package httpreq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/netutil/httpctype"
)

// Options alias of Option
type Options = Option

// Option struct
type Option struct {
	cli  *Client
	sent bool
	// Timeout for request. unit: ms
	Timeout int
	// Method for request
	Method string
	// HeaderMap data. eg: traceid
	HeaderMap map[string]string
	// ContentType header
	ContentType string

	// Logger for request
	Logger ReqLogger
	// Context for request
	Context context.Context

	// Data for request. can be used on any request method.
	//
	// type allow:
	// 	string, []byte, io.Reader, map[string]string, ...
	Data any
	// Body data for request. used on POST, PUT, PATCH method.
	//
	// eg: strings.NewReader("name=inhere")
	Body io.Reader
}

// OptionFn option func type
type OptionFn func(opt *Option)

// OptOrNew create a new Option if opt is nil
func OptOrNew(opt *Option) *Option {
	if opt == nil {
		opt = &Option{}
	}
	return opt
}

// NewOpt create a new Option and with option func
func NewOpt(fns ...OptionFn) *Option {
	return NewOption(fns)
}

// NewOption create a new Option and set option func
func NewOption(fns []OptionFn) *Option {
	opt := &Option{}
	return opt.WithOptionFn(fns...)
}

// WithOptionFn set option func
func (o *Option) WithOptionFn(fns ...OptionFn) *Option {
	return o.WithOptionFns(fns)
}

// WithOptionFns set option func
func (o *Option) WithOptionFns(fns []OptionFn) *Option {
	for _, fn := range fns {
		if fn != nil {
			fn(o)
		}
	}
	return o
}

// WithClient set client
func (o *Option) WithClient(cli *Client) *Option {
	o.cli = cli
	return o
}

// Copy option for new request, use for repeat send request
func (o *Option) Copy() *Option {
	o.Body = nil
	on := *o
	on.sent = false
	return &on
}

// WithMethod set method
func (o *Option) WithMethod(method string) *Option {
	if method != "" {
		o.Method = method
	}
	return o
}

// WithContentType set content type
func (o *Option) WithContentType(ct string) *Option {
	o.ContentType = ct
	return o
}

// WithHeaderMap set header map
func (o *Option) WithHeaderMap(m map[string]string) *Option {
	if o.HeaderMap == nil {
		o.HeaderMap = make(map[string]string)
	}
	for k, v := range m {
		o.HeaderMap[k] = v
	}
	return o
}

// WithHeader set header
func (o *Option) WithHeader(key, val string) *Option {
	if o.HeaderMap == nil {
		o.HeaderMap = make(map[string]string)
	}
	o.HeaderMap[key] = val
	return o
}

// WithData with custom data
func (o *Option) WithData(data any) *Option {
	o.Data = data
	return o
}

// AnyBody with custom body.
//
// Allow type:
//   - string, []byte, map[string][]string/url.Values, io.Reader(eg: bytes.Buffer, strings.Reader)
func (o *Option) AnyBody(data any) *Option {
	o.Body = ToRequestBody(data, o.ContentType)
	return o
}

// WithBody with custom body
func (o *Option) WithBody(r io.Reader) *Option {
	o.Body = r
	return o
}

// BytesBody with custom bytes body
func (o *Option) BytesBody(bs []byte) *Option {
	o.Body = bytes.NewReader(bs)
	return o
}

// FormBody with custom form body data
func (o *Option) FormBody(data any) *Option {
	o.ContentType = httpctype.Form
	o.Body = ToRequestBody(data, o.ContentType)
	return o
}

// WithJSON with custom JSON body
func (o *Option) WithJSON(data any) *Option {
	o.ContentType = httpctype.JSON
	o.Body = ToRequestBody(data, o.ContentType)
	return o
}

// JSONBytesBody with custom bytes body, and set JSON content type
func (o *Option) JSONBytesBody(bs []byte) *Option {
	o.ContentType = httpctype.JSON
	return o.WithBody(bytes.NewReader(bs))
}

// StringBody with custom string body
func (o *Option) StringBody(s string) *Option {
	o.Body = strings.NewReader(s)
	return o
}

//
// send request with options
//

// Get send GET request and return http response
func (o *Option) Get(url string, fns ...OptionFn) (*http.Response, error) {
	return o.Send(http.MethodGet, url, fns...)
}

// Post send POST request and return http response
func (o *Option) Post(url string, data any, fns ...OptionFn) (*http.Response, error) {
	return o.AnyBody(data).Send(http.MethodPost, url, fns...)
}

// Put send PUT request and return http response
func (o *Option) Put(url string, data any, fns ...OptionFn) (*http.Response, error) {
	return o.AnyBody(data).Send(http.MethodPut, url, fns...)
}

// Delete send DELETE request and return http response
func (o *Option) Delete(url string, fns ...OptionFn) (*http.Response, error) {
	return o.Send(http.MethodDelete, url, fns...)
}

// Send request and return http response
func (o *Option) Send(method, url string, fns ...OptionFn) (*http.Response, error) {
	cli := basefn.OrValue(o.cli != nil, o.cli, std)
	o.sent = true
	o.WithOptionFns(fns).WithMethod(method)

	return cli.SendWithOpt(url, o)
}

// MustSend request, will panic on error
func (o *Option) MustSend(method, url string, fns ...OptionFn) *http.Response {
	cli := basefn.OrValue(o.cli != nil, o.cli, std)
	o.sent = true
	o.WithOptionFns(fns).WithMethod(method)

	resp, err := cli.SendWithOpt(url, o)
	if err != nil {
		panic(err)
	}
	return resp
}
