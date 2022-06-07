package engine

import (
	"github.com/google/uuid"
	"infantry/bindings"
	"sync"
	"time"
)

func Start(plan bindings.Plan) (bindings.Report, error) {
	var report bindings.Report
	report.Id = uuid.New()
	report.Summary.StartTimestamp = time.Now().UTC().String()
	for _, stage := range plan.Setup.Stages {
		ExecuteStage(stage, plan.Proposals)
	}
	return bindings.Report{}, nil
}

func ExecuteStage(stage bindings.Stage, proposals []bindings.Proposal) {
	ticker := time.NewTicker(time.Duration(stage.Seconds) * time.Second)
	quit := make(chan struct{})
	var guard = make(chan struct{}, stage.MaxUsersAtOnce)
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case <-ticker.C:
				ExecuteProposal(proposals, stage.AddUsersPerSecond, guard)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	wg.Done()
}

func ExecuteProposal(proposals []bindings.Proposal, addUsers int, guard chan struct{}) {
	var wg sync.WaitGroup
	for i := 1; i <= addUsers; i++ {
		go func(n int) {
			guard <- struct{}{}
			var user = CreateVirtualUser()
			ExecuteTask(user, proposals)
			<-guard
		}(i)
	}
	wg.Done()
}

func ExecuteTask(user User, proposals []bindings.Proposal) {

}
