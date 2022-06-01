package engine

import (
	"fmt"
	"infantry/bindings"
	"time"
)

func Run(plan bindings.Plan, output string) {
	var report bindings.Report
	report.Summary.StartTimestamp = time.Now().UTC().String()

	fmt.Printf("DEBUG: %+v\n", plan)

}
