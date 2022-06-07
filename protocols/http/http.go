package http

import (
	"bytes"
	"encoding/json"
	"infantry/bindings"
	"net/http"
	"time"
)

// CreateTest Creates the base test struct for reporting
func CreateTest(method string, uri string) bindings.Test {
	var test bindings.Test
	test.Timestamp = time.Now().UTC().String()
	test.Method = method
	test.Uri = uri
	return test
}

// TestGet Tests a get request and appends custom headers
func TestGet(uri string, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodGet, uri)
	client := CreateHttpClient(skipSsl)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if headers != nil {
		SetTestHeaders(req, headers)
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
func TestPost(uri string, body interface{}, headers []bindings.Header, skipSsl bool) bindings.Test {
	test := CreateTest(http.MethodPost, uri)
	client := CreateHttpClient(skipSsl)
	jsonData, err := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if headers != nil {
		SetTestHeaders(req, headers)
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
