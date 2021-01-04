// +build stm32f10x

package hsi

var (
	// HSI gives public access to the oscillator
	HSI = Oscillator{
		Attributes: Attributes{
			DefaultFrequency: 8000000, // 8 MHz
		},
		ClockFrequency: 8000000,
	}
)
