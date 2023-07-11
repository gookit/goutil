// Package httpreq provide an simple http requester and some useful util functions.
package httpreq

import (
	"net/http"
	"sync"
	"time"

	"github.com/gookit/goutil/netutil/httpctype"
)

// AfterSendFn callback func
type AfterSendFn func(resp *http.Response, err error)

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

// ReqLogger request logger interface
type ReqLogger interface {
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
}

const defaultTimeout = 500 * time.Millisecond

var (
	// global lock
	_gl = sync.Mutex{}

	// client cache map
	cs = map[int]*Client{}
)

// NewClient create a new http client and cache it.
//
// Note: timeout unit is millisecond
func NewClient(timeout int) *Client {
	_gl.Lock()
	cli, ok := cs[timeout]

	if !ok {
		cli = NewWithTimeout(timeout)
		cs[timeout] = cli
	}

	_gl.Unlock()
	return cli
}

// MustResp check error and return response
func MustResp(r *http.Response, err error) *http.Response {
	if err != nil {
		panic(err)
	}
	return r
}

// MustRespX check error and create a new RespX instance
func MustRespX(r *http.Response, err error) *RespX {
	if err != nil {
		panic(err)
	}
	return NewResp(r)
}

// WithJSONType set request content type to JSON
func WithJSONType(opt *Option) {
	opt.ContentType = httpctype.JSON
}

// WithData set request data, will auto convert to body data or query string
func WithData(data any) OptionFn {
	return func(opt *Option) {
		opt.Data = data
	}
}
