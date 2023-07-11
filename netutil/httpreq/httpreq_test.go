package httpreq_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.NewEchoServer()
	defer s.Close()
	testSrvAddr = "http://" + s.Listener.Addr().String()
	fmt.Println("Test server listen on:", testSrvAddr)

	m.Run()
}

func TestStdClient(t *testing.T) {
	assert.NotNil(t, httpreq.Std())
	httpreq.SetTimeout(300)
	httpreq.Config(func(hc *http.Client) {
		hc.Timeout = 400 * time.Millisecond
	})

	t.Run("head", func(t *testing.T) {
		resp, err := httpreq.Send("head", testSrvAddr+"/head")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseRespToReply(resp)
		// dump.P(rr)
		assert.Equal(t, "HEAD", rr.Method)
	})

	t.Run("options", func(t *testing.T) {
		resp := httpreq.MustSend("options", testSrvAddr+"/options")
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseRespToReply(resp)
		dump.P(rr)
		assert.Eq(t, "OPTIONS", rr.Method)
	})

	t.Run("get", func(t *testing.T) {
		resp := httpreq.MustResp(httpreq.Get(testSrvAddr + "/get"))
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Equal(t, "GET", rr.Method)
	})

	t.Run("post", func(t *testing.T) {
		resp, err := httpreq.Post(testSrvAddr+"/post", "hi")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Equal(t, "POST", rr.Method)
		assert.Equal(t, "hi", rr.Body)
	})

	t.Run("post JSON", func(t *testing.T) {
		resp, err := httpreq.PostJSON(testSrvAddr+"/post", map[string]string{"name": "inhere"})
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Equal(t, "POST", rr.Method)
		assert.StrContains(t, rr.Body, `{"name":"inhere"}`)
	})

	t.Run("put", func(t *testing.T) {
		resp, err := httpreq.Put(testSrvAddr+"/put", "hi")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Equal(t, "PUT", rr.Method)
		assert.Equal(t, "hi", rr.Body)
	})

	t.Run("delete", func(t *testing.T) {
		resp, err := httpreq.Delete(testSrvAddr + "/delete")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Equal(t, "DELETE", rr.Method)
	})
}
