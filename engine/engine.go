package engine

import (
	"infantry/bindings"
	"infantry/protocols/http"
	"reflect"
)

// https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

var _report bindings.Report
var _plan bindings.Plan
var _executor bindings.Executor

func Start(plan bindings.Plan) (bindings.Report, error) {
	SetupReport(_report)
	FirePlanStartEvent(nil)
	CreateExecutor(_plan.Protocol)
	for _, stage := range plan.Setup.Stages {
		ExecuteStage(stage, plan.Proposals)
	}
	CreateSummary(_report)
	FinalizeReport(_report)
	FirePlanCompleteEvent(nil)
	return _report, nil
}

func CreateExecutor(protocol bindings.Protocol) {
	//TODO - verify this detects protocol
	if (reflect.DeepEqual(bindings.ProtocolHttp{}, protocol.Http)) {
		_executor = http.HttpExecutor{}
	}

	_executor.Initialize(protocol)
}

func ExecuteStage(stage bindings.Stage, proposals []bindings.Proposal) {
	FireStageStartedEvent(nil)

	var currentIteration = 0
	var currentUsers = stage.AddUsersPerIterations
	ExecuteEachSecondConcurrent(stage.Iterations, func() {
		if currentIteration > stage.Iterations {
			return
		}

		ExecuteUserProposals(proposals, currentUsers)

		var usersToAdd = stage.MaxUsersAtOnce - currentUsers
		if usersToAdd > 0 {
			currentUsers += usersToAdd
		}

		currentIteration++
	})

	FireStageCompletedEvent(nil)
}

func ExecuteUserProposals(proposals []bindings.Proposal, currentUsers int) {
	FireProposalStartedEvent(nil)
	ExecuteNumberOfTimesConcurrent(currentUsers, func() {
		ExecuteTasks(User{}, proposals)
	})
	FireProposalCompletedEvent(nil)
}

func ExecuteTasks(user User, proposals []bindings.Proposal) {
	for _, proposal := range proposals {
		FireProposalTaskStartedEvent(proposal)
		var resp, err = _executor.Execute((proposal))

		if err != nil {
			FireProposalTaskFailureEvent(resp)
			return
		}

		FireProposalTaskSuccessEvent(resp)
	}
}
