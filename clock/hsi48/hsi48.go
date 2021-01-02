package hsi48

import "device/stm32"

// State represents the state of the High Speed Internal oscillator
type State uint32

const (
	// StateOff indicates the oscillator is off
	StateOff State = 0x00000000

	// StateOn indicates the oscillator is on
	StateOn State = stm32.RCC_CRRCR_HSI48ON
)

// Config is the oscillator config
type Config struct {
	State State
}

// Oscillator represents the HSI48 oscillator
type Oscillator struct {
	Attributes Attributes
}

var (
	// HSI48 gives public access to the oscillator
	HSI48 = Oscillator{}
)

// Configure the oscillator,, waiting for completion
func (o *Oscillator) Configure(cfg *Config) {
	if cfg.State == StateOn {
		stm32.RCC.CRRCR.SetBits(stm32.RCC_CRRCR_HSI48ON)

		// Turning on, wait for ready
		for !stm32.RCC.CRRCR.HasBits(stm32.RCC_CRRCR_HSI48RDY) {
		}
	} else {
		stm32.RCC.CRRCR.ClearBits(stm32.RCC_CRRCR_HSI48ON)

		// Turning off, wait for not ready
		for stm32.RCC.CRRCR.HasBits(stm32.RCC_CRRCR_HSI48RDY) {
		}
	}
}

// Frequency of HSI48 is always 48 MHz
func (o *Oscillator) Frequency() int64 {
	return 48000000 // 48 MHz
}

// TimerMultiplier is always 1 for HSI48
func (o *Oscillator) TimerMultiplier() uint32 {
	return 1
}
