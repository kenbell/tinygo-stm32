package clock

import "runtime/volatile"

type PeripheralConfig struct {
	Default        Source
	EnableRegister *volatile.Register32
	EnableFlag     uint32
	SelectRegister *volatile.Register32
	SelectValues   map[Source]uint32
	SelectMask     uint32
}

func (c *PeripheralConfig) Apply(src Source) PeripheralClock {

	// Enable clock
	c.EnableRegister.SetBits(c.EnableFlag)

	if src == SourceNone {
		src = c.Default
	}

	// If no source provided, and no default the source is non-configurable
	// (e.g. like Timers), so return
	if src == SourceNone {
		return nil
	}

	// Select the desired source clock
	if c.SelectRegister != nil && len(c.SelectValues) > 0 {
		val := c.selectValue(src)
		reg := c.SelectRegister
		msk := c.SelectMask
		reg.ReplaceBits(val, msk, 0)
	}

	// Return the clock object
	return FromType(src)
}

func (c *PeripheralConfig) selectValue(t Source) uint32 {
	if c, ok := c.SelectValues[t]; ok {
		return c
	}

	return 0
}
