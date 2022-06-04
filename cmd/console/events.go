package console

import (
	"fmt"
	"github.com/gookit/event"
	"infantry/bindings"
)

func SetupEventListeners() {
	event.On(bindings.StartedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Started Plan...")
		return nil
	}), event.High)

	event.On(bindings.CompletedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Complete!")
		return nil
	}), event.High)
}
