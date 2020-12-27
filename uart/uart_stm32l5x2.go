// +build stm32l5x2

package uart

import (
	"device/stm32"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
	"github.com/kenbell/tinygo-stm32/port"
)

var (
	LPUART1 = &UART{
		USART_Type: stm32.LPUART1,
		Attributes: &Attributes{
			TxPort:            port.GPIOG,
			TxPin:             7,
			TxPinAltFunc:      8,
			TxReg:             &stm32.LPUART1.TDR,
			RxPort:            port.GPIOG,
			RxPin:             8,
			RxPinAltFunc:      8,
			RxReg:             &stm32.LPUART1.RDR,
			StatusReg:         &stm32.LPUART1.ISR,
			StatusTxEmptyFlag: stm32.USART_ISR_TXE,
			Clock: clock.PeripheralConfig{
				Default:        clock.SourcePCLK1,
				EnableRegister: &stm32.RCC.APB1ENR2,
				EnableFlag:     stm32.RCC_APB1ENR2_LPUART1EN,
				SelectRegister: &stm32.RCC.CCIPR1,
				SelectValues: map[clock.Source]uint32{
					clock.SourcePCLK1:  0,
					clock.SourceSYSCLK: 1 << stm32.RCC_CCIPR1_LPUART1SEL_Pos,
					clock.SourceHSI:    2 << stm32.RCC_CCIPR1_LPUART1SEL_Pos,
					clock.SourceLSE:    3 << stm32.RCC_CCIPR1_LPUART1SEL_Pos,
				},
				SelectMask: stm32.RCC_CCIPR1_LPUART1SEL_Msk,
			},
		},
		NewInterrupt: func(h nvic.InterruptHandler) interrupt.Interrupt {
			lpuart1InterruptHandler = h
			return interrupt.New(stm32.IRQ_LPUART1, handleLPUART1Interrupt)
		},
	}
)

var lpuart1InterruptHandler nvic.InterruptHandler

func handleLPUART1Interrupt(intr interrupt.Interrupt) {
	if lpuart1InterruptHandler != nil {
		lpuart1InterruptHandler(intr)
	}
}
