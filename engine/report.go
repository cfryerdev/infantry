package engine

import (
	"github.com/google/uuid"
	"infantry/bindings"
	"time"
)

func SetupReport(report bindings.Report) {
	report.Id = uuid.New()
	report.Summary.StartTimestamp = time.Now().UTC().String()
}

func FinalizeReport(report bindings.Report) {
	FireCreatingReportEvent(nil)
}

func CreateSummary(report bindings.Report) {
	report.Summary.EndTimestamp = time.Now().UTC().String()
}
