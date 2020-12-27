package msi

import "device/stm32"

// State represents the state of the High Speed Internal oscillator
type State uint32

const (
	// StateOff indicates the oscillator is off
	StateOff State = 0x00000000

	// StateOn indicates the oscillator is on
	StateOn State = stm32.RCC_CR_HSION
)

const (
	// CalibrationValueDefault is the default HSI calibration value
	CalibrationValueDefault = 0x0
)

var (
	// MSI gives public access to the oscillator
	MSI = &Oscillator{}
)

// Config is the configuration of the LSE oscillator
type Config struct {
	State State

	// CalibrationValue indicates the calibration value of the High Speed Internal oscillator
	CalibrationValue uint32

	// ClockRange indicates the range for the MSI clock
	ClockRange uint32
}

// Oscillator represents an HSE oscillator
type Oscillator struct {
	Attributes Attributes

	ClockFrequency int64
}

// Configure modifies the HSI state, waiting for completion
func (o *Oscillator) Configure(cfg *Config) {
	if cfg.State == StateOn {
		stm32.RCC.CR.SetBits(stm32.RCC_CR_MSION)

		// Turning on, wait for ready
		for !stm32.RCC.CR.HasBits(stm32.RCC_CR_MSIRDY) {
		}

		stm32.RCC.CR.Set(stm32.RCC_CR_MSIRGSEL)
		stm32.RCC.CR.ReplaceBits(cfg.ClockRange<<stm32.RCC_CR_MSIRANGE_Pos, stm32.RCC_CR_MSIRANGE_Msk, 0)
	} else {
		stm32.RCC.CR.ClearBits(stm32.RCC_CR_MSION)

		// Turning off, wait for not ready
		for stm32.RCC.CR.HasBits(stm32.RCC_CR_MSIRDY) {
		}
	}
}

func (o *Oscillator) Frequency() int64 {
	return o.Frequency()
}

func (o *Oscillator) TimerMultiplier() uint32 {
	return 1
}
