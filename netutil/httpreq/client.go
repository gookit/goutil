// Package httpreq an simple http requester
package httpreq

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gookit/goutil/netutil/httpctype"
)

// ReqOption struct
type ReqOption struct {
	// HeaderMap data. eg: traceid, parentid
	HeaderMap map[string]string
	// Timeout unit: ms
	Timeout int
	// ContentType header
	ContentType string
	// EncodeJSON req body
	EncodeJSON bool
	// Logger for request
	Logger ReqLogger
}

// Req alias of ReqClient
//
// Deprecated: rename to ReqClient
type Req = ReqClient

// ReqClient an simple http request client.
type ReqClient struct {
	client Doer
	// some config for request
	method  string
	baseURL string
	// custom set headers
	headerMap map[string]string
	// request body.
	// eg: strings.NewReader("name=inhere")
	body io.Reader
	// beforeSend callback
	beforeSend func(req *http.Request)
	afterSend  func(resp *http.Response)
}

// default std client
var std = New().Client(&http.Client{Timeout: 500 * time.Millisecond})
var emptyOpt = &ReqOption{}

// ConfigStd req client
func ConfigStd(fn func(hc *http.Client)) {
	fn(std.client.(*http.Client))
}

// Std instance
func Std() *ReqClient { return std }

// Get quick send a GET request by default client
func Get(url string, opt *ReqOption) (*http.Response, error) {
	return std.Method(http.MethodGet).SendWithOpt(url, opt)
}

// Post quick send a POS request by default client
func Post(url string, data any, opt *ReqOption) (*http.Response, error) {
	return std.Method(http.MethodPost).AnyBody(data).SendWithOpt(url, opt)
}

// New instance with base URL
func New(baseURL ...string) *ReqClient {
	h := &ReqClient{
		method: http.MethodGet,
		client: http.DefaultClient,
		// init map
		headerMap: make(map[string]string),
	}

	if len(baseURL) > 0 {
		h.baseURL = baseURL[0]
	}
	return h
}

// BaseURL with base URL
func (h *ReqClient) BaseURL(baseURL string) *ReqClient {
	h.baseURL = baseURL
	return h
}

// Method with custom method
func (h *ReqClient) Method(method string) *ReqClient {
	if method != "" {
		h.method = method
	}
	return h
}

// WithHeader with custom header
func (h *ReqClient) WithHeader(key, val string) *ReqClient {
	h.headerMap[key] = val
	return h
}

// WithHeaders with custom headers
func (h *ReqClient) WithHeaders(kvMap map[string]string) *ReqClient {
	for k, v := range kvMap {
		h.headerMap[k] = v
	}
	return h
}

// ContentType with custom content-Type header.
func (h *ReqClient) ContentType(cType string) *ReqClient {
	return h.WithHeader(httpctype.Key, cType)
}

// BeforeSend add callback before send.
func (h *ReqClient) BeforeSend(fn func(req *http.Request)) *ReqClient {
	h.beforeSend = fn
	return h
}

// AfterSend add callback after send.
func (h *ReqClient) AfterSend(fn func(resp *http.Response)) *ReqClient {
	h.afterSend = fn
	return h
}

// WithBody with custom body
func (h *ReqClient) WithBody(r io.Reader) *ReqClient {
	h.body = r
	return h
}

// BytesBody with custom bytes body
func (h *ReqClient) BytesBody(bs []byte) *ReqClient {
	h.body = bytes.NewReader(bs)
	return h
}

// JSONBytesBody with custom bytes body, and set JSON content type
func (h *ReqClient) JSONBytesBody(bs []byte) *ReqClient {
	h.body = bytes.NewReader(bs)
	h.ContentType(httpctype.JSON)
	return h
}

// StringBody with custom string body
func (h *ReqClient) StringBody(s string) *ReqClient {
	h.body = strings.NewReader(s)
	return h
}

// AnyBody with custom body.
//
// Allow type:
//   - string, []byte, map[string][]string/url.Values, io.Reader(eg: bytes.Buffer, strings.Reader)
func (h *ReqClient) AnyBody(data any) *ReqClient {
	h.body = ToRequestBody(data)
	return h
}

// Client custom http client
func (h *ReqClient) Client(c Doer) *ReqClient {
	h.client = c
	return h
}

// MustSend request, will panic on error
func (h *ReqClient) MustSend(url string) *http.Response {
	resp, err := h.Send(url)
	if err != nil {
		panic(err)
	}

	return resp
}

// Send request and return http response
func (h *ReqClient) Send(url string) (*http.Response, error) {
	return h.SendWithOpt(url, nil)
}

// SendWithOpt request and return http response
func (h *ReqClient) SendWithOpt(url string, opt *ReqOption) (*http.Response, error) {
	cli := h
	if len(cli.baseURL) > 0 {
		if !strings.HasPrefix(url, "http") {
			url = cli.baseURL + url
		} else if len(url) == 0 {
			url = cli.baseURL
		}
	}

	// create request
	req, err := http.NewRequest(cli.method, url, cli.body)
	if err != nil {
		return nil, err
	}
	return h.SendRequest(req, opt)
}

// SendRequest request and return http response
func (h *ReqClient) SendRequest(req *http.Request, opt *ReqOption) (*http.Response, error) {
	if opt == nil {
		opt = emptyOpt
	}

	cli := h
	if opt.Timeout > 0 {
		cli = New().Client(&http.Client{
			Timeout: time.Duration(opt.Timeout) * time.Millisecond,
		})
	}

	if len(cli.headerMap) > 0 {
		for k, v := range cli.headerMap {
			req.Header.Set(k, v)
		}
	}

	if cli.beforeSend != nil {
		cli.beforeSend(req)
	}

	resp, err := cli.client.Do(req)
	if err == nil && cli.afterSend != nil {
		cli.afterSend(resp)
	}
	return resp, err
}

// Doer get the http client
func (h *ReqClient) Doer() Doer {
	return h.client
}
