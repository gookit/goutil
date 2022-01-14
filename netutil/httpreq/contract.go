package httpreq

import "net/http"

// Doer interface for http client.
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// DoerFunc implements the Doer
type DoerFunc func(req *http.Request) (*http.Response, error)

// Do send request and return response.
func (do DoerFunc) Do(req *http.Request) (*http.Response, error) {
	return do(req)
}
