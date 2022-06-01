package bindings

import "github.com/google/uuid"

type Report struct {
	Id      uuid.UUID
	Summary Summary
	Results []Iteration
}

type Summary struct {
	StartTimestamp    string
	EndTimestamp      string
	Requests          int
	AvgResponseTimeMs int
}

type Iteration struct {
	Timestamp         string
	Requests          int
	AvgResponseTimeMs int
	Tests             []Test
}

type Test struct {
	Method         string
	Uri            string
	ResponseTimeMs int
	HttpStatus     int
}
