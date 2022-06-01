package engine

import (
	"infantry/bindings"
	"time"
)

func Run(plan bindings.Plan, output string) {
	var report bindings.Report
	report.Summary.StartTimestamp = time.Now().UTC().String()
}
