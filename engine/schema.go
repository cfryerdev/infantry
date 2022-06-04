package engine

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"infantry/bindings"
	"io/ioutil"
)

// LoadPlanSchemaFromPath Reads the plan yaml file
func LoadPlanSchemaFromPath(pathToConfig string) bindings.Plan {
	cfgPath := flag.String("p", pathToConfig, "Path to yaml config file")
	flag.Parse()
	cfg, err := ReadPlanFile(*cfgPath)
	if err != nil {
		panic(err.Error())
	}
	return *cfg
}

// ValidateSchema Validates the plan file to ensure it will function correctly
func ValidateSchema(plan bindings.Plan) {
	var planValidator = validator.New()
	err := planValidator.Struct(plan)
	if err != nil {

	}
}

// LoadEnvironmentOverrides Replaces any instances where env vars are used
func LoadEnvironmentOverrides(plan bindings.Plan) bindings.Plan {
	// v := reflect.ValueOf(plan)
	// https://pkg.go.dev/text/template
	// loop through each property, and find any wildcards -> {{.%s}}
	// Get the env variable, then replace value and return plan
	return plan
}

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
