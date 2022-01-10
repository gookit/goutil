package httpreq

import "net/http"

// HttpDoer interface for http client.
type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}
