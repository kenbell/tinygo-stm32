// +build stm32f10x

package uart

import (
	"device/stm32"
	"machine"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

var (
	USART1 = &UART{
		USART_Type: stm32.USART1,
		Attributes: &Attributes{
			PinMapping: []PinMapping{
				{TxPin: machine.PA9, RxPin: machine.PA10},
				{TxPin: machine.PB6, RxPin: machine.PB7},
			},
			TxReg:             &stm32.USART1.DR,
			TxPinConfig:       machine.PinConfig{Mode: machine.PinOutput50MHz + machine.PinOutputModeAltPushPull},
			RxReg:             &stm32.USART1.DR,
			RxPinConfig:       machine.PinConfig{Mode: machine.PinInputModeFloating},
			StatusReg:         &stm32.USART1.SR,
			StatusTxEmptyFlag: stm32.USART_SR_TXE,
			Clock: clock.PeripheralConfig{
				Default:        clock.SourcePCLK2,
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_USART1EN,
			},
		},
		NewInterrupt: func(h nvic.InterruptHandler) interrupt.Interrupt {
			uart1InterruptHandler = h
			return interrupt.New(stm32.IRQ_USART1, handleUSART1Interrupt)
		},
	}
)

var uart1InterruptHandler nvic.InterruptHandler

func handleUSART1Interrupt(intr interrupt.Interrupt) {
	if uart1InterruptHandler != nil {
		uart1InterruptHandler(intr)
	}
}
