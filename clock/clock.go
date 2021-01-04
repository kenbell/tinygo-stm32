package clock

import (
	"device/stm32"

	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
)

var (
	PLL = &pll.PLL1
	HSE = &hse.HSE
)

type HPREDivider uint32

const (
	HPREDividerDiv1   HPREDivider = stm32.RCC_CFGR_HPRE_Div1
	HPREDividerDiv2               = stm32.RCC_CFGR_HPRE_Div2
	HPREDividerDiv4               = stm32.RCC_CFGR_HPRE_Div4
	HPREDividerDiv8               = stm32.RCC_CFGR_HPRE_Div8
	HPREDividerDiv16              = stm32.RCC_CFGR_HPRE_Div16
	HPREDividerDiv64              = stm32.RCC_CFGR_HPRE_Div64
	HPREDividerDiv128             = stm32.RCC_CFGR_HPRE_Div128
	HPREDividerDiv256             = stm32.RCC_CFGR_HPRE_Div256
	HPREDividerDiv512             = stm32.RCC_CFGR_HPRE_Div512
)

type PPREDivider uint32

// Defined as values until generator fix so that PPRE1 constants
// always created.
const (
	PPREDividerDiv1  PPREDivider = 0x0
	PPREDividerDiv2              = 0x4
	PPREDividerDiv4              = 0x5
	PPREDividerDiv8              = 0x6
	PPREDividerDiv16             = 0x7
)

// Config is the common clock configuration pattern, but bit values vary by target
type Config struct {
	SYSCLKSource    SYSCLKSource
	AHBCLKDivider   HPREDivider
	APB1CLKDivider  PPREDivider
	APB2CLKDivider  PPREDivider
	FlashWaitStates uint32
}

// Apply the configuration
func (c Config) Apply() {
	// Increase flash latency to at least requested value
	if (stm32.FLASH.ACR.Get() & stm32_FLASH_ACR_LATENCY_Msk) < c.FlashWaitStates {
		stm32.FLASH.ACR.ReplaceBits(c.FlashWaitStates, stm32_FLASH_ACR_LATENCY_Msk, 0)
	}

	// SYSCLK
	src := uint32(c.SYSCLKSource)
	stm32.RCC.CFGR.ReplaceBits(src, stm32.RCC_CFGR_SW_Msk, 0)
	for stm32.RCC.CFGR.Get()&stm32.RCC_CFGR_SWS_Msk != src<<stm32.RCC_CFGR_SWS_Pos {
	}
	SYSCLK.ClockFrequency = sysclkFromSource(c.SYSCLKSource).Frequency()

	// HCLK (set APBx dividers to max to ensure we keep clocks in spec)
	stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16)<<stm32.RCC_CFGR_PPRE1_Pos, stm32.RCC_CFGR_PPRE1_Msk, 0)
	stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16)<<stm32.RCC_CFGR_PPRE2_Pos, stm32.RCC_CFGR_PPRE2_Msk, 0)
	stm32.RCC.CFGR.ReplaceBits(uint32(c.AHBCLKDivider)<<stm32.RCC_CFGR_HPRE_Pos, stm32.RCC_CFGR_HPRE_Msk, 0)
	HCLK.ClockFrequency = SYSCLK.Frequency() / int64(hclkDividerFactor(c.AHBCLKDivider))

	// Decrease flash latency down to at most the requested value
	if (stm32.FLASH.ACR.Get() & stm32_FLASH_ACR_LATENCY_Msk) > c.FlashWaitStates {
		stm32.FLASH.ACR.ReplaceBits(c.FlashWaitStates, stm32_FLASH_ACR_LATENCY_Msk, 0)
	}

	// PCLK1
	stm32.RCC.CFGR.ReplaceBits(uint32(c.APB1CLKDivider), 0x1C00, 0)
	PCLK1.ClockFrequency = HCLK.Frequency() / int64(pclkDividerFactor(c.APB1CLKDivider))

	// PCLK2
	stm32.RCC.CFGR.ReplaceBits(uint32(c.APB2CLKDivider)<<3, 0xE000, 0)
	PCLK2.ClockFrequency = HCLK.Frequency() / int64(pclkDividerFactor(c.APB2CLKDivider))
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
