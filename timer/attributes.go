package timer

import "github.com/kenbell/tinygo-stm32/clock"

type Features uint32

const (
	FeatureNone               Features = 0
	FeatureCounterModelSelect Features = 1 << iota
	FeatureClockDivision
	FeatureRepetitionCounter
)

// Attributes is an extension point to allow shared behaviour
// to be customized for particular targets and/or instances
type Attributes struct {
	Features Features
	Clock    clock.PeripheralConfig
}

func (a Attributes) HasFeatures(f Features) bool {
	return a.Features&f == f
}
