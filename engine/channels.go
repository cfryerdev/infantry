package engine

import (
	"github.com/gookit/event"
	"infantry/bindings"
)

func startEvent() {
	event.MustFire(bindings.StartedEvent, event.M{})
}

func completeEvent() {
	event.MustFire(bindings.CompletedEvent, event.M{})
}
