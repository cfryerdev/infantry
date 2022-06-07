package engine

import (
	"github.com/google/uuid"
)

type User struct {
	Id   uuid.UUID
	Task UserTask
}

type UserCluster struct {
	Cluster []User
}

type UserTask func(int, int) interface{}

// CreateVirtualUser Creates a virtual user for use with a single proposal
func CreateVirtualUser() User {
	return User{}
}
