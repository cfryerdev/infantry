package engine

import (
	"fmt"
	"github.com/google/uuid"
	"infantry/bindings"
	"time"
)

func Run(plan bindings.Plan) (bindings.Report, error) {
	var report bindings.Report
	report.Id = uuid.New()
	report.Summary.StartTimestamp = time.Now().UTC().String()

	fmt.Printf("DEBUG: %+v\n", plan)

	return bindings.Report{}, nil
}
