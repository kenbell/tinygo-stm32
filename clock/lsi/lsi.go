package lsi

import "device/stm32"

// State represents the state of the Low Speed Internal oscillator
type State uint32

const (
	// StateOff indicates the oscillator is off
	StateOff State = 0x00000000

	// StateOn indicates the oscillator is on
	StateOn State = stm32.RCC_CSR_LSION
)

// Config is the configuration of the LSE oscillator
type Config struct {
	State State
}

// Oscillator represents an LSE oscillator
type Oscillator struct {
	Attributes Attributes

	ClockFrequency int64
}

// Configure modifies the LSI state, waiting for completion
func (o *Oscillator) Configure(cfg *Config) {

	o.ClockFrequency = o.Attributes.DefaultFrequency

	if cfg.State == StateOn {
		stm32.RCC.CSR.SetBits(stm32.RCC_CSR_LSION)

		// Turning on, wait for ready
		for !stm32.RCC.CSR.HasBits(stm32.RCC_CSR_LSIRDY) {
		}
	} else {
		stm32.RCC.CSR.ClearBits(stm32.RCC_CSR_LSION)

		// Turning off, wait for not ready
		for stm32.RCC.CSR.HasBits(stm32.RCC_CSR_LSIRDY) {
		}
	}
}

func (o *Oscillator) Frequency() int64 {
	return o.ClockFrequency
}

func (o *Oscillator) TimerMultiplier() uint32 {
	return 1
}
