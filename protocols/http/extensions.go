package http

import (
	"crypto/tls"
	"infantry/bindings"
	"net/http"
)

// SetTestHeaders Sets the headers on the request based on header array
func SetTestHeaders(req *http.Request, headers []bindings.Header) {
	for i := 0; i < len(headers); i++ {
		var head = headers[i]
		req.Header.Set(head.Key, head.Value)
	}
}

// CreateHttpClient Creates a http client and enables or disables the ssl verification check
func CreateHttpClient(skipSsl bool) http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipSsl},
	}
	return http.Client{Transport: tr}
}
