package engine

import (
	"encoding/json"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"infantry/bindings"
	"io/ioutil"
)

// SaveReportFile Saves the overall output to a json file for reporting to read
func SaveReportFile(data bindings.Report, fileName string) {
	file, _ := json.MarshalIndent(data, "", "  ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

// ReadPlanFile Reads in a plan yaml file
func ReadPlanFile(path string) (*bindings.Plan, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(bindings.Plan)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}
