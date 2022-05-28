package engine

import (
	"infantry/bindings"
	"time"
)

func main(plan bindings.Plan, output string) {
	var report bindings.Report
	report.Summary.StartTimestamp = time.Now().UTC().String()
}
