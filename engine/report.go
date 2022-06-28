package engine

import (
	"encoding/json"
	"github.com/google/uuid"
	"infantry/bindings"
	"io/ioutil"
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

// SaveReportFile Saves the overall output to a json file for reporting to read
func SaveReportFile(data bindings.Report, fileName string) {
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}
