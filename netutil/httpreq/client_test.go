package httpreq_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil/assert"
)

func TestHttpReq_Send(t *testing.T) {
	resp, err := httpreq.New("https://httpbin.org").
		StringBody("hi").
		ContentType(httpctype.JSON).
		WithHeaders(map[string]string{"coustom1": "value1"}).
		Send("/get")

	assert.NoErr(t, err)
	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
	assert.False(t, httpreq.IsRedirect(sc))
	assert.False(t, httpreq.IsForbidden(sc))
	assert.False(t, httpreq.IsNotFound(sc))
	assert.False(t, httpreq.IsClientError(sc))
	assert.False(t, httpreq.IsServerError(sc))

	retMp := make(map[string]any)
	err = jsonutil.DecodeReader(resp.Body, &retMp)
	assert.NoErr(t, err)
	dump.P(retMp)
}

func TestHttpReq_MustSend(t *testing.T) {
	cli := httpreq.New().OnBeforeSend(func(req *http.Request) {
		assert.Eq(t, http.MethodPost, req.Method)
	}).OnAfterSend(func(resp *http.Response) {
		bodyStr, _ := io.ReadAll(resp.Body)
		assert.StrContains(t, string(bodyStr), "hi,goutil")
	})

	resp := cli.
		BaseURL("https://httpbin.org").
		BytesBody([]byte("hi,goutil")).
		Method("POST").
		MustSend("/post")

	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
}
