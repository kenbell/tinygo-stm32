package nvic

import (
	"device/arm"
	"runtime/interrupt"
)

type InterruptHandler func(interrupt.Interrupt)

type HandlerRecord struct {
	handler InterruptHandler
	next    *HandlerRecord
}

type HandlerChain struct {
	*HandlerRecord
}

func (ch *HandlerChain) Register(handler InterruptHandler) bool {
	state := arm.DisableInterrupts()

	record := &HandlerRecord{
		handler: handler,
		next:    ch.HandlerRecord,
	}
	ch.HandlerRecord = record

	arm.EnableInterrupts(state)

	return record.next != nil
}

func (ch *HandlerChain) Call(intr interrupt.Interrupt) {
	h := ch.HandlerRecord
	for h != nil {
		if h.handler != nil {
			h.handler(intr)
		}

		h = h.next
	}
}
