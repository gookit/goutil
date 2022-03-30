// Package httpreq an simple http requester
package httpreq

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gookit/goutil/netutil/httpctype"
)

// HttpReq an simple http requester.
type HttpReq struct {
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
}

// New instance with base URL
func New(baseURL ...string) *HttpReq {
	h := &HttpReq{
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
func (h *HttpReq) BaseURL(baseURL string) *HttpReq {
	h.baseURL = baseURL
	return h
}

// Method with custom method
func (h *HttpReq) Method(method string) *HttpReq {
	h.method = method
	return h
}

// WithHeader with custom header
func (h *HttpReq) WithHeader(key, val string) *HttpReq {
	h.headerMap[key] = val
	return h
}

// WithHeaders with custom headers
func (h *HttpReq) WithHeaders(kvMap map[string]string) *HttpReq {
	for k, v := range kvMap {
		h.headerMap[k] = v
	}
	return h
}

// ContentType with custom content-Type header.
func (h *HttpReq) ContentType(cType string) *HttpReq {
	return h.WithHeader("Content-Type", cType)
}

// BeforeSend add callback before send.
func (h *HttpReq) BeforeSend(fn func(req *http.Request)) *HttpReq {
	h.beforeSend = fn
	return h
}

// WithBody with custom body
func (h *HttpReq) WithBody(r io.Reader) *HttpReq {
	h.body = r
	return h
}

// BytesBody with custom bytes body
func (h *HttpReq) BytesBody(bs []byte) *HttpReq {
	h.body = bytes.NewReader(bs)
	return h
}

// JSONBytesBody with custom bytes body, and set JSON content type
func (h *HttpReq) JSONBytesBody(bs []byte) *HttpReq {
	h.body = bytes.NewReader(bs)
	h.ContentType(httpctype.JSON)
	return h
}

// StringBody with custom string body
func (h *HttpReq) StringBody(s string) *HttpReq {
	h.body = strings.NewReader(s)
	return h
}

// Client custom http client
func (h *HttpReq) Client(c Doer) *HttpReq {
	h.client = c
	return h
}

// MustSend request, will panic on error
func (h *HttpReq) MustSend(url string) *http.Response {
	resp, err := h.Send(url)
	if err != nil {
		panic(err)
	}

	return resp
}

// Send request and return http response
func (h *HttpReq) Send(url string) (*http.Response, error) {
	if len(h.baseURL) > 0 {
		if !strings.HasPrefix(url, "http") {
			url = h.baseURL + url
		} else if len(url) == 0 {
			url = h.baseURL
		}
	}

	// create request
	req, err := http.NewRequest(h.method, url, h.body)
	if err != nil {
		return nil, err
	}

	if len(h.headerMap) > 0 {
		for k, v := range h.headerMap {
			req.Header.Set(k, v)
		}
	}

	if h.beforeSend != nil {
		h.beforeSend(req)
	}
	return h.client.Do(req)
}
