package bindings

type DataEvent struct {
	Payload   interface{}
	EventType int
}

const (
	StartedEvent = iota
	CompletedEvent
	IterationEvent
	SuccessEvent
	FailureEvent
)
