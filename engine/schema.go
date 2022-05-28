package engine

import (
	"flag"
	"infantry/bindings"
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
func ValidateSchema() {}

// LoadEnvironmentOverrides Replaces any instances where env vars are used
func LoadEnvironmentOverrides() {}
