package timer

import (
	"device/stm32"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/nvic"
)

// Timer represents on-chip TIMx timers
type Timer struct {
	*stm32.TIM_Type
	Attributes *Attributes
	Clock      clock.PeripheralClock

	interruptHandler nvic.InterruptHandler
	NewInterrupt     func(nvic.InterruptHandler) *interrupt.Interrupt
}

// ConfigureBasic sets up the timer as a basic upwards counter
//
// This functionality is compatible with the 'Basic' timers found in
// STM32.
func (t *Timer) ConfigureBasic(cfg *Config) {
	// Enable the clock to this timer
	t.Attributes.Clock.EnableRegister.SetBits(t.Attributes.Clock.EnableFlag)

	attrs := t.Attributes
	cr1 := t.CR1.Get()

	if attrs.HasFeatures(FeatureCounterModelSelect) {
		cr1 &= ^uint32(stm32.TIM_CR1_DIR_Msk | stm32.TIM_CR1_CMS_Msk)
		cr1 |= uint32(cfg.CounterMode)
	}

	if attrs.HasFeatures(FeatureClockDivision) {
		cr1 &= ^uint32(stm32.TIM_CR1_CKD_Msk)
		cr1 |= uint32(cfg.ClockDivision)
	}

	cr1 &= ^uint32(stm32.TIM_CR1_ARPE)
	cr1 |= uint32(cfg.AutoReloadPreload)

	t.CR1.Set(cr1)
	t.ARR.Set(cfg.Period)
	t.PSC.Set(cfg.Prescaler)

	if attrs.HasFeatures(FeatureRepetitionCounter) {
		t.RCR.Set(cfg.RepetitionCounter)
	}

	if cfg.Repeat {
		t.EGR.SetBits(stm32.TIM_EGR_UG)
	} else {
		t.EGR.ClearBits(stm32.TIM_EGR_UG)
	}
}

// StartWithInterrupts starts the timer, with update interrupts enabled
//
// The caller is responsible for registering a handler on the corresponding
// interrupt.
func (t *Timer) StartWithInterrupts() {
	// Clear update flag
	t.SR.ClearBits(stm32.TIM_SR_UIF)

	// Enable the hardware interrupt
	t.DIER.SetBits(stm32.TIM_DIER_UIE)

	// Enable the hardware
	t.CR1.SetBits(stm32.TIM_CR1_CEN)
}

// Stop the timer
func (t *Timer) Stop() {
	// Disable the timer
	t.CR1.ClearBits(stm32.TIM_CR1_CEN)
}

// GetAndClearUpdateFlag indicates if the Update Interrupt Flag is set
//
// The result is true if the flag was set.  In which case the flag is
// also reset.
func (t *Timer) GetAndClearUpdateFlag() bool {
	if !t.SR.HasBits(stm32.TIM_SR_UIF) {
		return false
	}

	t.SR.ClearBits(stm32.TIM_SR_UIF)
	return true
}
