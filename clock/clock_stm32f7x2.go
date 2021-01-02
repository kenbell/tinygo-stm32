// +build stm32f7x2

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

type HPREDivider uint32

const (
	HPREDividerDiv1   HPREDivider = 0x00000000
	HPREDividerDiv2               = 0x00000080
	HPREDividerDiv4               = 0x00000090
	HPREDividerDiv8               = 0x000000A0
	HPREDividerDiv16              = 0x000000B0
	HPREDividerDiv64              = 0x000000C0
	HPREDividerDiv128             = 0x000000D0
	HPREDividerDiv256             = 0x000000E0
	HPREDividerDiv512             = 0x000000F0
)

type PPREDivider uint32

const (
	PPREDividerDiv1  PPREDivider = 0x00000000
	PPREDividerDiv2              = 0x00001000
	PPREDividerDiv4              = 0x00001400
	PPREDividerDiv8              = 0x00001800
	PPREDividerDiv16             = 0x00001C00
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

func hclkDividerFactor(div HPREDivider) uint32 {
	switch div {
	case HPREDividerDiv1:
		return 1
	case HPREDividerDiv2:
		return 2
	case HPREDividerDiv4:
		return 4
	case HPREDividerDiv8:
		return 8
	case HPREDividerDiv16:
		return 16
	case HPREDividerDiv64:
		return 64
	case HPREDividerDiv128:
		return 128
	case HPREDividerDiv256:
		return 256
	case HPREDividerDiv512:
		return 512
	}

	return 0
}

func pclkDividerFactor(div PPREDivider) uint32 {
	switch div {
	case PPREDividerDiv1:
		return 1
	case PPREDividerDiv2:
		return 2
	case PPREDividerDiv4:
		return 4
	case PPREDividerDiv8:
		return 8
	case PPREDividerDiv16:
		return 16
	}

	return 0
}
