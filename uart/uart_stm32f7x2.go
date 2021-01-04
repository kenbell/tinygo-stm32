// +build stm32f7x2

package uart

import (
	"device/stm32"
	"machine"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

var (
	USART3 = &UART{
		USART_Type: stm32.USART3,
		Attributes: &Attributes{
			PinMapping: []PinMapping{
				{TxPin: machine.PD8, TxPinAltFunc: 7, RxPin: machine.PD9, RxPinAltFunc: 7},
			},
			TxReg:             &stm32.USART3.TDR,
			RxReg:             &stm32.USART3.RDR,
			StatusReg:         &stm32.USART3.ISR,
			StatusTxEmptyFlag: stm32.USART_ISR_TXE,
			Clock: clock.PeripheralConfig{
				Default:        clock.SourcePCLK1,
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_USART3EN,
				SelectRegister: &stm32.RCC.DCKCFGR2,
				SelectValues: map[clock.Source]uint32{
					clock.SourcePCLK1:  0,
					clock.SourceSYSCLK: 1 << stm32.RCC_DCKCFGR2_USART3SEL_Pos,
					clock.SourceHSI:    2 << stm32.RCC_DCKCFGR2_USART3SEL_Pos,
					clock.SourceLSE:    3 << stm32.RCC_DCKCFGR2_USART3SEL_Pos,
				},
				SelectMask: stm32.RCC_DCKCFGR2_USART3SEL_Msk,
			},
		},
		NewInterrupt: func(h nvic.InterruptHandler) interrupt.Interrupt {
			uart3InterruptHandler = h
			return interrupt.New(stm32.IRQ_USART3, handleUSART3Interrupt)
		},
	}
)

var uart3InterruptHandler nvic.InterruptHandler

func handleUSART3Interrupt(intr interrupt.Interrupt) {
	if uart3InterruptHandler != nil {
		uart3InterruptHandler(intr)
	}
}
