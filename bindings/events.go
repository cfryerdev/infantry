package bindings

type DataEvent struct {
	Payload   interface{}
	EventType int
}

const (
	StartedEvent           = "started_event"
	CompletedEvent         = "completed_event"
	IterationStartEvent    = "iteration_start_event"
	IterationCompleteEvent = "iteration_completed_event"
	SuccessEvent           = "success_event"
	FailureEvent           = "failure_event"
)
