package engine

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gookit/event"
	"infantry/bindings"
	"io/ioutil"
	"time"
)

func SetupReport(report bindings.Report) {
	SetupReportEventListeners(report)
}

func FinalizeReport(report bindings.Report) {
	FireCreatingReportEvent(nil)
}

func CreateSummary(report bindings.Report) {
	report.Summary.EndTimestamp = time.Now().UTC().String()
}

// SaveReportFile Saves the overall output to a json file for reporting to read
func SaveReportFile(data bindings.Report, fileName string) {
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func SetupReportEventListeners(report bindings.Report) {
	event.On(bindings.PlanStartedEvent, event.ListenerFunc(func(e event.Event) error {
		report.Id = uuid.New()
		report.Summary.StartTimestamp = time.Now().UTC().String()
		return nil
	}), event.High)

	event.On(bindings.PlanCompletedEvent, event.ListenerFunc(func(e event.Event) error {
		return nil
	}), event.High)

	event.On(bindings.ProposalStartedEvent, event.ListenerFunc(func(e event.Event) error {
		return nil
	}), event.High)

	event.On(bindings.ProposalCompletedEvent, event.ListenerFunc(func(e event.Event) error {
		// SaveReportFile(report, "")
		return nil
	}), event.High)
}
