// +build stm32f7x2

package lse

import "device/stm32"

var (
	LSE = &Oscillator{
		Attributes: Attributes{
			ClockEnableRegister: &stm32.RCC.APB1ENR,
			ClockEnableFlag:     stm32.RCC_APB1ENR_PWREN,
		},
	}
)
