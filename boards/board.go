package boards

import (
	"device/arm"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/power"
	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

const tickDebug = false

const TICK_RATE = 1000 // 1 KHz
const TICKS_PER_NS = 1000000000 / TICK_RATE

var tickCount volatile.Register64
var timerWakeup volatile.Register8

// BasicBoard is an implementation of tinygo machine.GenericBoard for STM32 MCU.
//
// BasicBoard implements basic tick/sleep functionality.  The user is expected to
// provide the clock configuration.
type BasicBoard struct {
	tickTimer  *timer.Timer
	sleepTimer *timer.Timer
	uart       *uart.UART
	led        machine.Pin
}

type BasicBoardConfig struct {
	SleepTimer       *timer.Timer
	TickTimer        *timer.Timer
	UART             *uart.UART
	UARTConfig       *uart.Config
	OscillatorConfig *clock.OscillatorConfig
	ClockConfig      *clock.Config
	LED              machine.Pin
}

func NewBasicBoard(config *BasicBoardConfig) *BasicBoard {
	b := &BasicBoard{}

	b.initClocks(config.OscillatorConfig, config.ClockConfig)
	b.initTickTimer(config.TickTimer)
	b.initSleepTimer(config.SleepTimer)

	b.uart = config.UART
	b.uart.Configure(config.UARTConfig)

	b.led = config.LED

	return b
}

func (b *BasicBoard) initClocks(oscCfg *clock.OscillatorConfig, clkCfg *clock.Config) {
	power.EnableOverdrive()

	oscCfg.Apply()
	clkCfg.Apply(7)
}

func (b *BasicBoard) SleepTicks(d int64) {
	timerWakeup.Set(0)

	b.startSleepTimer(d)

	// wait till timer wakes up
	for timerWakeup.Get() == 0 {
		arm.Asm("wfi")
	}
}

func (b *BasicBoard) Ticks() int64 {
	return int64(tickCount.Get())
}

func (b *BasicBoard) TicksToNanoseconds(ticks int64) int64 {
	return ticks * TICKS_PER_NS
}

func (b *BasicBoard) NanosecondsToTicks(ns int64) int64 {
	return ns / TICKS_PER_NS
}

func (b *BasicBoard) UART() machine.GenericUART {
	return b.uart
}

func (b *BasicBoard) LED() machine.Pin {
	return b.led
}

func (b *BasicBoard) initSleepTimer(t *timer.Timer) {
	b.sleepTimer = t

	intr := b.sleepTimer.NewInterrupt(b.handleWakeup)
	intr.SetPriority(0xc3)
	intr.Enable()
}

func (b *BasicBoard) startSleepTimer(ticks int64) {
	cfg := timer.Config{}
	cfg.SetDelay(b.TicksToNanoseconds(ticks), b.sleepTimer.Clock)
	b.sleepTimer.ConfigureBasic(&cfg)
	b.sleepTimer.StartWithInterrupts()
}

func (b *BasicBoard) initTickTimer(t *timer.Timer) {
	b.tickTimer = t

	if tickDebug {
		machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// Repeating timer, with prescale and period calculated
	// from the tick rate
	cfg := timer.Config{}
	cfg.SetFrequency(TICK_RATE, b.tickTimer.Clock)

	b.tickTimer.ConfigureBasic(&cfg)

	intr := b.tickTimer.NewInterrupt(b.handleTick)
	intr.SetPriority(0xc1)
	intr.Enable()

	b.tickTimer.StartWithInterrupts()
}

func (b *BasicBoard) handleWakeup(interrupt.Interrupt) {
	if b.sleepTimer.GetAndClearUpdateFlag() {
		// Repeat is disable, but we also stop the timer when
		// not waiting
		b.sleepTimer.Stop()

		// timer was triggered
		timerWakeup.Set(1)
	}
}

var debugLEDState = false

func (b *BasicBoard) handleTick(interrupt.Interrupt) {
	if b.tickTimer.GetAndClearUpdateFlag() {
		c := tickCount.Get()

		// 1Hz LED flash
		if tickDebug && c%500 == 0 {
			debugLEDState = !debugLEDState
			machine.LED.Set(debugLEDState)
		}

		tickCount.Set(c + 1)
	}
}
