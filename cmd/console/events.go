package console

import (
	"fmt"
	"github.com/gookit/event"
	"infantry/bindings"
	"time"
)

func SetupEventListeners() {
	event.On(bindings.PlanStartedEvent, event.ListenerFunc(func(e event.Event) error {
		TimestampLog("Started Plan...")
		return nil
	}), event.High)

	event.On(bindings.PlanCompletedEvent, event.ListenerFunc(func(e event.Event) error {
		TimestampLog("Complete!")
		return nil
	}), event.High)

	event.On(bindings.ProposalTaskStartedEvent, event.ListenerFunc(func(e event.Event) error {
		TimestampLog("-- Task Started...")
		return nil
	}), event.High)
}

func TimestampLog(message string) {
	fmt.Println(time.Now().Format("[15:04:05-06] ") + message)
}
