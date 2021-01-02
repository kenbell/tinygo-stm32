package msi

const (
	khz = 1000
	mhz = 1000000
)

var freqTable = []int64{
	100 * khz,
	200 * khz,
	400 * khz,
	800 * khz,
	1 * mhz,
	2 * mhz,
	4 * mhz,
	8 * mhz,
	16 * mhz,
	24 * mhz,
	32 * mhz,
	48 * mhz,
}

func rangeToFreq(r uint32) int64 {
	return freqTable[r]
}
