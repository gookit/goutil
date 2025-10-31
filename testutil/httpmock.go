package testutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/basefn"
)

// some data.
type (
	// M short name for a string-map
	M map[string]string
	// MD simple request data
	MD struct {
		// Headers headers
		Headers M
		// Body reader. eg: strings.NewReader("name=inhere")
		Body io.Reader
		// BodyString quick adds string body.
		BodyString string
		// BeforeSend callback
		BeforeSend func(req *http.Request)
	}
)

// NewHTTPRequest quick create request for http testing
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
	URL    string `json:"url"`
	Method string `json:"method"`
	// Query data
	Query map[string]any `json:"query,omitempty"`
	// Headers data.
	//
	// If value is one elem, will return string, otherwise will return []string
	//
	// Example:
	// 	map[string]any{
	//		"Connection": "close",
	//		"Vary": []string{"Accept-Encoding", "Accept-Encoding"},
	//	}
	Headers map[string]any `json:"headers,omitempty"`
	// Form data.
	//
	// If value is one elem, will return string, otherwise will return []string
	Form map[string]any `json:"form,omitempty"`
	// Body data string from request body
	Body string `json:"body,omitempty"`
	// JSON data on Content-Type: application/json
	//  - 通常是 map[string]any 类型数据
	JSON any `json:"json,omitempty"`
	// Files data.
	Files map[string]any `json:"files,omitempty"`
}

// ContentType get content type
func (r *EchoReply) ContentType() string {
	return r.Headers["Content-Type"].(string)
}

// JSONMap assert JSON data to map[string]any
func (r *EchoReply) JSONMap() map[string]any {
	if r.JSON == nil {
		return nil
	}
	if m, ok := r.JSON.(map[string]any); ok {
		return m
	}
	return nil
}

// HeaderString get header value as string
func (r *EchoReply) HeaderString(name string) string {
	if r.Headers == nil {
		return ""
	}

	val := r.Headers[name]
	if s, ok := val.(string); ok {
		return s
	}
	return fmt.Sprint(val)
}

// EchoServer for testing http request.
type EchoServer struct {
	*httptest.Server
}

// HostAddr get host address. eg: 127.0.0.1:8999
func (s *EchoServer) HostAddr() string {
	return s.Listener.Addr().String()
}

// HTTPHost get http host address. eg: http://127.0.0.1:8999
func (s *EchoServer) HTTPHost() string {
	return "http://" + s.HostAddr()
}

// PrintHttpHost print host address to console
func (s *EchoServer) PrintHttpHost() string {
	baseUrl := s.HTTPHost()
	fmt.Println("Test server listen on:", baseUrl)
	return baseUrl
}

// HandleRequest handle request
func (s *EchoServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	pathName := strings.Trim(r.URL.Path, "/")
	if !strings.Contains(pathName, "/") {
		switch pathName {
		case "404": // eg. GET /404
			w.WriteHeader(http.StatusNotFound)
			return
		case "500": // eg. GET /500
			w.WriteHeader(http.StatusInternalServerError)
			return
		default:
			// custom reply status code. eg: /status-{code}
			if strings.HasPrefix(pathName, "status-") {
				stCode := pathName[7:]
				if codeVal := strutil.SafeInt(stCode); codeVal > 0 {
					w.WriteHeader(codeVal)
					return
				}
			} else {
				// 405 eg: "GET /post"
				pathMethod := strings.ToUpper(pathName)
				if httpreq.IsValidMethod(pathMethod) && pathMethod != r.Method {
					w.WriteHeader(http.StatusMethodNotAllowed)
					return
				}
			}
		}
	}

	// default: 200 ok
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Server", "goutil/echo-server")
	w.WriteHeader(http.StatusOK)
	// w.Header().Set("Connection", "close")

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	basefn.MustOK(enc.Encode(BuildEchoReply(r)))
}

// MockHttpServer create an echo server for testing. alias of NewEchoServer
func MockHttpServer() *EchoServer { return NewEchoServer() }

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
//		testSrvAddr = s.PrintHttpHost()
//
//		m.Run()
//	}
//
//	// in a test case ...
//	res := http.Get(testSrvAddr + "/get/some-one")
//	rpl := testutil.ParseRespToReply(res)
//	// assert ...
func NewEchoServer() *EchoServer {
	es := &EchoServer{}
	es.Server = httptest.NewServer(http.HandlerFunc(es.handleRequest))

	return es
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
		form = stringsMapToAnyMap(r.PostForm)

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
		if len(r.PostForm) > 0 {
			bodyStr = r.PostForm.Encode()
		} else if r.Body != nil {
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
		URL:     r.URL.String(),
		Origin:  r.RemoteAddr,
		Method:  method,
		Body:    bodyStr,
		JSON:    jsonBody,
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
		return &EchoReply{}
	}

	if w.Request != nil && w.Request.Method == "HEAD" {
		req := w.Request
		rpl := &EchoReply{
			URL:     req.URL.String(),
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
