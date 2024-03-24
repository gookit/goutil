package httpreq

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/strutil"
)

// BasicAuthConf struct
type BasicAuthConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// IsValid value
func (ba *BasicAuthConf) IsValid() bool {
	return ba.Password != "" && ba.Username != ""
}

// Value build to auth header "Authorization".
func (ba *BasicAuthConf) Value() string {
	return BuildBasicAuth(ba.Username, ba.Password)
}

// String build to auth header "Authorization".
func (ba *BasicAuthConf) String() string {
	return ba.Username + ":" + ba.Password
}

// IsOK check response status code is 200
func IsOK(statusCode int) bool {
	return statusCode == http.StatusOK
}

// IsSuccessful check response status code is in 200 - 300
func IsSuccessful(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < 300
}

// IsRedirect check response status code is in [301, 302, 303, 307]
func IsRedirect(statusCode int) bool {
	return statusCode == http.StatusMovedPermanently ||
		statusCode == http.StatusFound ||
		statusCode == http.StatusSeeOther ||
		statusCode == http.StatusTemporaryRedirect
}

// IsForbidden is this response forbidden(403)
func IsForbidden(statusCode int) bool {
	return statusCode == http.StatusForbidden
}

// IsNotFound is this response not found(404)
func IsNotFound(statusCode int) bool {
	return statusCode == http.StatusNotFound
}

// IsClientError check response is client error (400 - 500)
func IsClientError(statusCode int) bool {
	return statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError
}

// IsServerError check response is server error (500 - 600)
func IsServerError(statusCode int) bool {
	return statusCode >= http.StatusInternalServerError && statusCode <= 600
}

// IsNoBodyMethod check
func IsNoBodyMethod(method string) bool {
	return method != "POST" && method != "PUT" && method != "PATCH"
}

// BuildBasicAuth returns the base64 encoded username:password for basic auth.
// Then set to header "Authorization".
//
// copied from net/http.
func BuildBasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

// AddHeaders adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func AddHeaders(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

// SetHeaders sets the key, value pairs from the given http.Header to the
// request. Values for existing keys are overwritten.
func SetHeaders(req *http.Request, headers ...http.Header) {
	for _, header := range headers {
		for key, values := range header {
			req.Header[key] = values
		}
	}
}

// AddHeaderMap to reqeust instance.
func AddHeaderMap(req *http.Request, headerMap map[string]string) {
	for k, v := range headerMap {
		req.Header.Add(k, v)
	}
}

// SetHeaderMap to reqeust instance.
func SetHeaderMap(req *http.Request, headerMap map[string]string) {
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
}

// HeaderToStringMap convert
func HeaderToStringMap(rh http.Header) map[string]string {
	if len(rh) == 0 {
		return nil
	}

	mp := make(map[string]string, len(rh))
	for name, values := range rh {
		mp[name] = strings.Join(values, "; ")
	}
	return mp
}

// MakeQuery make query string, convert data to url.Values
func MakeQuery(data any) url.Values {
	return ToQueryValues(data)
}

// ToQueryValues convert string-map or any-map to url.Values
//
// data support:
//   - url.Values
//   - []byte
//   - string
//   - map[string][]string
//   - map[string]string
//   - map[string]any
func ToQueryValues(data any) url.Values {
	uv := make(url.Values)

	switch typData := data.(type) {
	// use url.Values directly if we have it
	case url.Values:
		return typData
	case map[string][]string:
		return typData
	case []byte:
		m, err := url.ParseQuery(string(typData))
		if err != nil {
			return uv
		}
		return m
	case string:
		m, err := url.ParseQuery(typData)
		if err != nil {
			return uv
		}
		return m
	case map[string]string:
		for k, v := range typData {
			uv.Add(k, v)
		}
	case map[string]any:
		for k, v := range typData {
			uv.Add(k, strutil.QuietString(v))
		}
	}
	return uv
}

// MergeURLValues merge url.Values by overwrite.
//
// values support: url.Values, map[string]string, map[string][]string
func MergeURLValues(uv url.Values, values ...any) url.Values {
	if uv == nil {
		uv = make(url.Values)
	}

	for _, v := range values {
		switch tv := v.(type) {
		case url.Values:
			for k, vs := range tv {
				uv[k] = vs
			}
		case map[string]any:
			for k, v := range tv {
				uv[k] = arrutil.AnyToStrings(v)
			}
		case map[string]string:
			for k, v := range tv {
				uv[k] = []string{v}
			}
		case map[string][]string:
			for k, vs := range tv {
				uv[k] = vs
			}
		}
	}

	return uv
}

