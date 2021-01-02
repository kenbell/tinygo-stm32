// +build stm32l5x2

package lse

import "device/stm32"

var (
	LSE = Oscillator{
		Attributes: Attributes{
			ClockEnableRegister: &stm32.RCC.APB1ENR1,
			ClockEnableFlag:     stm32.RCC_APB1ENR1_PWREN,
			DefaultFrequency:    32768, // 32.768 KHz
		},
	}
)
