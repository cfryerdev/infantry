package types

import (
	"fmt"
	"infantry/bindings"
)

// Convert This converts an artillery file in json or yaml, into a plan yaml
func Convert(scenario string) bindings.Plan {
	fmt.Printf("DEBUG: %s", scenario)
	return bindings.Plan{}
}
