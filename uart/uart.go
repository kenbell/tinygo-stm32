package uart

import (
	"device/stm32"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

type Config struct {
	// BaudRate is the desired baud rate of the UART
	BaudRate uint32

	// Clock providing the clock signal to the UART
	Clock clock.Source

	//  PinMapping to use, typically 0 (most chips), or optionally 1 (some chips)
	PinMapping uint32
}

// UART represents the family of UART devices (USART, LPUART, UART)
type UART struct {
	*stm32.USART_Type
	Attributes *Attributes
	Clock      clock.PeripheralClock

	interruptHandler nvic.InterruptHandler
	NewInterrupt     func(nvic.InterruptHandler) interrupt.Interrupt
	Receive          func(b byte)
}

// Configure the UART.
func (uart *UART) Configure(config *Config) {
	// Default baud rate to 115200.
	if config.BaudRate == 0 {
		config.BaudRate = 115200
	}

	uart.configClock(config.Clock)

	// Configure pins (enable, set mode, etc)
	uart.configPins(config.PinMapping)

	// Set baud rate
	uart.SetBaudRate(config.BaudRate)

	// Enable USART port, tx, rx and rx interrupts
	uart.CR1.Set(stm32.USART_CR1_TE | stm32.USART_CR1_RE | stm32.USART_CR1_RXNEIE | stm32.USART_CR1_UE)

	// Enable RX IRQ
	intr := uart.NewInterrupt(uart.handleInterrupt)
	intr.SetPriority(0xc0)
	intr.Enable()
}

// SetReceiveCallback sets the callback to be invoked when a byte is received
// from the UART.
//
// This function signature is chosen to be identical to machine.GenericUART in
// tinygo to enable this UART to be used as an implementation of the tinygo
// UART.
func (uart *UART) SetReceiveCallback(fn func(b byte)) {
	uart.Receive = fn
}

// WriteByte writes a byte of data to the UART.
//
// This function signature is chosen to be identical to machine.GenericUART in
// tinygo to enable this UART to be used as an implementation of the tinygo
// UART.
func (uart *UART) WriteByte(c byte) error {
	uart.Attributes.TxReg.Set(uint32(c))

	for !uart.Attributes.StatusReg.HasBits(uart.Attributes.StatusTxEmptyFlag) {
	}
	return nil
}

// handleInterrupt should be called from the appropriate interrupt handler for
// this UART instance
func (uart *UART) handleInterrupt(interrupt.Interrupt) {
	b := byte(uart.Attributes.RxReg.Get() & 0xFF)
	if uart.Receive != nil {
		uart.Receive(b)
	}
}

// SetBaudRate sets the communication speed for the UART
func (uart *UART) SetBaudRate(br uint32) {
	mult := uart.Attributes.BaudMultiplier
	if mult == 0 {
		mult = 1
	}

	divider := mult * uint32(uart.Clock.Frequency()/int64(br))
	uart.BRR.Set(divider)
}

func (uart *UART) configClock(src clock.Source) {
	uart.Clock = uart.Attributes.Clock.Apply(src)
}

func (uart *UART) configPins(pinMapping uint32) {
	mapping := &uart.Attributes.PinMapping[pinMapping]

	mapping.TxPin.ConfigureAltFunc(
		uart.Attributes.TxPinConfig,
		mapping.TxPinAltFunc)

	mapping.RxPin.ConfigureAltFunc(
		uart.Attributes.RxPinConfig,
		mapping.RxPinAltFunc)
}
