package http

import (
	"bytes"
	"encoding/json"
	"infantry/bindings"
	"net/http"
	"time"
)

type Agent struct {
	Host          string
	Url           string
	Body          interface{}
	Timeout       int
	SkipVerifySsl bool
	Headers       []bindings.Header
}

// CreateTest Creates the base test struct for reporting
func CreateTest(method string, uri string) bindings.Test {
	var test bindings.Test
	test.Timestamp = time.Now().UTC().String()
	test.Method = method
	test.Uri = uri
	return test
}

// TestGet Tests a get request and appends custom headers
func TestGet(agent Agent) bindings.Test {
	test := CreateTest(http.MethodGet, agent.Host+agent.Url)
	client := CreateHttpClient(agent.SkipVerifySsl)
	client.Timeout = time.Duration(agent.Timeout)
	req, err := http.NewRequest(http.MethodGet, agent.Host+agent.Url, nil)
	if agent.Headers != nil {
		SetTestHeaders(req, agent.Headers)
	}
	if err != nil {
		test.Error = err.Error()
		return test
	}

	ts := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		test.Error = err.Error()
		return test
	}

	test.ResponseTimeMs = time.Since(ts).Milliseconds()
	test.HttpStatus = resp.StatusCode
	return test
}

// TestPost Tests a post request, adds body as json, and appends custom headers
func TestPost(agent Agent) bindings.Test {
	test := CreateTest(http.MethodPost, agent.Host+agent.Url)
	client := CreateHttpClient(agent.SkipVerifySsl)
	client.Timeout = time.Duration(agent.Timeout)
	jsonData, err := json.Marshal(agent.Body)
	req, err := http.NewRequest(http.MethodPost, agent.Host+agent.Url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if agent.Headers != nil {
		SetTestHeaders(req, agent.Headers)
	}
	if err != nil {
		test.Error = err.Error()
		return test
	}

	ts := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		test.Error = err.Error()
		return test
	}
	defer resp.Body.Close()

	test.ResponseTimeMs = time.Since(ts).Milliseconds()
	test.HttpStatus = resp.StatusCode
	return test
}

func TestPatch(uri string, body interface{}, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodPatch, uri)
	return test
}

func TestDelete(uri string, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodDelete, uri)
	return test
}

func TestHead(uri string, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodHead, uri)
	return test
}

func TestOptions(uri string, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodOptions, uri)
	return test
}
