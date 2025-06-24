package httpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gookit/goutil/netutil/httpctype"
)

// Resp alias of RespX
type Resp = RespX

// RespX wrap http.Response and add some useful methods.
type RespX struct {
	*http.Response
	// CostTime for a request-response
	CostTime int64
	// body data buffer
	bodyBuf *bytes.Buffer
}

// WrapResp wrap http.Response to RespX
func WrapResp(hr *http.Response, err error) (*RespX, error) {
	if err != nil {
		return nil, err
	}
	return &RespX{Response: hr}, nil
}

// NewResp instance
func NewResp(hr *http.Response) *RespX { return &RespX{Response: hr} }

// IsFail check status code is not equals to 200
func (r *RespX) IsFail() bool { return r.StatusCode != http.StatusOK }

// IsOk check status code is equals to 200
func (r *RespX) IsOk() bool { return r.StatusCode == http.StatusOK }

// IsSuccessful check status code is in 200-300
func (r *RespX) IsSuccessful() bool { return IsSuccessful(r.StatusCode) }

// IsEmptyBody check response body is empty
func (r *RespX) IsEmptyBody() bool { return r.ContentLength <= 0 }

// ContentType get response content type
func (r *RespX) ContentType() string { return r.Header.Get(httpctype.Key) }

// BodyString get body as string.
func (r *RespX) String() string { return ResponseToString(r.Response) }

//
// ------------------------ read body ------------------------
//

// ReadBody read body to buffer. allow reading body multiple times.
//
// NOTE: will close the response.Body
func (r *RespX) ReadBody() error {
	if r.bodyBuf != nil {
		return nil
	}

	// prof: assign memory before read
	if r.ContentLength > bytes.MinRead {
		r.bodyBuf = bytes.NewBuffer(make([]byte, 0, r.ContentLength))
	} else {
		r.bodyBuf = bytes.NewBuffer(make([]byte, 0, bytes.MinRead))
	}

	// NOTICE: must close resp body.
	defer r.SafeCloseBody()
	_, err := r.bodyBuf.ReadFrom(r.Body)
	return err
}

// BodyBuffer read body to buffer. NOTE: will close the response.Body
func (r *RespX) BodyBuffer() *bytes.Buffer {
	if err := r.ReadBody(); err != nil {
		panic(err)
	}
	return r.bodyBuf
}

// BodyString get body as string.
func (r *RespX) BodyString() string { return r.BodyBuffer().String() }

// BindJSONOnOk body data on response status is in 200-300.
// If ptr is nil, will do nothing.
func (r *RespX) BindJSONOnOk(ptr any) error { return r.bindJSON(ptr, true) }

// BindJSON body data to a ptr, will don't check status code.
// If ptr is nil, will do nothing.
func (r *RespX) BindJSON(ptr any) error { return r.bindJSON(ptr, false) }

func (r *RespX) bindJSON(ptr any, checkStatus bool) error {
	if checkStatus && !r.IsSuccessful() {
		return errors.New("response status is not equals to 200")
	}

	if err := r.ReadBody(); err != nil {
		return err
	}

	if ptr == nil {
		return nil
	}
	return json.NewDecoder(r.bodyBuf).Decode(ptr)
}

// CloseBuffer close body buffer
func (r *RespX) CloseBuffer() {
	if r.bodyBuf != nil {
		r.bodyBuf.Reset()
		r.bodyBuf = nil
	}
}

// CloseBody close resp body
func (r *RespX) CloseBody() error { return r.Body.Close() }

// SafeCloseBody close resp body, ignore error
func (r *RespX) SafeCloseBody() {
	if r.Body != nil {
		_ = r.Body.Close()
	}
}
