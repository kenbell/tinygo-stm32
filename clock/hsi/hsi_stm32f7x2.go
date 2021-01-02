// +build stm32f7x2

package hsi

var (
	// HSI gives public access to the oscillator
	HSI = Oscillator{
		Attributes: Attributes{
			DefaultFrequency: 16000000, // 16 MHz
		},
		ClockFrequency: 16000000,
	}
)
