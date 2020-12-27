package port

import (
	"device/stm32"
	"runtime/volatile"
)

type Port struct {
	*stm32.GPIO_Type
	ClockEnableRegister *volatile.Register32
	ClockEnableFlag     uint32
	PowerEnableRegister *volatile.Register32
	PowerEnableFlag     uint32
}

func (p *Port) SetAltFunc(pin uint8, af uint8) {
	pos := (pin % 8) * 4

	if pin < 8 {
		p.AFRL.ReplaceBits(uint32(af), 0xf, pos)
	} else {
		p.AFRH.ReplaceBits(uint32(af), 0xf, pos)
	}
}

func (p *Port) EnableClock() {
	// On some chips, some ports must be explicitly powered
	if p.PowerEnableRegister != nil {
		p.PowerEnableRegister.SetBits(p.PowerEnableFlag)
	}

	p.ClockEnableRegister.SetBits(p.ClockEnableFlag)
}
