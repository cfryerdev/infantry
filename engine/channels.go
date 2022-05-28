package engine

import (
	"github.com/gookit/event"
	"infantry/bindings"
)

// https://github.com/gookit/event

func startEvent(message bindings.DataEvent) {
	event.MustFire(bindings.StartedEvent, event.M{"argo0": message})
}

func completeEvent(message bindings.DataEvent) {
	event.MustFire(bindings.CompletedEvent, event.M{"argo0": message})
}
