package http

import (
	"infantry/bindings"
	"net/http"
	"time"
)

// TestGet Tests a get request and appends custom headers
func TestGet(uri string, headers []bindings.Header, skipSsl bool) bindings.Test {
	var test bindings.Test
	test.Timestamp = time.Now().UTC().String()
	test.Method = http.MethodGet
	test.Uri = uri

	var client = CreateHttpClient(skipSsl)
	var req, err = http.NewRequest(http.MethodGet, uri, nil)
	if headers != nil {
		SetTestHeaders(req, headers)
	}
	if err != nil {
		test.Error = err.Error()
		return test
	}

	var ts = time.Now()
	resp, err := client.Do(req)
	if err != nil {
		test.Error = err.Error()
		return test
	}

	test.ResponseTimeMs = time.Since(ts).Milliseconds()
	test.HttpStatus = resp.StatusCode
	return test
}
