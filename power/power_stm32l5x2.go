// +build stm32l5x2

package power

import "device/stm32"

// EnableOverdrive configures the system to reach max clock
func EnableOverdrive() {
	if !stm32.RCC.APB1ENR1.HasBits(stm32.RCC_APB1ENR1_PWREN) {
		stm32.RCC.APB1ENR1.SetBits(stm32.RCC_APB1ENR1_PWREN)
		_ = stm32.RCC.APB1ENR1.Get()
	}

	// PWR_VOLTAGESCALING_CONFIG
	stm32.PWR.CR1.ReplaceBits(0, stm32.PWR_CR1_VOS_Msk, 0)
	_ = stm32.PWR.CR1.Get()
}
