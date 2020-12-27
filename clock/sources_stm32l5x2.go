// +build stm32l5x2

package clock

import (
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/hsi48"
	"github.com/kenbell/tinygo-stm32/clock/lse"
	"github.com/kenbell/tinygo-stm32/clock/lsi"
	"github.com/kenbell/tinygo-stm32/clock/msi"
)

const (
	SourceNone Source = iota
	SourcePCLK1
	SourcePCLK2
	SourceSYSCLK
	SourceHSE
	SourceHSI
	SourceLSE
	SourceLSI
	SourceMSI
	SourceHSI48
)

// Expose all the clocks from this package for convenience of consumption
var (
	PCLK1  = &BaseClock{ClockTimerMultiplier: 1}
	PCLK2  = &BaseClock{ClockTimerMultiplier: 1}
	SYSCLK = &BaseClock{ClockTimerMultiplier: 1}
	HCLK   = &BaseClock{ClockTimerMultiplier: 1}
	HSI    = hsi.HSI
	LSI    = lsi.LSI
	LSE    = lse.LSE
	MSI    = msi.MSI
	HSI48  = hsi48.HSI48
)

// Lookup table to convert from symbolic clock source to the actual clock
// instance
var clocks = map[Source]PeripheralClock{
	SourcePCLK1:  PCLK1,
	SourcePCLK2:  PCLK2,
	SourceSYSCLK: SYSCLK,
	SourceHSE:    HSE,
	SourceHSI:    HSI,
	SourceLSE:    LSE,
	SourceLSI:    LSI,
	SourceMSI:    MSI,
	SourceHSI48:  HSI48,
}
