package httpreq

import "net/http"

// default standard client instance
var std = NewClient(500)

// Std instance
func Std() *Client { return std }

// SetTimeout set default timeout(ms) for std client
//
// Note: timeout unit is millisecond
func SetTimeout(ms int) {
	std = NewClient(ms)
}

// Config std http client
func Config(fn func(hc *http.Client)) {
	fn(std.client.(*http.Client))
}

//
// send request by default client
//

// Get quick send a GET request by default client
func Get(url string, optFns ...OptionFn) (*http.Response, error) {
	return std.Get(url, optFns...)
}

// Post quick send a POST request by default client
func Post(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	return std.Post(url, data, optFns...)
}

// Put quick send a PUT request by default client
func Put(url string, data any, optFns ...OptionFn) (*http.Response, error) {
	return std.Put(url, data, optFns...)
}

// Delete quick send a DELETE request by default client
func Delete(url string, optFns ...OptionFn) (*http.Response, error) {
	return std.Delete(url, optFns...)
}

// Send quick send a request by default client
func Send(method, url string, optFns ...OptionFn) (*http.Response, error) {
	return std.Send(method, url, optFns...)
}

// MustSend quick send a request by default client
func MustSend(method, url string, optFns ...OptionFn) *http.Response {
	return std.MustSend(method, url, optFns...)
}

// SendRequest quick send a request by default client
func SendRequest(req *http.Request, opt *Option) (*http.Response, error) {
	return std.SendRequest(req, opt)
}
