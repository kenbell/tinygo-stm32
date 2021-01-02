// +build stm32l5x2

package hsi

const (
	// CalibrationValueDefault is the default HSI calibration value
	CalibrationValueDefault = 0x10

	// CalibrationValueMax is the maximum calibration value
	CalibrationValueMax = 0x1F
)

var (
	// HSI gives public access to the oscillator
	HSI = Oscillator{
		Attributes: Attributes{
			DefaultFrequency: 16000000, // 16 MHz
		},
		ClockFrequency: 16000000,
	}
)
