package console

import (
	"fmt"
	"github.com/gookit/event"
	"infantry/bindings"
)

func SetupEventListeners() {
	event.On(bindings.PlanStartedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Started Plan...")
		return nil
	}), event.High)

	event.On(bindings.PlanCompletedEvent, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Complete!")
		return nil
	}), event.High)
}
