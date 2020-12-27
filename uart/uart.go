package uart

import (
	"device/stm32"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

type Config struct {
	BaudRate uint32
	Clock    clock.Source
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
func (uart *UART) Configure(config Config) {
	// Default baud rate to 115200.
	if config.BaudRate == 0 {
		config.BaudRate = 115200
	}

	uart.configClock(config.Clock)

	// Configure pins (enable, set mode, etc)
	uart.configRxPin()
	uart.configTxPin()

	// Set baud rate
	uart.SetBaudRate(config.BaudRate)

	// Enable USART port, tx, rx and rx interrupts
	uart.CR1.Set(stm32.USART_CR1_TE | stm32.USART_CR1_RE | stm32.USART_CR1_RXNEIE | stm32.USART_CR1_UE)

	// Enable RX IRQ
	intr := uart.NewInterrupt(uart.handleInterrupt)
	intr.SetPriority(0xc0)
	intr.Enable()
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
	divider := uint32(uart.Clock.Frequency() / int64(br))
	uart.BRR.Set(divider)
}

// WriteByte writes a byte of data to the UART.
func (uart *UART) WriteByte(c byte) error {
	uart.Attributes.TxReg.Set(uint32(c))

	for !uart.Attributes.StatusReg.HasBits(uart.Attributes.StatusTxEmptyFlag) {
	}
	return nil
}

func (uart *UART) configClock(src clock.Source) {
	uart.Clock = uart.Attributes.Clock.Apply(src)
}

func (uart *UART) configTxPin() {
	port := uart.Attributes.TxPort
	pos := uart.Attributes.TxPin * 2

	port.EnableClock()

	port.MODER.ReplaceBits(stm32.GPIOModeOutputAltFunc, 0x3, pos)
	port.OSPEEDR.ReplaceBits(stm32.GPIOSpeedHigh, 0x3, pos)
	port.PUPDR.ReplaceBits(stm32.GPIOPUPDRPullUp, 0x3, pos)

	port.SetAltFunc(uart.Attributes.TxPin, uart.Attributes.TxPinAltFunc)
}

func (uart *UART) configRxPin() {
	port := uart.Attributes.RxPort
	pos := uart.Attributes.RxPin * 2

	port.EnableClock()

	port.MODER.ReplaceBits(stm32.GPIOModeOutputAltFunc, 0x3, pos)
	port.PUPDR.ReplaceBits(stm32.GPIOPUPDRFloating, 0x3, pos)

	port.SetAltFunc(uart.Attributes.RxPin, uart.Attributes.RxPinAltFunc)
}
