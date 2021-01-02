package lse

import "device/stm32"

// State represents the state of the Low Speed External oscillator
type State uint32

const (
	// StateOff indicates the oscillator is off
	StateOff State = 0x00000000

	// StateOn indicates a crystal is used as the source
	StateOn State = stm32.RCC_BDCR_LSEON

	// StateBypass indicates an oscillator is used as the source
	StateBypass State = stm32.RCC_BDCR_LSEBYP | stm32.RCC_BDCR_LSEON
)

// Config is the configuration of the LSE oscillator
type Config struct {
	State     State
	Frequency int64
}

// Oscillator represents an LSE oscillator
type Oscillator struct {
	Attributes Attributes

	ClockFrequency int64
}

// Configure modifies the LSE state, waiting for completion
func (o *Oscillator) Configure(cfg *Config) {
	// Use default frequency if not overridden
	if cfg.Frequency != 0 {
		o.ClockFrequency = cfg.Frequency
	} else {
		o.ClockFrequency = o.Attributes.DefaultFrequency
	}

	// Apply power to peripheral clock, if needed
	powerStateOff := (o.Attributes.ClockEnableRegister.Get() & o.Attributes.ClockEnableFlag) == 0
	if powerStateOff {
		o.Attributes.ClockEnableRegister.SetBits(o.Attributes.ClockEnableFlag)
		_ = o.Attributes.ClockEnableRegister.Get()
	}

	// Disable Backup domain protection
	if !stm32.PWR.CR1.HasBits(stm32.PWR_CR1_DBP) {
		stm32.PWR.CR1.SetBits(stm32.PWR_CR1_DBP)
		for !stm32.PWR.CR1.HasBits(stm32.PWR_CR1_DBP) {
		}
	}

	// Set LSE Drive to LOW
	//stm32.RCC.BDCR.ReplaceBits(0, stm32.RCC_BDCR_LSEDRV_Msk, 0)

	waitReady := true

	switch cfg.State {
	case StateOn:
		stm32.RCC.BDCR.SetBits(stm32.RCC_BDCR_LSEON)
	case StateBypass:
		stm32.RCC.BDCR.SetBits(stm32.RCC_BDCR_LSEON)
		stm32.RCC.BDCR.SetBits(stm32.RCC_BDCR_LSEBYP)
	default:
		stm32.RCC.BDCR.ClearBits(stm32.RCC_BDCR_LSEON | stm32.RCC_BDCR_LSEBYP)
		waitReady = false
	}

	// Turning on, wait for ready - turning off, wait for not ready
	for stm32.RCC.BDCR.HasBits(stm32.RCC_BDCR_LSERDY) != waitReady {
	}

	//stm32.RCC.BDCR.ClearBits(stm32.RCC_BDCR_LSESYSEN)
	//for stm32.RCC.BDCR.HasBits(stm32.RCC_BDCR_LSESYSEN) {
	//}

	// Power down peripheral clock again, if was not already powered up
	if powerStateOff {
		o.Attributes.ClockEnableRegister.ClearBits(o.Attributes.ClockEnableFlag)
	}
}

// Frequency gets the configured frequency of the LSE
func (o *Oscillator) Frequency() int64 {
	return o.ClockFrequency
}

// TimerMultiplier for LSE is always 1
func (o *Oscillator) TimerMultiplier() uint32 {
	return 1
}
