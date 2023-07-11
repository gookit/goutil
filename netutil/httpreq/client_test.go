package httpreq_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestClient_Send(t *testing.T) {
	resp, err := httpreq.New(testSrvAddr).
		ContentType(httpctype.JSON).
		DefaultMethod(http.MethodPost).
		DefaultHeaderMap(map[string]string{"coustom1": "value1"}).
		Send("POST", "/json", func(opt *httpreq.Option) {
			opt.Body = io.NopCloser(strings.NewReader(`{"name": "inhere"}`))
		})

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

func TestClient_RSET(t *testing.T) {
	cli := httpreq.New(testSrvAddr).
		ContentType(httpctype.JSON).
		DefaultHeader("custom1", "value1").
		DefaultHeaderMap(map[string]string{
			"custom2": "value2",
		})

	assert.NotNil(t, cli.Doer())

	t.Run("Get", func(t *testing.T) {
		resp, err := cli.Get("/get", httpreq.WithData("name=inhere&age=18"))
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "GET", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.Equal(t, "value2", rr.Headers["Custom2"])
		assert.Equal(t, "inhere", rr.Query["name"])
		// dump.P(rr)
	})

	t.Run("Post", func(t *testing.T) {
		resp, err := cli.Post("/post", `{"name": "inhere"}`, httpreq.WithJSONType)
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "POST", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.StrContains(t, rr.Headers["Content-Type"].(string), httpctype.MIMEJSON)
		assert.Eq(t, `{"name": "inhere"}`, rr.Body)
		dump.P(rr)
	})

	t.Run("Put", func(t *testing.T) {
		resp, err := cli.Put("/put", `{"name": "inhere"}`, httpreq.WithJSONType)
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "PUT", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.StrContains(t, rr.Headers["Content-Type"].(string), httpctype.MIMEJSON)
		assert.Eq(t, `{"name": "inhere"}`, rr.Body)
		// dump.P(rr)
	})

	t.Run("Delete", func(t *testing.T) {
		resp, err := cli.Delete("/delete", httpreq.WithData("name=inhere&age=18"))
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "DELETE", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.Equal(t, "value2", rr.Headers["Custom2"])
		assert.Equal(t, "inhere", rr.Query["name"])
		// dump.P(rr)
	})
}

func TestHttpReq_MustSend(t *testing.T) {
	cli := httpreq.New().OnBeforeSend(func(req *http.Request) {
		assert.Eq(t, http.MethodPost, req.Method)
	}).OnAfterSend(func(resp *http.Response, err error) {
		bodyStr, _ := io.ReadAll(resp.Body)
		assert.StrContains(t, string(bodyStr), "hi,goutil")
	})

	resp := cli.BaseURL(testSrvAddr).
		BytesBody([]byte("hi,goutil")).
		MustSend("POST", "/post")

	sc := resp.StatusCode
	assert.True(t, httpreq.IsOK(sc))
	assert.True(t, httpreq.IsSuccessful(sc))
}
