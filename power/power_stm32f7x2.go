// +build stm32f7x2

package power

import "device/stm32"

// EnableOverdrive configures the system to reach 216 MHz
func EnableOverdrive() {
	if !stm32.RCC.APB1ENR.HasBits(stm32.RCC_APB1ENR_PWREN) {
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_PWREN)
		_ = stm32.RCC.APB1ENR.Get()
	}

	stm32.PWR.CR1.SetBits(stm32.PWR_CR1_ODEN)
	for !stm32.PWR.CSR1.HasBits(stm32.PWR_CSR1_ODRDY) {
	}

	stm32.PWR.CR1.SetBits(stm32.PWR_CR1_ODSWEN)
	for !stm32.PWR.CSR1.HasBits(stm32.PWR_CSR1_ODSWRDY) {
	}

}
