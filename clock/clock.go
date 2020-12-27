package clock

import (
	"device/stm32"

	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
)

var (
	PLL = pll.PLL1
	HSE = hse.HSE
)

type Type uint32

const (
	TypeSYSCLK Type = 0x00000001
	TypeHCLK   Type = 0x00000002
	TypePCLK1  Type = 0x00000004
	TypePCLK2  Type = 0x00000008
)

// Config is the common clock configuration pattern, but bit values vary by target
type Config struct {
	Types          Type
	SYSCLKSource   SYSCLKSource
	AHBCLKDivider  HPREDivider
	APB1CLKDivider PPREDivider
	APB2CLKDivider PPREDivider
}

// Apply the configuration
func (c Config) Apply(flashLatency uint32) {
	// Increase flash latency to at least requested value
	if (stm32.FLASH.ACR.Get() & stm32_FLASH_ACR_LATENCY_Msk) < flashLatency {
		stm32.FLASH.ACR.ReplaceBits(flashLatency, stm32_FLASH_ACR_LATENCY_Msk, 0)
	}

	if c.Types&TypeHCLK == TypeHCLK {
		// Set APBx dividers to max to ensure we keep clocks in spec
		if c.Types&TypePCLK1 == TypePCLK1 {
			stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16), stm32.RCC_CFGR_PPRE1_Msk, 0)
		}
		if c.Types&TypePCLK2 == TypePCLK2 {
			stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16)<<3, stm32.RCC_CFGR_PPRE2_Msk, 0)
		}

		// Set HCLK divider
		stm32.RCC.CFGR.ReplaceBits(uint32(c.AHBCLKDivider), stm32.RCC_CFGR_HPRE_Msk, 0)
	}

	if c.Types&TypeSYSCLK == TypeSYSCLK {
		src := uint32(c.SYSCLKSource)
		stm32.RCC.CFGR.ReplaceBits(src, stm32_RCC_CFGR_SW_Msk, 0)
		for stm32.RCC.CFGR.Get()&stm32_RCC_CFGR_SWS_Msk != src<<stm32_RCC_CFGR_SWS_Pos {
		}
	}

	// Decrease flash latency down to at most the requested value
	if (stm32.FLASH.ACR.Get() & stm32_FLASH_ACR_LATENCY_Msk) > flashLatency {
		stm32.FLASH.ACR.ReplaceBits(flashLatency, stm32_FLASH_ACR_LATENCY_Msk, 0)
	}

	if c.Types&TypePCLK1 == TypePCLK1 {
		stm32.RCC.CFGR.ReplaceBits(uint32(c.APB1CLKDivider), 0x1C00, 0)
	}

	if c.Types&TypePCLK2 == TypePCLK2 {
		stm32.RCC.CFGR.ReplaceBits(uint32(c.APB2CLKDivider)<<3, 0xE000, 0)
	}
}
