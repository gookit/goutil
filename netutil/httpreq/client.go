package httpreq

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/strutil"
)

// Client a simple http request client.
type Client struct {
	client Doer
	// default config for request
	method  string
	baseURL string
	timeout int // unit: ms
	// custom set headers
	headerMap map[string]string

	// beforeSend callback
	beforeSend func(req *http.Request)
	afterSend  AfterSendFn
}

// New instance with base URL and use http.Client as default http client
func New(baseURL ...string) *Client {
	h := NewWithDoer(&http.Client{Timeout: defaultTimeout})
	if len(baseURL) > 0 && baseURL[0] != "" {
		h.baseURL = baseURL[0]
	}
	return h
}

// NewWithTimeout new instance use http.Client and with custom timeout(ms)
func NewWithTimeout(ms int) *Client {
	return NewWithDoer(&http.Client{Timeout: time.Duration(ms) * time.Millisecond})
}

// NewWithDoer instance with custom http client
func NewWithDoer(d Doer) *Client {
	return &Client{
		client: d,
		method: http.MethodGet,
		// init map
		headerMap: make(map[string]string),
	}
}

// Doer get the http client
func (h *Client) Doer() Doer {
	return h.client
}

// SetClient custom set http client doer
func (h *Client) SetClient(c Doer) *Client {
	h.client = c
	return h
}

// SetTimeout set default timeout for http client doer
func (h *Client) SetTimeout(ms int) *Client {
	h.timeout = ms
	if hc, ok := h.client.(*http.Client); ok {
		hc.Timeout = time.Duration(ms) * time.Millisecond
	}
	return h
}

// BaseURL set request base URL
func (h *Client) BaseURL(baseURL string) *Client {
	h.baseURL = baseURL
	return h
}

// DefaultMethod set default request method
func (h *Client) DefaultMethod(method string) *Client {
	if method != "" {
		h.method = method
	}
	return h
}

// ContentType set default content-Type header.
func (h *Client) ContentType(cType string) *Client {
	return h.DefaultHeader(httpctype.Key, cType)
}

// DefaultHeader set default header for all requests
func (h *Client) DefaultHeader(key, val string) *Client {
	h.headerMap[key] = val
	return h
}

// DefaultHeaderMap set default headers for all requests
func (h *Client) DefaultHeaderMap(kvMap map[string]string) *Client {
	for k, v := range kvMap {
		h.headerMap[k] = v
	}
	return h
}

// OnBeforeSend add callback before send.
func (h *Client) OnBeforeSend(fn func(req *http.Request)) *Client {
	h.beforeSend = fn
	return h
}

// OnAfterSend add callback after send.
func (h *Client) OnAfterSend(fn AfterSendFn) *Client {
	h.afterSend = fn
	return h
}

//
// build request options
//

// WithOption with custom request options
func (h *Client) WithOption(optFns ...OptionFn) *Option {
	return NewOption(optFns).WithClient(h)
}

func optWithClient(cli *Client) *Option {
	return &Option{cli: cli}
}

// WithData with custom request data
func (h *Client) WithData(data any) *Option {
	return optWithClient(h).WithData(data)
}

// WithBody with custom body
func (h *Client) WithBody(r io.Reader) *Option {
	return optWithClient(h).WithBody(r)
}

// BytesBody with custom bytes body
func (h *Client) BytesBody(bs []byte) *Option {
	return optWithClient(h).BytesBody(bs)
}

// StringBody with custom string body
func (h *Client) StringBody(s string) *Option {
	return optWithClient(h).StringBody(s)
}

// FormBody with custom form data body
func (h *Client) FormBody(data any) *Option {
	return optWithClient(h).FormBody(data)
}

// JSONBody with custom JSON data body
func (h *Client) JSONBody(data any) *Option {
	return optWithClient(h).WithJSON(data)
}

// JSONBytesBody with custom bytes body, and set JSON content type
func (h *Client) JSONBytesBody(bs []byte) *Option {
	return optWithClient(h).JSONBytesBody(bs)
}

