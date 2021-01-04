// +build stm32f10x

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
	SourceHSI Source = stm32.RCC_CFGR_PLLSRC_HSI_Div2

	// SourceHSE indicates the High Speed External oscillator as PLL source
	SourceHSE Source = stm32.RCC_CFGR_PLLSRC_HSE_Div_PREDIV
)

type HSEPreDiv uint32

const (
	HSEPreDiv1 HSEPreDiv = 0
	HSEPreDiv2           = 1
)

// Config holds the config state of the PLL
type Config struct {
	// State of the PLL after configuration (on/off)
	State State

	// Source of the clock signal for the PLL
	Source Source

	// Prediv controls if HSE source is pre-divided by 2 before
	// the PLL
	Prediv HSEPreDiv

	// Mul is the multipler for the input clock signal
	Mul uint32
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
	stm32.RCC.CFGR.Set(uint32(c.Source)<<stm32.RCC_CFGR_PLLSRC_Pos |
		uint32(c.Prediv)<<stm32.RCC_CFGR_PLLXTPRE_Pos |
		(c.Mul-2)<<stm32.RCC_CFGR_PLLMUL_Pos)

	// Enable the PLL, wait until ready
	stm32.RCC.CR.SetBits(stm32.RCC_CR_PLLON)
	for !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
	}

	return c.calcPLLCLK()
}

func (c *Config) calcPLLCLK() int64 {
	// HSI freq is always halfed
	baseFreq := hsi.HSI.Frequency() / 2

	// HSE freq is halfed if requested
	if c.Source == SourceHSE {
		baseFreq = hse.HSE.Frequency()
		if c.Prediv == HSEPreDiv2 {
			baseFreq /= 2
		}
	}

	return baseFreq * int64(c.Mul)
}
