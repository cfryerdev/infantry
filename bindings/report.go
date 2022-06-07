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
	NetworkErrors     int
	Codes             []ResponseCodes
	AvgResponseTimeMs int64
}

type Iteration struct {
	Timestamp         string
	Requests          int
	AvgResponseTimeMs int64
	NetworkErrors     int
	Codes             []ResponseCodes
	Tests             []Test
}

type ResponseCodes struct {
	Informational100s int
	Successful200s    int
	Redirection300s   int
	ClientError400s   int
	ServerError500s   int
}

type Test struct {
	Timestamp      string
	Method         string
	Uri            string
	ResponseTimeMs int64
	HttpStatus     int
	Error          string
}

type MetaData struct {
	Key   string
	Value []MetaData
}
