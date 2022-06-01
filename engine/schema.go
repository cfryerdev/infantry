package engine

import (
	"flag"
	"github.com/go-playground/validator/v10"
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
func ValidateSchema(plan bindings.Plan) {
	var planValidator = validator.New()
	err := planValidator.Struct(plan)
	if err != nil {

	}
}

// LoadEnvironmentOverrides Replaces any instances where env vars are used
func LoadEnvironmentOverrides(plan bindings.Plan) bindings.Plan {
	// v := reflect.ValueOf(plan)
	// loop through each property, and find any wildcards -> {{environment.CUSTOM_KEY_HERE}}
	// Get the env variable, then replace value and return plan
	return plan
}
