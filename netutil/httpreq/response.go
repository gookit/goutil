package httpreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gookit/goutil/netutil/httpctype"
)

// Resp struct
type Resp struct {
	*http.Response
	// CostTime for a request-response
	CostTime int64
}

// NewResp instance
func NewResp(hr *http.Response) *Resp {
	return &Resp{Response: hr}
}

// IsFail check
func (r *Resp) IsFail() bool {
	return r.StatusCode != http.StatusOK
}

// IsOk check
func (r *Resp) IsOk() bool {
	return r.StatusCode == http.StatusOK
}

// IsSuccessful check
func (r *Resp) IsSuccessful() bool {
	return IsSuccessful(r.StatusCode)
}

// IsEmptyBody check response body is empty
func (r *Resp) IsEmptyBody() bool {
	return r.ContentLength <= 0
}

// ContentType get response content type
func (r *Resp) ContentType() string {
	return r.Header.Get(httpctype.Key)
}

// BodyString get body as string.
func (r *Resp) String() string {
	return ResponseToString(r.Response)
}

// BodyString get body as string.
func (r *Resp) BodyString() string {
	return r.BodyBuffer().String()
}

// BodyBuffer read body to buffer.
//
// NOTICE: must close resp body.
func (r *Resp) BodyBuffer() *bytes.Buffer {
	buf := &bytes.Buffer{}
	// prof: assign memory before read
	if r.ContentLength > bytes.MinRead {
		buf.Grow(int(r.ContentLength) + 2)
	}

	// NOTICE: must close resp body.
	defer r.QuiteCloseBody()
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		panic(err)
	}

	return buf
}

// BindJSONOnOk body data on status is 200
//
// NOTICE: must close resp body.
func (r *Resp) BindJSONOnOk(ptr any) error {
	// NOTICE: must close resp body.
	defer r.QuiteCloseBody()

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

// BindJSON body data to a ptr
//
// NOTICE: must close resp body.
func (r *Resp) BindJSON(ptr any) error {
	// NOTICE: must close resp body.
	defer r.QuiteCloseBody()

	if ptr == nil {
		_, _ = io.Copy(io.Discard, r.Body) // <-- add this line
		return nil
	}

	return json.NewDecoder(r.Body).Decode(ptr)
}

// CloseBody close resp body
func (r *Resp) CloseBody() error {
	return r.Body.Close()
}

// QuiteCloseBody close resp body, ignore error
func (r *Resp) QuiteCloseBody() {
	_ = r.Body.Close()
}
