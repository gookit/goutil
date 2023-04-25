package httpreq

import (
	"net/http"
	"sync"
	"time"
)

// default std client
var std = NewClient(500)

// global lock
var _gl = sync.Mutex{}

// client map
var cs = map[int]*ReqClient{}

// NewClient create a new http client
func NewClient(timeout int) *ReqClient {
	_gl.Lock()
	cli, ok := cs[timeout]

	if !ok {
		cli = New().Client(&http.Client{
			Timeout: time.Duration(timeout) * time.Millisecond,
		})
		cs[timeout] = cli
	}

	_gl.Unlock()
	return cli
}
