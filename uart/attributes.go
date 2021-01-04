package uart

import (
	"device/stm32"
	"machine"
	"runtime/volatile"

	"github.com/kenbell/tinygo-stm32/clock"
)

type PinMapping struct {
	TxPin        machine.Pin
	TxPinAltFunc stm32.AltFunc
	RxPin        machine.Pin
	RxPinAltFunc stm32.AltFunc
}

// Attributes is an extension point to allow shared behaviour
// to be customized for particular targets and/or instances
type Attributes struct {
	PinMapping  []PinMapping
	TxPinConfig machine.PinConfig
	TxReg       *volatile.Register32

	RxPinConfig machine.PinConfig
	RxReg       *volatile.Register32

	// Some chips have a fixed multiplier on the baud rate configured
	// into the BRR register.  Default is multiplier of 1)
	BaudMultiplier uint32

	StatusReg         *volatile.Register32
	StatusTxEmptyFlag uint32

	Clock clock.PeripheralConfig
}
