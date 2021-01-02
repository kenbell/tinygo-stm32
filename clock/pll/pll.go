package pll

// State indicates the PLL state
type State uint32

const (
	// StateNone indicates the state is unspecified
	StateNone State = 0x00000000

	// StateOff indicates the PLL state is off
	StateOff State = 0x00000001

	// StateOn indicates the PLL state is on
	StateOn State = 0x00000002
)

type PLL struct {
	ClockFrequency int64
}

func (pll *PLL) Configure(cfg *Config) {
	pll.ClockFrequency = cfg.Apply()
}

// Frequency gets the PLLCLK frequency, which is selectable as an input to be
// the SYSCLK frequency.
func (pll *PLL) Frequency() int64 {
	return pll.ClockFrequency
}

// TimerMultiplier is always 1 for PLL since not used directly by timers
func (pll *PLL) TimerMultiplier() uint32 {
	return 1
}
