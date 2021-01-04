package clock

var (
	SYSCLK = BaseClock{ClockTimerMultiplier: 1}
	HCLK   = BaseClock{ClockTimerMultiplier: 1}
	PCLK1  = BaseClock{ClockTimerMultiplier: 1}
	PCLK2  = BaseClock{ClockTimerMultiplier: 1}
)

type Source uint8

// PeripheralClock represents the clock that drives a peripheral
type PeripheralClock interface {
	// The base clock frequency
	Frequency() int64

	// The multipler that applies when driving TIMx peripherals
	TimerMultiplier() uint32
}

// BaseClock can be used when there is no target-specific clock
// implementation
type BaseClock struct {
	ClockFrequency       int64
	ClockTimerMultiplier uint32
}

func (c *BaseClock) Frequency() int64 {
	return c.ClockFrequency
}

func (c *BaseClock) TimerMultiplier() uint32 {
	return c.ClockTimerMultiplier
}

// FromType returns the instance of a clock given it's type
//
// Return value is `nil` if the clock type is unknown / none.
func FromType(t Source) PeripheralClock {
	if c, ok := clocks[t]; ok {
		return c
	}

	return nil
}
