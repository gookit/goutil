package httpreq

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

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

// AddHeaderMap to reqeust instance.
func AddHeaderMap(req *http.Request, headerMap map[string]string) {
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

// ToQueryValues convert string-map or any-map to url.Values
func ToQueryValues(data any) url.Values {
	// use url.Values directly if we have it
	if uv, ok := data.(url.Values); ok {
		return uv
	}

	uv := make(url.Values)
	if strMp, ok := data.(map[string]string); ok {
		for k, v := range strMp {
			uv.Add(k, v)
		}
	} else if kvMp, ok := data.(map[string]any); ok {
		for k, v := range kvMp {
			uv.Add(k, strutil.QuietString(v))
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

// IsNoBodyMethod check
func IsNoBodyMethod(method string) bool {
	return method != "POST" && method != "PUT" && method != "PATCH"
}

// ToRequestBody convert handle
//
// Allow type for data:
//   - string
//   - []byte
//   - map[string]string
//   - map[string][]string/url.Values
//   - io.Reader(eg: bytes.Buffer, strings.Reader)
func ToRequestBody(data any) io.Reader {
	var reqBody io.Reader
	switch typVal := data.(type) {
	case io.Reader:
		reqBody = typVal
	case map[string]string:
		reqBody = bytes.NewBufferString(ToQueryValues(typVal).Encode())
	case map[string][]string:
		reqBody = bytes.NewBufferString(url.Values(typVal).Encode())
	case url.Values:
		reqBody = bytes.NewBufferString(typVal.Encode())
	case string:
		reqBody = bytes.NewBufferString(typVal)
	case []byte:
		reqBody = bytes.NewBuffer(typVal)
	default:
		// auto encode body data to json
		if data != nil {
			buf := &bytes.Buffer{}
			enc := json.NewEncoder(buf)
			// close escape  &, <, >  TO  \u0026, \u003c, \u003e
			enc.SetEscapeHTML(false)
			if err := enc.Encode(data); err != nil {
				panic("auto encode data error=" + err.Error())
			}

			reqBody = buf
		}

		// nobody
	}

	return reqBody
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
	}

	return buf.String()
}
