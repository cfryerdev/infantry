package engine

import (
	"fmt"
	"infantry/bindings"
)

// https://www.geeksforgeeks.org/function-as-a-field-in-golang-structure/

var report bindings.Report

func Start(plan bindings.Plan) (bindings.Report, error) {
	SetupReport(report)
	FirePlanStartEvent(nil)
	for _, stage := range plan.Setup.Stages {
		ExecuteStage(stage, plan.Proposals)
	}
	CreateSummary(report)
	FinalizeReport(report)
	FirePlanCompleteEvent(nil)
	return report, nil
}

func ExecuteStage(stage bindings.Stage, proposals []bindings.Proposal) {
	FireStageStartedEvent(nil)
	ExecuteEachSecondConcurrent(stage.Seconds, func() {
		ExecuteUserProposals(proposals, stage.AddUsersPerSecond, stage.MaxUsersAtOnce)
	})
	FireStageCompletedEvent(nil)
}

func ExecuteUserProposals(proposals []bindings.Proposal, addUsers int, maxUsers int) {
	ExecuteNumberOfTimesConcurrent(addUsers, maxUsers, func() {
		FireProposalStartedEvent(nil)
		var user = CreateVirtualUser()
		ExecuteTask(user, proposals)
		FireProposalCompletedEvent(nil)
	})
}

func ExecuteTask(user User, proposals []bindings.Proposal) {
	for _, proposal := range proposals {
		FireProposalTaskStartedEvent(nil)
		fmt.Printf("PROPOSAL: %+v\n", proposal)
		// if err != nill { FireProposalTaskFailureEvent(nil) }
		// else { FireProposalTaskSuccessEvent(nil) }
	}
}
