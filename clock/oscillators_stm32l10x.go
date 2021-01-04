// +build stm32f10x

package clock

import (
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/lse"
	"github.com/kenbell/tinygo-stm32/clock/lsi"
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

	// PLL indicates the configuration of the main PLL
	PLL *pll.Config
}

// Apply changes the MCU oscillator config
func (c OscillatorConfig) Apply() {
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

	if c.PLL != nil {
		PLL.Configure(c.PLL)
	}
}
