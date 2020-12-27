// +build stm32l5x2

package clock

import (
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/hsi48"
	"github.com/kenbell/tinygo-stm32/clock/lse"
	"github.com/kenbell/tinygo-stm32/clock/lsi"
	"github.com/kenbell/tinygo-stm32/clock/msi"
	"github.com/kenbell/tinygo-stm32/clock/pll"
)

// OscillatorConfig configures the oscillators and main PLL
type OscillatorConfig struct {
	// HSE indicates the configuration of the High Speed External oscillator
	HSE *hse.Config

	// LSE indicates the configuration of the Low Speed External oscillator
	LSE *lse.Config

	// HSI indicates the configuration of the High Speed Internal oscillator
	HSI *hsi.Config

	// LSI indicates the configuration of the Low Speed Internal oscillator
	LSI *lsi.Config

	// MSI indicates the configuration of the Medium Speed Internal oscillator
	MSI *msi.Config

	// HSI48 indicates the configuration of the High Speed 48 MHz Internal oscillator
	HSI48 *hsi48.Config

	// PLL indicates the configuration of the main PLL
	PLL *pll.Config
}

// Apply changes the MCU oscillator config
func (c OscillatorConfig) Apply() {
	if c.MSI != nil {
		MSI.Configure(c.MSI)
	}

	if c.HSE != nil {
		HSE.Configure(c.HSE)
	}

	if c.HSI != nil {
		HSI.Configure(c.HSI)
	}

	if c.LSI != nil {
		LSI.Configure(c.LSI)
	}

	if c.LSE != nil {
		LSE.Configure(c.LSE)
	}

	if c.HSI48 != nil {
		HSI48.Configure(c.HSI48)
	}

	if c.PLL != nil {
		c.PLL.Apply()
	}
}