// AppendQueryToURL appends the given query string to the given url.
func AppendQueryToURL(reqURL *url.URL, uv url.Values) error {
	urlValues, err := url.ParseQuery(reqURL.RawQuery)
	if err != nil {
		return err
	}

	for key, values := range uv {
		for _, value := range values {
			urlValues.Add(key, value)
		}
	}

	// url.Values format to a sorted "url encoded" string.
	// e.g. "key=val&foo=bar"
	reqURL.RawQuery = urlValues.Encode()
	return nil
}

// AppendQueryToURLString appends the given query data to the given url.
func AppendQueryToURLString(urlStr string, query url.Values) string {
	if len(query) == 0 {
		return urlStr
	}

	if strings.ContainsRune(urlStr, '?') {
		return urlStr + "&" + query.Encode()
	}
	return urlStr + "?" + query.Encode()
}

// MakeBody make request body, convert data to io.Reader
func MakeBody(data any, cType string) io.Reader {
	return ToRequestBody(data, cType)
}

// ToRequestBody make request body, convert data to io.Reader
//
// Allow type for data:
//   - string
//   - []byte
//   - map[string]string
//   - map[string][]string/url.Values
//   - io.Reader(eg: bytes.Buffer, strings.Reader)
func ToRequestBody(data any, cType string) io.Reader {
	if data == nil {
		return nil // nobody
	}

	var reader io.Reader
	kind := httpctype.ToKind(cType, "")

	switch typVal := data.(type) {
	case io.Reader:
		reader = typVal
	case []byte:
		reader = bytes.NewBuffer(typVal)
	case string:
		reader = bytes.NewBufferString(typVal)
	case url.Values:
		reader = bytes.NewBufferString(typVal.Encode())
	case map[string]string:
		if kind == httpctype.KindJSON {
			reader = toJSONReader(data)
		} else {
			reader = bytes.NewBufferString(ToQueryValues(typVal).Encode())
		}
	case map[string][]string:
		if kind == httpctype.KindJSON {
			reader = toJSONReader(data)
		} else {
			reader = bytes.NewBufferString(url.Values(typVal).Encode())
		}
	default:
		// encode body data to json
		if kind == httpctype.KindJSON {
			reader = toJSONReader(data)
		} else {
			panic("httpreq: invalid data type for request body, content-type: " + cType)
		}
	}

	return reader
}

func toJSONReader(data any) io.Reader {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	// close escape  &, <, >  TO  \u0026, \u003c, \u003e
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		panic("encode data as json fail. error=" + err.Error())
	}
	return buf
}

// HeaderToString convert http Header to string
func HeaderToString(h http.Header) string {
	var sb strings.Builder
	for key, values := range h {
		sb.WriteString(key)
		sb.WriteString(": ")
		sb.WriteString(strings.Join(values, ";"))
		sb.WriteByte('\n')
	}
	return sb.String()
}

// RequestToString convert http Request to string
func RequestToString(r *http.Request) string {
	buf := &bytes.Buffer{}
	buf.WriteString(r.Method)
	buf.WriteByte(' ')
	buf.WriteString(r.URL.String())
	buf.WriteByte(' ')
	buf.WriteString(r.Proto)
	buf.WriteByte('\n')

	for key, values := range r.Header {
		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(strings.Join(values, ";"))
		buf.WriteByte('\n')
	}

	if r.Body != nil {
		buf.WriteByte('\n')
		_, _ = buf.ReadFrom(r.Body)
	}
	return buf.String()
}

// ResponseToString convert http Response to string
func ResponseToString(w *http.Response) string {
	if w == nil {
		return ""
	}

	buf := &bytes.Buffer{}
	buf.WriteString(w.Proto)
	buf.WriteByte(' ')
	buf.WriteString(w.Status)
	buf.WriteByte('\n')

	if len(w.Header) > 0 {
		for key, values := range w.Header {
			buf.WriteString(key)
			buf.WriteString(": ")
			buf.WriteString(strings.Join(values, ";"))
			buf.WriteByte('\n')
		}
	}

	if w.Body != nil {
		buf.WriteByte('\n')
		_, _ = buf.ReadFrom(w.Body)
		_ = w.Body.Close()
	}

	return buf.String()
}

// ParseAccept header to strings. referred from gin framework
//
// eg: acceptHeader = "application/json, text/plain, */*"
func ParseAccept(acceptHeader string) []string {
	if acceptHeader == "" {
		return []string{}
	}

	parts := strings.Split(acceptHeader, ",")
	outs := make([]string, 0, len(parts))

	for _, part := range parts {
		if part = strings.TrimSpace(strings.Split(part, ";")[0]); part != "" {
			outs = append(outs, part)
		}
	}
	return outs
}
