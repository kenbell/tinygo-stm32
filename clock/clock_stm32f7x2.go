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
	stm32_RCC_CFGR_SW_Msk  = stm32.RCC_CFGR_SW0_Msk | stm32.RCC_CFGR_SW1_Msk
	stm32_RCC_CFGR_SWS_Msk = stm32.RCC_CFGR_SWS0_Msk | stm32.RCC_CFGR_SWS1_Msk
	stm32_RCC_CFGR_SWS_Pos = stm32.RCC_CFGR_SWS0_Pos

	stm32_FLASH_ACR_LATENCY_Msk = stm32.FLASH_ACR_LATENCY_Msk
)
