package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

// some data.
type (
	// M short name for map
	M map[string]string
	// MD simple request data
	MD struct {
		// Headers headers
		Headers M
		// Body body. eg: strings.NewReader("name=inhere")
		Body io.Reader
		// BodyString quick add body.
		BodyString string
		// BeforeSend callback
		BeforeSend func(req *http.Request)
	}
)

// NewHttpRequest for http testing
// Usage:
//
// 	req := NewHttpRequest("GET", "/path", nil)
//
// 	// with data 1
// 	body := strings.NewReader("string ...")
// 	req := NewHttpRequest("POST", "/path", &MD{
// 		Body: body,
// 		Headers: M{"x-head": "val"}
// 	})
//
// 	// with data 2
// 	req := NewHttpRequest("POST", "/path", &MD{
// 		BodyString: "data string",
// 		Headers: M{"x-head": "val"}
// 	})
func NewHttpRequest(method, path string, data *MD) *http.Request {
	var body io.Reader
	if data != nil {
		if data.Body != nil {
			body = data.Body
		} else if data.BodyString != "" {
			body = strings.NewReader(data.BodyString)
		}
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

	return req
}

// MockRequest mock an HTTP Request
//
// Usage:
// 	handler := router.New()
// 	res := MockRequest(handler, "GET", "/path", nil)
//
// 	// with data 1
// 	body := strings.NewReader("string ...")
// 	res := MockRequest(handler, "POST", "/path", &MD{
// 		Body: body,
// 		Headers: M{"x-head": "val"}
// 	})
//
// 	// with data 2
// 	res := MockRequest(handler, "POST", "/path", &MD{
// 		BodyString: "data string",
// 		Headers: M{"x-head": "val"}
// 	})
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder {
	// w.Result() will return http.Response
	w := httptest.NewRecorder()
	r := NewHttpRequest(method, path, data)

	// s := httptest.NewServer()
	h.ServeHTTP(w, r)
	return w
}
