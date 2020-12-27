// +build stm32f7x2

package timer

import (
	"device/stm32"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

var (
	TIM1 = &Timer{
		TIM_Type: stm32.TIM1,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision | FeatureRepetitionCounter,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_TIM1EN,
			},
		},
		Clock:        clock.PCLK2,
		NewInterrupt: newIRQ_TIM1_UP_TIM10Handler,
	}

	TIM2 = &Timer{
		TIM_Type: stm32.TIM2,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM2EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM2Handler,
	}

	TIM3 = &Timer{
		TIM_Type: stm32.TIM3,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM3EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM3Handler,
	}

	TIM4 = &Timer{
		TIM_Type: stm32.TIM4,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM4EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM4Handler,
	}

	TIM5 = &Timer{
		TIM_Type: stm32.TIM5,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM5EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM5Handler,
	}

	TIM6 = &Timer{
		TIM_Type: stm32.TIM6,
		Attributes: &Attributes{
			Features: FeatureNone,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM6EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM6_DACHandler,
	}

	TIM7 = &Timer{
		TIM_Type: stm32.TIM7,
		Attributes: &Attributes{
			Features: FeatureNone,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM7EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM7Handler,
	}

	TIM8 = &Timer{
		TIM_Type: stm32.TIM8,
		Attributes: &Attributes{
			Features: FeatureCounterModelSelect | FeatureClockDivision | FeatureRepetitionCounter,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_TIM8EN,
			},
		},
		Clock:        clock.PCLK2,
		NewInterrupt: newIRQ_TIM8_UP_TIM13Handler,
	}

	TIM9 = &Timer{
		TIM_Type: stm32.TIM9,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_TIM9EN,
			},
		},
		Clock:        clock.PCLK2,
		NewInterrupt: newIRQ_TIM1_BRK_TIM9Handler,
	}

	TIM10 = &Timer{
		TIM_Type: stm32.TIM10,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_TIM10EN,
			},
		},
		Clock:        clock.PCLK2,
		NewInterrupt: newIRQ_TIM1_UP_TIM10Handler,
	}

	TIM11 = &Timer{
		TIM_Type: stm32.TIM11,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB2ENR,
				EnableFlag:     stm32.RCC_APB2ENR_TIM11EN,
			},
		},
		Clock:        clock.PCLK2,
		NewInterrupt: newIRQ_TIM1_TRG_COM_TIM11Handler,
	}

	TIM12 = &Timer{
		TIM_Type: stm32.TIM12,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM12EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM8_BRK_TIM12Handler,
	}

	TIM13 = &Timer{
		TIM_Type: stm32.TIM13,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM13EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM8_UP_TIM13Handler,
	}

	TIM14 = &Timer{
		TIM_Type: stm32.TIM14,
		Attributes: &Attributes{
			Features: FeatureClockDivision,
			Clock: clock.PeripheralConfig{
				EnableRegister: &stm32.RCC.APB1ENR,
				EnableFlag:     stm32.RCC_APB1ENR_TIM14EN,
			},
		},
		Clock:        clock.PCLK1,
		NewInterrupt: newIRQ_TIM8_TRG_COM_TIM14Handler,
	}
)

var (
	chainIRQ_TIM1_UP_TIM10      nvic.HandlerChain
	chainIRQ_TIM2               nvic.HandlerChain
	chainIRQ_TIM3               nvic.HandlerChain
	chainIRQ_TIM4               nvic.HandlerChain
	chainIRQ_TIM5               nvic.HandlerChain
	chainIRQ_TIM6_DAC           nvic.HandlerChain
	chainIRQ_TIM7               nvic.HandlerChain
	chainIRQ_TIM8_UP_TIM13      nvic.HandlerChain
	chainIRQ_TIM1_BRK_TIM9      nvic.HandlerChain
	chainIRQ_TIM1_TRG_COM_TIM11 nvic.HandlerChain
	chainIRQ_TIM8_BRK_TIM12     nvic.HandlerChain
	chainIRQ_TIM8_TRG_COM_TIM14 nvic.HandlerChain
)

func newIRQ_TIM1_UP_TIM10Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM1_UP_TIM10.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM1_UP_TIM10, handleIRQ_TIM1_UP_TIM10Interrupt)
	return &intr
}

func handleIRQ_TIM1_UP_TIM10Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM1_UP_TIM10.Call(intr)
}

func newIRQ_TIM2Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM2.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM2, handleIRQ_TIM2Interrupt)
	return &intr
}

func handleIRQ_TIM2Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM2.Call(intr)
}

func newIRQ_TIM3Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM3.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM3, handleIRQ_TIM3Interrupt)
	return &intr
}

func handleIRQ_TIM3Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM3.Call(intr)
}

func newIRQ_TIM4Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM4.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM4, handleIRQ_TIM4Interrupt)
	return &intr
}

func handleIRQ_TIM4Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM4.Call(intr)
}

func newIRQ_TIM5Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM5.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM5, handleIRQ_TIM5Interrupt)
	return &intr
}

func handleIRQ_TIM5Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM5.Call(intr)
}

func newIRQ_TIM6_DACHandler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM6_DAC.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM6_DAC, handleIRQ_TIM6_DACInterrupt)
	return &intr
}

func handleIRQ_TIM6_DACInterrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM6_DAC.Call(intr)
}

func newIRQ_TIM7Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM7.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM7, handleIRQ_TIM7Interrupt)
	return &intr
}

func handleIRQ_TIM7Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM7.Call(intr)
}

func newIRQ_TIM8_UP_TIM13Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM8_UP_TIM13.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM8_UP_TIM13, handleIRQ_TIM8_UP_TIM13Interrupt)
	return &intr
}

func handleIRQ_TIM8_UP_TIM13Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM8_UP_TIM13.Call(intr)
}

func newIRQ_TIM1_BRK_TIM9Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM1_BRK_TIM9.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM1_BRK_TIM9, handleIRQ_TIM1_BRK_TIM9Interrupt)
	return &intr
}

func handleIRQ_TIM1_BRK_TIM9Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM1_BRK_TIM9.Call(intr)
}

func newIRQ_TIM1_TRG_COM_TIM11Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM1_TRG_COM_TIM11.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM1_TRG_COM_TIM11, handleIRQ_TIM1_TRG_COM_TIM11Interrupt)
	return &intr
}

func handleIRQ_TIM1_TRG_COM_TIM11Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM1_TRG_COM_TIM11.Call(intr)
}

func newIRQ_TIM8_BRK_TIM12Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM8_BRK_TIM12.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM8_BRK_TIM12, handleIRQ_TIM8_BRK_TIM12Interrupt)
	return &intr
}

func handleIRQ_TIM8_BRK_TIM12Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM8_BRK_TIM12.Call(intr)
}

func newIRQ_TIM8_TRG_COM_TIM14Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM8_TRG_COM_TIM14.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM8_TRG_COM_TIM14, handleIRQ_TIM8_TRG_COM_TIM14Interrupt)
	return &intr
}

func handleIRQ_TIM8_TRG_COM_TIM14Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM8_TRG_COM_TIM14.Call(intr)
}
