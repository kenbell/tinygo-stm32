package lse

import "runtime/volatile"

// Attributes is an extension point to allow shared behaviour
// to be customized for particular targets and/or instances
type Attributes struct {
	ClockEnableRegister         *volatile.Register32
	ClockEnableFlag             uint32
	DefaultFrequency            int64
	DisableBackupDomainRegister *volatile.Register32
	DisableBackupDomainFlag     uint32
}
