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
}

func (pll *PLL) Configure(cfg *Config) {
	cfg.Apply()
}
