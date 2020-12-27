package timer

import "github.com/kenbell/tinygo-stm32/clock"

type TIMCounterMode uint32
type TIMClockDivision uint32
type TIMAutoReloadPreload uint32

type Config struct {
	Prescaler         uint32
	CounterMode       TIMCounterMode
	Period            uint32
	ClockDivision     TIMClockDivision
	RepetitionCounter uint32
	AutoReloadPreload TIMAutoReloadPreload
	Repeat            bool
}

// SetFrequency is a convenience function to calculate an appropriate
// prescaler and period combination to achieve a desired frequency
// for a timer using the indicated peripheral clock
func (c *Config) SetFrequency(hz int64, clk clock.PeripheralClock) {
	psc := uint32((clk.Frequency() * int64(clk.TimerMultiplier())) / hz)
	period := uint32(1)

	// Get the pre-scale into range, with interrupt firing
	// once per tick.
	for psc > 0x10000 || period == 1 {
		psc >>= 1
		period <<= 1
	}

	// Clamp overflow
	if period > 0x10000 {
		period = 0x10000
	}

	// The actual values for the timer are offset by one
	c.Prescaler = psc - 1
	c.Period = period - 1
	c.Repeat = true
}

// SetDelay is a convenience function to calculate an appropriate prescaler
// and period combination to achieve a desired delay for a timer using
// the indicated peripheral lock.
//
// At high clock speeds, the maximum delay is limited by the processor registers
// to 2^32 clock cycles (e.g. for a 50MHz peripheral clock approx 40 secs).  On
// overflow the delay is clamped to the maximum delay configurable (i.e. timer
// will fire early).  For long delays, either use an slower clock speed and/or
// repeat the delay as many times as needed.
func (c *Config) SetDelay(ns int64, clk clock.PeripheralClock) {
	// Calculate initial pre-scale value.
	// delay (in ns) and clock freq are both large values, so do the nanosecs
	// conversion (divide by 1G) by pre-dividing each by 1000 to avoid overflow
	// in any meaningful time period.  This reduces max theoretical resolution
	// to approx 1 us.
	psc := ((ns / 1000) * ((clk.Frequency() * int64(clk.TimerMultiplier())) / 1000) / 1000)
	period := int64(1)

	// Get the pre-scale into range, with interrupt firing
	// once per tick.
	for psc > 0x10000 || period == 1 {
		psc >>= 1
		period <<= 1
	}

	// Clamp overflow
	if period > 0x10000 {
		period = 0x10000
	}

	// We use a basic timer set to not reload when period expires
	// effectively giving us a single-shot timer.
	c.Prescaler = uint32(psc) - 1
	c.Period = uint32(period) - 1
	c.Repeat = false
}
