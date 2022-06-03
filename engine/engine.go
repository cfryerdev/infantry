package engine

import (
	"fmt"
	"infantry/bindings"
	"time"
)

func Run(plan bindings.Plan) bindings.Report {
	var report bindings.Report
	report.Summary.StartTimestamp = time.Now().UTC().String()

	fmt.Printf("DEBUG: %+v\n", plan)

	return bindings.Report{}
}
