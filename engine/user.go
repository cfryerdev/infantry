package engine

import (
	"fmt"
	"infantry/bindings"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
}

// ExecuteProposal Creates a virtual user and executes the proposal
func (user User) ExecuteProposal(protocol bindings.Protocol, proposal bindings.Proposal) (bindings.Test, error) {
	fmt.Printf("protocol: %+v\n", protocol)
	/*if protocol.Http != bindings.ProtocolHttp{} {

	}*/
	return bindings.Test{}, nil
}
