// +build stm32f10x

package clock

import (
	"device/stm32"
)

type SYSCLKSource uint32

const (
	SYSCLKSourceHSI SYSCLKSource = 0
	SYSCLKSourceHSE              = 1
	SYSCLKSourcePLL              = 2
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
