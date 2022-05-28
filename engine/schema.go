package engine

import (
	"flag"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"infantry/bindings"
	"io/ioutil"
)

func readFile(path string) (*bindings.Plan, error) {
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

// LoadPlanSchemaFromPath Reads the plan yaml file
func LoadPlanSchemaFromPath(pathToConfig string) bindings.Plan {
	cfgPath := flag.String("p", pathToConfig, "Path to yaml config file")
	flag.Parse()
	cfg, err := readFile(*cfgPath)
	if err != nil {
		panic(err.Error())
	}
	return *cfg
}
