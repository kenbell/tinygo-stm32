package hse

import "device/stm32"

// State represents the state of the High Speed External oscillator
type State uint32

const (
	// StateOff indicates the oscillator is off
	StateOff State = 0x00000000

	// StateOn indicates a crystal is used as the source
	StateOn State = stm32.RCC_CR_HSEON

	// StateBypass indicates an oscillator is used as the source
	StateBypass State = stm32.RCC_CR_HSEBYP | stm32.RCC_CR_HSEON
)

var (
	// HSE gives public access to the oscillator
	HSE = &Oscillator{}
)

// Config is the configuration of the LSE oscillator
type Config struct {
	State     State
	Frequency int64
}

// Oscillator represents an HSE oscillator
type Oscillator struct {
	Attributes Attributes

	ClockFrequency int64
}

// Configure modifies the HSE state, waiting for completion
func (o *Oscillator) Configure(cfg *Config) {

	if cfg.Frequency != 0 {
		o.ClockFrequency = cfg.Frequency
	}

	waitReady := true

	switch cfg.State {
	case StateOn:
		stm32.RCC.CR.SetBits(stm32.RCC_CR_HSEON)
	case StateBypass:
		stm32.RCC.CR.SetBits(stm32.RCC_CR_HSEON)
		stm32.RCC.CR.SetBits(stm32.RCC_CR_HSEBYP)
	default:
		stm32.RCC.CR.ClearBits(stm32.RCC_CR_HSEON | stm32.RCC_CR_HSEBYP)
		waitReady = false
	}

	// Turning on, wait for ready - turning off, wait for not ready
	for stm32.RCC.CR.HasBits(stm32.RCC_CR_HSERDY) != waitReady {
	}
}

func (o *Oscillator) Frequency() int64 {
	return o.Frequency()
}

func (o *Oscillator) TimerMultiplier() uint32 {
	return 1
}
