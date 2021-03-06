// +build stm32f7x2

package clock

import (
	"device/stm32"
)

type SYSCLKSource uint32

const (
	SYSCLKSourceHSI SYSCLKSource = stm32.RCC_CFGR_SW_HSI
	SYSCLKSourceHSE              = stm32.RCC_CFGR_SW_HSE
	SYSCLKSourcePLL              = stm32.RCC_CFGR_SW_PLL
)

// Work-around SVD definitions to create useful constants
const (
	stm32_FLASH_ACR_LATENCY_Msk = stm32.FLASH_ACR_LATENCY_Msk
)

func sysclkFromSource(src SYSCLKSource) PeripheralClock {
	switch src {
	case SYSCLKSourceHSE:
		return HSE
	case SYSCLKSourceHSI:
		return HSI
	case SYSCLKSourcePLL:
		return PLL
	}

	return nil
}
