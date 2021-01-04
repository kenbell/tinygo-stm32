// +build stm32f10x

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
	SourceHSE
	SourceHSI
	SourceLSE
	SourceLSI
)

var (
	HSI = &hsi.HSI
	LSI = &lsi.LSI
	LSE = &lse.LSE
)

var clocks = map[Source]PeripheralClock{
	SourcePCLK1:  &PCLK1,
	SourcePCLK2:  &PCLK2,
	SourceSYSCLK: &SYSCLK,
	SourceHSE:    HSE,
	SourceHSI:    HSI,
	SourceLSE:    LSE,
	SourceLSI:    LSI,
}