// AnyBody with custom body.
//
// Allow type:
//   - string, []byte, map[string][]string/url.Values, io.Reader(eg: bytes.Buffer, strings.Reader)
func (h *Client) AnyBody(data any) *Option {
	return optWithClient(h).AnyBody(data)
}

//
// ------------ send request with options ------------
//

// Get send GET request with options, return http response
func (h *Client) Get(url string, optFns ...OptionFn) (*http.Response, error) {
	return h.Send(http.MethodGet, url, optFns...)
}

// Post send POST request with options, return http response
func (h *Client) Post(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	opt := NewOption(optFns).WithMethod(http.MethodPost).AnyBody(data)
	return h.SendWithOpt(url, opt)
}

// PostJSON send JSON POST request with options, return http response
func (h *Client) PostJSON(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	opt := NewOption(optFns).WithMethod(http.MethodPost).WithJSON(data)
	return h.SendWithOpt(url, opt)
}

// Put send PUT request with options, return http response
func (h *Client) Put(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	opt := NewOption(optFns).WithMethod(http.MethodPut).AnyBody(data)
	return h.SendWithOpt(url, opt)
}

// Delete send DELETE request with options, return http response
func (h *Client) Delete(url string, optFns ...OptionFn) (*http.Response, error) {
	return h.Send(http.MethodDelete, url, optFns...)
}

// Send request with option func, return http response
func (h *Client) Send(method, url string, optFns ...OptionFn) (*http.Response, error) {
	return h.SendWithOpt(url, NewOption(optFns).WithMethod(method))
}

// MustSend request, will panic on error
func (h *Client) MustSend(method, url string, optFns ...OptionFn) *http.Response {
	resp, err := h.SendWithOpt(url, NewOption(optFns).WithMethod(method))
	if err != nil {
		panic(err)
	}
	return resp
}

// SendWithOpt request and return http response
func (h *Client) SendWithOpt(url string, opt *Option) (*http.Response, error) {
	cli := h
	if len(cli.baseURL) > 0 {
		if !strings.HasPrefix(url, "http") {
			url = cli.baseURL + url
		} else if len(url) == 0 {
			url = cli.baseURL
		}
	}

	opt = OptOrNew(opt)
	ctx := opt.Context
	if ctx == nil {
		ctx = context.Background()
	}

	// create request
	method := strings.ToUpper(strutil.OrElse(opt.Method, cli.method))

	if opt.Data != nil {
		if IsNoBodyMethod(method) {
			url = AppendQueryToURLString(url, MakeQuery(opt.Data))
			opt.Body = nil
		} else if opt.Body == nil {
			cType := strutil.OrElse(h.headerMap[httpctype.Key], opt.ContentType)
			opt.Body = MakeBody(opt.Data, cType)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, opt.Body)
	if err != nil {
		return nil, err
	}

	return h.sendRequest(req, opt)
}

// SendRequest send request and return http response
func (h *Client) SendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	return h.sendRequest(req, opt)
}

// send request and return http response
func (h *Client) sendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	// apply default headers
	if len(h.headerMap) > 0 {
		for k, v := range h.headerMap {
			req.Header.Set(k, v)
		}
	}

	// apply options
	if opt.ContentType != "" {
		req.Header.Set(httpctype.Key, opt.ContentType)
	}

	// - apply header map
	if len(opt.HeaderMap) > 0 {
		for k, v := range opt.HeaderMap {
			req.Header.Set(k, v)
		}
	}

	cli := h // if timeout changed, create new client
	if opt.Timeout > 0 && opt.Timeout != cli.timeout {
		cli = NewClient(opt.Timeout)
	}

	if h.beforeSend != nil {
		h.beforeSend(req)
	}

	resp, err := cli.client.Do(req)
	if h.afterSend != nil {
		h.afterSend(resp, err)
	}
	return resp, err
}
