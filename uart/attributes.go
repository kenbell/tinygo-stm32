package uart

import (
	"runtime/volatile"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/port"
)

// Attributes is an extension point to allow shared behaviour
// to be customized for particular targets and/or instances
type Attributes struct {
	TxPort       *port.Port
	TxPin        uint8
	TxPinAltFunc uint8
	TxReg        *volatile.Register32

	RxPort       *port.Port
	RxPin        uint8
	RxPinAltFunc uint8
	RxReg        *volatile.Register32

	// Some chips have a fixed multiplier on the baud rate configured
	// into the BRR register.  Default is multiplier of 1)
	BaudMultiplier uint32

	StatusReg         *volatile.Register32
	StatusTxEmptyFlag uint32

	Clock clock.PeripheralConfig
}
