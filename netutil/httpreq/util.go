package httpreq

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
)

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
// copied from net/http.
func BuildBasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

// AddHeadersToRequest adds the key, value pairs from the given http.Header to the
// request. Values for existing keys are appended to the keys values.
func AddHeadersToRequest(req *http.Request, header http.Header) {
	for key, values := range header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
}

// ToQueryValues convert string map to url.Values
func ToQueryValues(data interface{}) url.Values {
	// use url.Values directly if we have it
	if uv, ok := data.(url.Values); ok {
		return uv
	}

	uv := make(url.Values)
	if strMp, ok := data.(map[string]string); ok {
		for k, v := range strMp {
			uv.Add(k, v)
		}
	}

	return uv
}

// RequestToString convert http Request to string
func RequestToString(r *http.Request) string {
	buf := &bytes.Buffer{}
	buf.WriteString(r.Method)
	buf.WriteByte(' ')
	buf.WriteString(r.URL.String())
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

	for key, values := range w.Header {
		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(strings.Join(values, ";"))
		buf.WriteByte('\n')
	}

	if w.Body != nil {
		buf.WriteByte('\n')
		_, _ = buf.ReadFrom(w.Body)
	}

	return buf.String()
}
