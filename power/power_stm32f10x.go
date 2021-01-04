// +build stm32f10x

package power

import "device/stm32"

// EnableOverdrive configures the system to reach 72 MHz
func EnableOverdrive() {
	if !stm32.RCC.APB1ENR.HasBits(stm32.RCC_APB1ENR_PWREN) {
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_PWREN)
		_ = stm32.RCC.APB1ENR.Get()
	}
}
