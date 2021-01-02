package nucleof722ze

import (
	"github.com/kenbell/tinygo-stm32/clock/"
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
)

var (
	// DefaultPLLConfig is a basic oscillator config that provides a 216 MHz PLLCLK
	DefaultPLLConfig = clock.OscillatorConfig{
		HSE: &hse.Config{
			State:     hse.StateBypass,
			Frequency: 8000000, // 8 MHz
		},
		PLL: &pll.Config{
			Source: pll.SourceHSE,
			State:  pll.StateOn,
			M:      4,
			N:      216,
			P:      2,
			Q:      9,
		},
	}

	// DefaultPeripheralClocks is a default config compatible with DefaultPLLConfig
	DefaultPeripheralClocks = clock.Config{
		SYSCLKSource:   clock.SYSCLKSourcePLL,
		AHBCLKDivider:  clock.HPREDividerDiv1,
		APB1CLKDivider: clock.PPREDividerDiv4,
		APB2CLKDivider: clock.PPREDividerDiv2,
	}
)
