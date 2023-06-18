# HTTP Request

`httpreq` provide an simple http requester and some useful util functions.

- provide a simple and useful HTTP request client
- provide some useful http utils functions

## Install

```bash
go get github.com/gookit/goutil/netutil/httpreq
```

## Go docs

- [Go docs](https://pkg.go.dev/github.com/gookit/goutil/netutil/httpreq)

## Usage

```go
package main

import (
	"fmt"

	"github.com/gookit/goutil/netutil/httpreq"
)

func main() {
	// Send a GET request
	resp, err := httpreq.Get("http://httpbin.org/get")
	fmt.Println(httpreq.ResponseToString(resp), err)

	// Send a POST request
	resp, err = httpreq.Post("http://httpbin.org/post", `{"name":"inhere"}`, httpreq.WithJSONType)
	fmt.Println(httpreq.ResponseToString(resp), err)
}
```

## HTTP Client

```go
package main

import (
    "fmt"

    "github.com/gookit/goutil/netutil/httpreq"
)

func main() {
    // create a client
    client := httpreq.New("http://httpbin.org")

    // Send a GET request
    resp, err := client.Get("/get")
    fmt.Println(httpreq.ResponseToString(resp), err)

    // Send a POST request
    resp, err = client.Post("/post", `{"name":"inhere"}`, httpreq.WithJSONType)
    fmt.Println(httpreq.ResponseToString(resp), err)
}
```

## Package docs

```go

func AddHeaderMap(req *http.Request, headerMap map[string]string)
func AddHeaders(req *http.Request, header http.Header)
func AppendQueryToURL(reqURL *url.URL, uv url.Values) error
func AppendQueryToURLString(urlStr string, query url.Values) string
func BuildBasicAuth(username, password string) string
func Config(fn func(hc *http.Client))
func Delete(url string, optFns ...OptionFn) (*http.Response, error)
func Get(url string, optFns ...OptionFn) (*http.Response, error)
func HeaderToString(h http.Header) string
func HeaderToStringMap(rh http.Header) map[string]string
func IsClientError(statusCode int) bool
func IsForbidden(statusCode int) bool
func IsNoBodyMethod(method string) bool
func IsNotFound(statusCode int) bool
func IsOK(statusCode int) bool
func IsRedirect(statusCode int) bool
func IsServerError(statusCode int) bool
func IsSuccessful(statusCode int) bool
func MakeBody(data any, cType string) io.Reader
func MakeQuery(data any) url.Values
func MustResp(r *http.Response, err error) *http.Response
func MustSend(method, url string, optFns ...OptionFn) *http.Response
func Post(url string, data any, optFns ...OptionFn) (*http.Response, error)
func Put(url string, data any, optFns ...OptionFn) (*http.Response, error)
func RequestToString(r *http.Request) string
func ResponseToString(w *http.Response) string
func Send(method, url string, optFns ...OptionFn) (*http.Response, error)
func SendRequest(req *http.Request, opt *Option) (*http.Response, error)
func SetTimeout(ms int)
func ToQueryValues(data any) url.Values
func ToRequestBody(data any, cType string) io.Reader
func WithJSONType(opt *Option)
type AfterSendFn func(resp *http.Response, err error)
type BasicAuthConf struct{ ... }
type Client struct{ ... }
    func New(baseURL ...string) *Client
    func NewClient(timeout int) *Client
    func NewWithDoer(d Doer) *Client
    func Std() *Client
type Option struct{ ... }
    func MakeOpt(opt *Option) *Option
    func NewOpt(fns ...OptionFn) *Option
    func NewOption(fns []OptionFn) *Option
type OptionFn func(opt *Option)
    func WithData(data any) OptionFn
type RespX struct{ ... }
    func MustRespX(r *http.Response, err error) *RespX
    func NewResp(hr *http.Response) *RespX
```

## Testings

```shell
go test -v ./netutil/httpreq/...
```

Test limit by regexp:

```shell
go test -v -run ^TestSetByKeys ./netutil/httpreq/...
```
