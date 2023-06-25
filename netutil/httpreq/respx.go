package httpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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
}

// WrapResp wrap http.Response to RespX
func WrapResp(hr *http.Response, err error) (*RespX, error) {
	if err != nil {
		return nil, err
	}
	return &RespX{Response: hr}, nil
}

// NewResp instance
func NewResp(hr *http.Response) *RespX {
	return &RespX{Response: hr}
}

// IsFail check
func (r *RespX) IsFail() bool {
	return r.StatusCode != http.StatusOK
}

// IsOk check
func (r *RespX) IsOk() bool {
	return r.StatusCode == http.StatusOK
}

// IsSuccessful check
func (r *RespX) IsSuccessful() bool {
	return IsSuccessful(r.StatusCode)
}

// IsEmptyBody check response body is empty
func (r *RespX) IsEmptyBody() bool {
	return r.ContentLength <= 0
}

// ContentType get response content type
func (r *RespX) ContentType() string {
	return r.Header.Get(httpctype.Key)
}

// BodyString get body as string.
func (r *RespX) String() string {
	return ResponseToString(r.Response)
}

// BodyString get body as string.
func (r *RespX) BodyString() string {
	return r.BodyBuffer().String()
}

// BodyBuffer read body to buffer.
//
// NOTICE: must close resp body.
func (r *RespX) BodyBuffer() *bytes.Buffer {
	buf := &bytes.Buffer{}
	// prof: assign memory before read
	if r.ContentLength > bytes.MinRead {
		buf.Grow(int(r.ContentLength) + 2)
	}

	// NOTICE: must close resp body.
	defer r.SafeCloseBody()
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		panic(err)
	}

	return buf
}

// BindJSONOnOk body data on response status is 200.
// if ptr is nil, will discard body data.
//
// NOTICE: must close resp body.
func (r *RespX) BindJSONOnOk(ptr any) error {
	// NOTICE: must close resp body.
	defer r.SafeCloseBody()

	if r.IsFail() {
		_, _ = io.Copy(io.Discard, r.Body) // <-- add this line
		return errors.New("response status is not equals to 200")
	}

	if ptr == nil {
		_, _ = io.Copy(io.Discard, r.Body) // <-- add this line
		return nil
	}
	return json.NewDecoder(r.Body).Decode(ptr)
}

// BindJSON body data to a ptr, will don't check status code.
// if ptr is nil, will discard body data.
//
// NOTICE: must close resp body.
func (r *RespX) BindJSON(ptr any) error {
	// NOTICE: must close resp body.
	defer r.SafeCloseBody()

	if ptr == nil {
		_, _ = io.Copy(io.Discard, r.Body) // <-- add this line
		return nil
	}
	return json.NewDecoder(r.Body).Decode(ptr)
}

// CloseBody close resp body
func (r *RespX) CloseBody() error {
	return r.Body.Close()
}

// SafeCloseBody close resp body, ignore error
func (r *RespX) SafeCloseBody() {
	_ = r.Body.Close()
}
