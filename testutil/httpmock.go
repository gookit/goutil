package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// some data.
type (
	// M short name for map
	M map[string]string
	// MD simple request data
	MD struct {
		// Body body
		Body io.Reader
		// Headers headers
		Headers M
		// BeforeSend callback
		BeforeSend func(req *http.Request)
	}
)

// MockRequest mock an HTTP Request
// Usage:
// 	handler := router.New()
// 	res := MockRequest(handler, "GET", "/path", nil)
// 	// with data
// 	body := strings.NewReader("string ...")
// 	res := MockRequest(handler, "POST", "/path", &MD{Body: "data", Headers: M{"x-head": "val"}})
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder {
	var body io.Reader
	if data != nil && data.Body != nil {
		body = data.Body
	}

	// create fake request
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}

	req.RequestURI = req.URL.String()
	if data != nil {
		if len(data.Headers) > 0 {
			// req.Header.Set("Content-Type", "text/plain")
			for k, v := range data.Headers {
				req.Header.Set(k, v)
			}
		}

		if data.BeforeSend != nil {
			data.BeforeSend(req)
		}
	}

	// w.Result() will return http.Response
	w := httptest.NewRecorder()
	// s := httptest.NewServer()
	h.ServeHTTP(w, req)
	return w
}
