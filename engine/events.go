package engine

import (
	"github.com/gookit/event"
	"infantry/bindings"
)

// https://github.com/gookit/event

// TODO: Setup actual structs for messages with required data

func FirePlanStartEvent(message interface{}) {
	event.MustFire(bindings.PlanStartedEvent, event.M{"argo0": message})
}

func FirePlanCompleteEvent(message interface{}) {
	event.MustFire(bindings.PlanCompletedEvent, event.M{"argo0": message})
}

func FireStageStartedEvent(message interface{}) {
	event.MustFire(bindings.StageStartedEvent, event.M{"argo0": message})
}

func FireStageCompletedEvent(message interface{}) {
	event.MustFire(bindings.StageCompletedEvent, event.M{"argo0": message})
}

func FireProposalStartedEvent(message interface{}) {
	event.MustFire(bindings.ProposalStartedEvent, event.M{"argo0": message})
}

func FireProposalCompletedEvent(message interface{}) {
	event.MustFire(bindings.ProposalCompletedEvent, event.M{"argo0": message})
}

func FireProposalTaskStartedEvent(message interface{}) {
	event.MustFire(bindings.ProposalTaskStartedEvent, event.M{"argo0": message})
}

func FireProposalTaskSuccessEvent(message interface{}) {
	event.MustFire(bindings.ProposalTaskSuccessEvent, event.M{"argo0": message})
}

func FireProposalTaskFailureEvent(message interface{}) {
	event.MustFire(bindings.ProposalTaskFailureEvent, event.M{"argo0": message})
}

func FireCreatingReportEvent(message interface{}) {
	event.MustFire(bindings.CreatingReportEvent, event.M{"argo0": message})
}
