package lsi

var (
	// LSI gives public access to the oscillator
	LSI = Oscillator{
		Attributes: Attributes{
			DefaultFrequency: 32000, // 32 KHz
		},
	}
)
