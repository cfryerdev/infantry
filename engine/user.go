package engine

import (
	"github.com/google/uuid"
)

type User struct {
	Id   uuid.UUID
	Host string
}

// CreateVirtualUser Creates a virtual user for use with proposals
func CreateVirtualUser() User {
	return User{}
}
