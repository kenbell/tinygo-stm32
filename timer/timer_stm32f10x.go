// +build stm32f10x

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
		Clock:        &clock.PCLK2,
		NewInterrupt: newIRQ_TIM1_UPHandler,
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
		Clock:        &clock.PCLK1,
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
		Clock:        &clock.PCLK1,
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
		Clock:        &clock.PCLK1,
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
		Clock:        &clock.PCLK1,
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
		Clock:        &clock.PCLK1,
		NewInterrupt: newIRQ_TIM6Handler,
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
		Clock:        &clock.PCLK1,
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
		Clock:        &clock.PCLK2,
		NewInterrupt: newIRQ_TIM8_UPHandler,
	}
)

var (
	chainIRQ_TIM1_UP            nvic.HandlerChain
	chainIRQ_TIM2               nvic.HandlerChain
	chainIRQ_TIM3               nvic.HandlerChain
	chainIRQ_TIM4               nvic.HandlerChain
	chainIRQ_TIM5               nvic.HandlerChain
	chainIRQ_TIM6               nvic.HandlerChain
	chainIRQ_TIM7               nvic.HandlerChain
	chainIRQ_TIM8_UP            nvic.HandlerChain
	chainIRQ_TIM1_BRK_TIM9      nvic.HandlerChain
	chainIRQ_TIM1_TRG_COM_TIM11 nvic.HandlerChain
	chainIRQ_TIM8_BRK_TIM12     nvic.HandlerChain
	chainIRQ_TIM8_TRG_COM_TIM14 nvic.HandlerChain
)

func newIRQ_TIM1_UPHandler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM1_UP.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM1_UP, handleIRQ_TIM1_UPInterrupt)
	return &intr
}

func handleIRQ_TIM1_UPInterrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM1_UP.Call(intr)
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

func newIRQ_TIM6Handler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM6.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM6, handleIRQ_TIM6Interrupt)
	return &intr
}

func handleIRQ_TIM6Interrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM6.Call(intr)
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

func newIRQ_TIM8_UPHandler(h nvic.InterruptHandler) *interrupt.Interrupt {
	if chainIRQ_TIM8_UP.Register(h) {
		return nil
	}
	intr := interrupt.New(stm32.IRQ_TIM8_UP, handleIRQ_TIM8_UPInterrupt)
	return &intr
}

func handleIRQ_TIM8_UPInterrupt(intr interrupt.Interrupt) {
	chainIRQ_TIM8_UP.Call(intr)
}
