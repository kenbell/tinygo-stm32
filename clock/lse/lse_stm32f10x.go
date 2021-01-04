// +build stm32f10x

package lse

import "device/stm32"

var (
	LSE = Oscillator{
		Attributes: Attributes{
			ClockEnableRegister:         &stm32.RCC.APB1ENR,
			ClockEnableFlag:             stm32.RCC_APB1ENR_PWREN,
			DefaultFrequency:            32768, // 32.768 KHz
			DisableBackupDomainRegister: &stm32.PWR.CR,
			DisableBackupDomainFlag:     stm32.PWR_CR_DBP,
		},
	}
)
