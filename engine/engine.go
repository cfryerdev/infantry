package engine

import (
	"infantry/bindings"
)

// https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

var _report bindings.Report
var _plan bindings.Plan

func Start(plan bindings.Plan) (bindings.Report, error) {
	SetupReport(_report)
	FirePlanStartEvent(nil)
	for _, stage := range plan.Setup.Stages {
		ExecuteStage(stage, plan.Proposals)
	}
	CreateSummary(_report)
	FinalizeReport(_report)
	FirePlanCompleteEvent(nil)
	return _report, nil
}

func ExecuteStage(stage bindings.Stage, proposals []bindings.Proposal) {
	FireStageStartedEvent(nil)
	ExecuteEachSecondConcurrent(stage.Iterations, func() {
		ExecuteUserProposals(proposals, stage.AddUsersPerIterations, stage.MaxUsersAtOnce)
	})
	FireStageCompletedEvent(nil)
}

func ExecuteUserProposals(proposals []bindings.Proposal, addUsers int, maxUsers int) {
	FireProposalStartedEvent(nil)
	ExecuteNumberOfTimesConcurrent(addUsers, maxUsers, func() {
		ExecuteTasks(User{}, proposals)
	})
	FireProposalCompletedEvent(nil)
}

func ExecuteTasks(user User, proposals []bindings.Proposal) {
	for _, proposal := range proposals {
		FireProposalTaskStartedEvent(proposal)
		var resp, err = user.ExecuteProposal(_plan.Protocol, proposal)
		if err != nil {
			FireProposalTaskFailureEvent(resp)
		}
		//else { FireProposalTaskSuccessEvent(resp) }
	}
}
