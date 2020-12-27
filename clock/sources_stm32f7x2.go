// +build stm32f7x2

package clock

import (
	"github.com/kenbell/tinygo-stm32/clock/hsi"
	"github.com/kenbell/tinygo-stm32/clock/lse"
	"github.com/kenbell/tinygo-stm32/clock/lsi"
)

const (
	SourceNone Source = iota
	SourcePCLK1
	SourcePCLK2
	SourceSYSCLK
	SourceHSI
	SourceLSE
	SourceLSI
)

var (
	PCLK1  = &BaseClock{ClockTimerMultiplier: 2}
	PCLK2  = &BaseClock{ClockTimerMultiplier: 2}
	SYSCLK = &BaseClock{ClockTimerMultiplier: 1}
	HCLK   = &BaseClock{ClockTimerMultiplier: 1}
	HSI    = hsi.HSI
	LSI    = lsi.LSI
	LSE    = lse.LSE
)

var clocks = map[Source]PeripheralClock{
	SourcePCLK1:  PCLK1,
	SourcePCLK2:  PCLK2,
	SourceSYSCLK: SYSCLK,
	SourceHSI:    HSI,
	SourceLSE:    LSE,
	SourceLSI:    LSI,
}
