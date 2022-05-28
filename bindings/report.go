package bindings

type Report struct {
	Summary    Summary
	Iterations []Iteration
}

type Summary struct {
	StartTimestamp string
	EndTimestamp   string
}

type Iteration struct {
	Timestamp         string
	Requests          int
	RequestsPerSecond int
	Codes             map[string]string
}
