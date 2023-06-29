package testutil

import (
	"encoding/json"
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

// NewHttpRequest for http testing, alias of NewHTTPRequest()
//
// Deprecated: use NewHTTPRequest() instead.
func NewHttpRequest(method, path string, data *MD) *http.Request {
	return NewHTTPRequest(method, path, data)
}

// NewHTTPRequest for http testing
// Usage:
//
//	req := NewHttpRequest("GET", "/path", nil)
//
//	// with data 1
//	body := strings.NewReader("string ...")
//	req := NewHttpRequest("POST", "/path", &MD{
//		Body: body,
//		Headers: M{"x-head": "val"}
//	})
//
//	// with data 2
//	req := NewHttpRequest("POST", "/path", &MD{
//		BodyString: "data string",
//		Headers: M{"x-head": "val"}
//	})
func NewHTTPRequest(method, path string, data *MD) *http.Request {
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
//
//	handler := router.New()
//	res := MockRequest(handler, "GET", "/path", nil)
//
//	// with data 1
//	body := strings.NewReader("string ...")
//	res := MockRequest(handler, "POST", "/path", &MD{
//		Body: body,
//		HeaderM: M{"x-head": "val"}
//	})
//
//	// with data 2
//	res := MockRequest(handler, "POST", "/path", &MD{
//		BodyString: "data string",
//		HeaderM: M{"x-head": "val"}
//	})
func MockRequest(h http.Handler, method, path string, data *MD) *httptest.ResponseRecorder {
	// w.Result() will return http.Response
	w := httptest.NewRecorder()
	r := NewHTTPRequest(method, path, data)

	// s := httptest.NewServer()
	h.ServeHTTP(w, r)
	return w
}

// EchoReply http response data reply model
type EchoReply struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
	Method string `json:"method"`
	// Query data
	Query   map[string]any `json:"query,omitempty"`
	Headers map[string]any `json:"headers,omitempty"`
	Form    map[string]any `json:"form,omitempty"`
	// Body data string
	Body  string         `json:"body,omitempty"`
	Json  any            `json:"json,omitempty"`
	Files map[string]any `json:"files,omitempty"`
}

// ContentType get content type
func (r *EchoReply) ContentType() string {
	return r.Headers["Content-Type"].(string)
}

// NewEchoServer create an echo server for testing.
//
// Usage on testing:
//
//	var testSrvAddr string
//
//	func TestMain(m *testing.M) {
//		// create server
//		s := testutil.NewEchoServer()
//		defer s.Close()
//		testSrvAddr = "http://" + s.Listener.Addr().String()
//		fmt.Println("Test server listen on:", testSrvAddr)
//
//		m.Run()
//	}
//
//	// in a test case ...
//	res := http.Get(testSrvAddr)
//	rpl := testutil.ParseRespToReply(res)
//	// assert ...
func NewEchoServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Server", "goutil/echo-server")
		w.WriteHeader(http.StatusOK)
		// w.Header().Set("Connection", "close")

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		err := enc.Encode(BuildEchoReply(r))
		if err != nil {
			_, _ = w.Write([]byte(`{"error": "encode error"}`))
		}
	}))
}

// BuildEchoReply build reply body data
func BuildEchoReply(r *http.Request) *EchoReply {
	// get headers
	headers := stringsMapToAnyMap(r.Header)

	// get query args
	args := stringsMapToAnyMap(r.URL.Query())
	cType := r.Header.Get("Content-Type")
	method := strings.ToUpper(r.Method)

	var jsonBody any
	var bodyStr string
	var form, files map[string]any

	if method == "POST" || method == "PUT" || method == "PATCH" {
		// get form data
		_ = r.ParseForm()
		form = stringsMapToAnyMap(r.Form)

		// get form files
		if r.MultipartForm != nil {
			files = make(map[string]any)
			for k, fhs := range r.MultipartForm.File {
				if len(fhs) == 1 {
					files[k] = fhs[0]
				} else {
					files[k] = fhs
				}
			}
		}

		// get body data
		if r.Body != nil {
			// defer r.Body.Close()
			body, _ := io.ReadAll(r.Body)
			bodyStr = string(body)

			// try to parse json
			if len(body) > 0 && strings.HasPrefix(cType, "application/json") {
				_ = json.Unmarshal(body, &jsonBody)
			}
		}
	}

	return &EchoReply{
		Url:     r.URL.String(),
		Origin:  r.RemoteAddr,
		Method:  method,
		Body:    bodyStr,
		Json:    jsonBody,
		Query:   args,
		Form:    form,
		Files:   files,
		Headers: headers,
	}
}

/*
// HTTP tool for testing
var HTTP = &HTTPTool{}

// HTTPTool http tool for testing
type HTTPTool struct {
}

func (ht *HTTPTool) ParseRespToReply(r *http.Response) *EchoReply {
	return ParseRespToReply(r)
}

func (ht *HTTPTool) ParseBodyToReply(bd io.ReadCloser) *EchoReply {
	return ParseBodyToReply(bd)
}
*/

// ParseRespToReply parse http response to reply
func ParseRespToReply(w *http.Response) *EchoReply {
	if w.Body == nil {
		return nil
	}

	if w.Request != nil && w.Request.Method == "HEAD" {
		req := w.Request
		rpl := &EchoReply{
			Url:     req.URL.String(),
			Method:  req.Method,
			Headers: stringsMapToAnyMap(req.Header),
		}
		return rpl
	}

	return ParseBodyToReply(w.Body)
}

// ParseBodyToReply parse http body to reply
func ParseBodyToReply(bd io.ReadCloser) *EchoReply {
	rpl := &EchoReply{}
	if bd == nil {
		return rpl
	}

	err := json.NewDecoder(bd).Decode(rpl)
	if err != nil {
		panic(err)
	}
	return rpl
}

func stringsMapToAnyMap(ssMp map[string][]string) map[string]any {
	if len(ssMp) == 0 {
		return nil
	}

	anyMp := make(map[string]any, len(ssMp))
	for k, v := range ssMp {
		if len(v) == 1 {
			anyMp[k] = v[0]
			continue
		}
		anyMp[k] = v
	}
	return anyMp
}
