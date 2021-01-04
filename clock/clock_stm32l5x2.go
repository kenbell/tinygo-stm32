// +build stm32l5x2

package clock

import (
	"device/stm32"
)

type SYSCLKSource uint32

const (
	SYSCLKSourceMSI SYSCLKSource = 0
	SYSCLKSourceHSI              = 1
	SYSCLKSourceHSE              = 2
	SYSCLKSourcePLL              = 3
)

// Work-around SVD definitions to create useful constants
const (
	stm32_FLASH_ACR_LATENCY_Msk = stm32.Flash_ACR_LATENCY_Msk
)

func sysclkFromSource(src SYSCLKSource) PeripheralClock {
	switch src {
	case SYSCLKSourceMSI:
		return MSI
	case SYSCLKSourceHSE:
		return HSE
	case SYSCLKSourceHSI:
		return HSI
	case SYSCLKSourcePLL:
		return PLL
	}

	return nil
}
