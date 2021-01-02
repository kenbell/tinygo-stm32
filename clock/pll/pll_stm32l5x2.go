// +build stm32l5x2

package pll

import (
	"device/stm32"

	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/msi"
)

var (
	PLL1 = PLL{}
)

// Source indicates the PLL source
type Source uint32

const (
	// SourceNone indicates no clock as PLL source
	SourceNone Source = 0

	// SourceMSI indicates the Medium Speed Internal oscillator as PLL source
	SourceMSI Source = 1

	// SourceHSI indicates the High Speed Internal oscillator as PLL source
	SourceHSI Source = 2

	// SourceHSE indicates the High Speed External oscillator as PLL source
	SourceHSE Source = 3
)

// Config holds the config state of the PLL
type Config struct {
	State  State
	Source Source
	M      uint32
	N      uint32
	P      uint32
	Q      uint32
	R      uint32
}

const (
	stm32_RCC_PLL_SAI3CLK = stm32.RCC_PLLCFGR_PLLPEN
	stm32_RCC_PLL_48M1CLK = stm32.RCC_PLLCFGR_PLLQEN
	stm32_RCC_PLL_SYSCLK  = stm32.RCC_PLLCFGR_PLLREN
)

// Apply modifies the PLL state, waiting for completion
//
// The new PLLCLK frequency is returned, which is the clock
// available as an input to SYSCLK.
func (c *Config) Apply() int64 {
	if c.State == StateNone {
		return 0
	}

	// Disable the PLL, wait until reset
	stm32.RCC.CR.ClearBits(stm32.RCC_CR_PLLON)
	for stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
	}

	// If turning off
	if c.State == StateOff {
		// Save power - disable PLL inputs if no PLLs on
		if stm32.RCC.CR.Get()&(stm32.RCC_CR_PLLSAI1RDY|stm32.RCC_CR_PLLSAI2RDY) == 0 {
			stm32.RCC.PLLCFGR.ReplaceBits(0, stm32.RCC_PLLCFGR_PLLSRC_Msk, 0)
		}

		stm32.RCC.PLLCFGR.ClearBits(stm32_RCC_PLL_SYSCLK | stm32_RCC_PLL_48M1CLK | stm32_RCC_PLL_SAI3CLK)
		for stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
		}

		return 0
	}

	// Configure the PLL
	stm32.RCC.PLLCFGR.ReplaceBits(uint32(c.Source)|
		(c.M-1)<<stm32.RCC_PLLCFGR_PLLM_Pos|
		(c.N<<stm32.RCC_PLLCFGR_PLLN_Pos)|
		(((c.Q>>1)-1)<<stm32.RCC_PLLCFGR_PLLQ_Pos)|
		(((c.R>>1)-1)<<stm32.RCC_PLLCFGR_PLLR_Pos)|
		(c.P<<stm32.RCC_PLLCFGR_PLLPDIV_Pos),
		stm32.RCC_PLLCFGR_PLLSRC_Msk|stm32.RCC_PLLCFGR_PLLM_Msk|
			stm32.RCC_PLLCFGR_PLLN_Msk|stm32.RCC_PLLCFGR_PLLP_Msk|
			stm32.RCC_PLLCFGR_PLLR_Msk|stm32.RCC_PLLCFGR_PLLPDIV_Msk,
		0)

	// Enable the PLL, wait until ready
	stm32.RCC.CR.SetBits(stm32.RCC_CR_PLLON)
	stm32.RCC.PLLCFGR.SetBits(stm32.RCC_PLLCFGR_PLLREN) // = RCC_PLL_SYSCLK
	for !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
	}

	return c.calcPLLCLK()
}

func (c *Config) calcPLLCLK() int64 {
	var baseFreq int64
	switch c.Source {
	case SourceMSI:
		baseFreq = msi.MSI.Frequency()
	case SourceHSI:
		baseFreq = hsi.HSI.Frequency()
	case SourceHSE:
		baseFreq = hse.HSE.Frequency()
	}

	return ((baseFreq / int64(c.M)) * int64(c.N)) / int64(c.R)
}
