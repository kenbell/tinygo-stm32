// +build stm32f7x2

package pll

import (
	"device/stm32"

	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/hsi"
)

var (
	PLL1 = PLL{}
)

// Source indicates the PLL source
type Source uint32

const (
	// SourceHSI indicates the High Speed Internal oscillator as PLL source
	SourceHSI Source = 0

	// SourceHSE indicates the High Speed External oscillator as PLL source
	SourceHSE Source = (1 << 22)
)

// Config holds the config state of the PLL
type Config struct {
	State  State
	Source Source
	M      uint32
	N      uint32
	P      uint32
	Q      uint32
}

// Apply modifies the PLL state, waiting for completion
func (c *Config) Apply() int64 {
	if c.State == StateNone {
		return 0
	}

	// Disable the PLL, wait until reset
	stm32.RCC.CR.ClearBits(stm32.RCC_CR_PLLON)
	for stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
	}

	// If turning off, we're done
	if c.State == StateOff {
		return 0
	}

	// Configure the PLL
	stm32.RCC.PLLCFGR.Set(0x20000000 |
		uint32(c.Source) |
		c.M |
		(c.N << stm32.RCC_PLLCFGR_PLLN_Pos) |
		(((c.P >> 1) - 1) << stm32.RCC_PLLCFGR_PLLP_Pos) |
		(c.Q << stm32.RCC_PLLCFGR_PLLQ_Pos))

	// Enable the PLL, wait until ready
	stm32.RCC.CR.SetBits(stm32.RCC_CR_PLLON)
	for !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
	}

	return c.calcPLLCLK()
}

func (c *Config) calcPLLCLK() int64 {
	baseFreq := hsi.HSI.Frequency()
	if c.Source == SourceHSE {
		baseFreq = hse.HSE.Frequency()
	}

	return ((baseFreq / int64(c.M)) * int64(c.N)) / int64(c.P)
}
