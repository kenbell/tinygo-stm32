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

// Config is the common clock configuration pattern, but bit values vary by target
type Config struct {
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

	// SYSCLK
	src := uint32(c.SYSCLKSource)
	stm32.RCC.CFGR.ReplaceBits(src, stm32.RCC_CFGR_SW_Msk, 0)
	for stm32.RCC.CFGR.Get()&stm32.RCC_CFGR_SWS_Msk != src<<stm32.RCC_CFGR_SWS_Pos {
	}
	SYSCLK.ClockFrequency = sysclkFromSource(c.SYSCLKSource).Frequency()

	// HCLK (set APBx dividers to max to ensure we keep clocks in spec)
	stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16), stm32.RCC_CFGR_PPRE1_Msk, 0)
	stm32.RCC.CFGR.ReplaceBits(uint32(PPREDividerDiv16)<<3, stm32.RCC_CFGR_PPRE2_Msk, 0)
	stm32.RCC.CFGR.ReplaceBits(uint32(c.AHBCLKDivider), stm32.RCC_CFGR_HPRE_Msk, 0)
	HCLK.ClockFrequency = SYSCLK.Frequency() / int64(hclkDividerFactor(c.AHBCLKDivider))

	// Decrease flash latency down to at most the requested value
	if (stm32.FLASH.ACR.Get() & stm32_FLASH_ACR_LATENCY_Msk) > flashLatency {
		stm32.FLASH.ACR.ReplaceBits(flashLatency, stm32_FLASH_ACR_LATENCY_Msk, 0)
	}

	// PCLK1
	stm32.RCC.CFGR.ReplaceBits(uint32(c.APB1CLKDivider), 0x1C00, 0)
	PCLK1.ClockFrequency = HCLK.Frequency() / int64(pclkDividerFactor(c.APB1CLKDivider))

	// PCLK2
	stm32.RCC.CFGR.ReplaceBits(uint32(c.APB2CLKDivider)<<3, 0xE000, 0)
	PCLK2.ClockFrequency = HCLK.Frequency() / int64(pclkDividerFactor(c.APB2CLKDivider))
}
