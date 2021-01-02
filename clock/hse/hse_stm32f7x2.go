package hse

var (
	// HSE gives public access to the oscillator
	HSE = Oscillator{
		Attributes: Attributes{
			DefaultFrequency: 8000000, // 8 MHz
		},
		ClockFrequency: 8000000, // 8 MHz
	}
)
