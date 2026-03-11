package httpreq_test

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

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

func TestClient_NewWithTimeout(t *testing.T) {
	cli := httpreq.NewWithTimeout(5000)
	assert.NotNil(t, cli)
	assert.NotNil(t, cli.Doer())
}

func TestClient_NewWithDoer(t *testing.T) {
	customClient := &http.Client{Timeout: 3000 * time.Millisecond}
	cli := httpreq.NewWithDoer(customClient)
	assert.NotNil(t, cli)
	assert.Eq(t, customClient, cli.Doer())
}

func TestClient_SetClient(t *testing.T) {
	cli := httpreq.New()
	originalDoer := cli.Doer()

	newClient := &http.Client{Timeout: 2000 * time.Millisecond}
	ret := cli.SetClient(newClient)

	assert.Eq(t, newClient, cli.Doer())
	assert.Eq(t, cli, ret)
	assert.NotEq(t, originalDoer, cli.Doer())
}

func TestClient_SetTimeout(t *testing.T) {
	cli := httpreq.New()
	ret := cli.SetTimeout(5000)

	assert.NotNil(t, ret)
	assert.Eq(t, cli, ret)

	hc, ok := cli.Doer().(*http.Client)
	assert.True(t, ok)
	assert.Eq(t, 5000*time.Millisecond, hc.Timeout)
}

func TestClient_PostJSON(t *testing.T) {
	cli := httpreq.New(testSrvAddr)
	data := map[string]string{"name": "inhere", "city": "chengdu"}

	resp, err := cli.PostJSON("/post", data)
	assert.NoErr(t, err)
	assert.True(t, httpreq.IsOK(resp.StatusCode))

	rr := testutil.ParseRespToReply(resp)
	assert.Eq(t, "POST", rr.Method)
	assert.StrContains(t, rr.Body, `"name":"inhere"`)
}

func TestClient_WithData(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.WithData("name=inhere&age=18")
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Data)
}

func TestClient_WithBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)
	body := strings.NewReader("test body")

	opt := cli.WithBody(body)
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
}

func TestClient_BytesBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.BytesBody([]byte("bytes body"))
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
}

func TestClient_StringBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.StringBody("string body")
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
}

func TestClient_FormBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.FormBody(map[string]string{"name": "inhere"})
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
}

func TestClient_JSONBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.JSONBody(map[string]string{"name": "inhere"})
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
	assert.Eq(t, httpctype.JSON, opt.ContentType)
}

func TestClient_JSONBytesBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.JSONBytesBody([]byte(`{"name":"inhere"}`))
	assert.NotNil(t, opt)
	assert.NotNil(t, opt.Body)
	assert.Eq(t, httpctype.JSON, opt.ContentType)
}

func TestClient_AnyBody(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	t.Run("string body", func(t *testing.T) {
		opt := cli.AnyBody("string data")
		assert.NotNil(t, opt.Body)
	})

	t.Run("bytes body", func(t *testing.T) {
		opt := cli.AnyBody([]byte("bytes data"))
		assert.NotNil(t, opt.Body)
	})

	t.Run("map body", func(t *testing.T) {
		opt := cli.AnyBody(map[string]string{"key": "value"})
		assert.NotNil(t, opt.Body)
	})
}

func TestClient_WithOption(t *testing.T) {
	cli := httpreq.New(testSrvAddr)

	opt := cli.WithOption(httpreq.WithJSONType)
	assert.NotNil(t, opt)
	assert.Eq(t, httpctype.JSON, opt.ContentType)
}

func TestClient_SendWithOpt(t *testing.T) {
	cli := httpreq.New(testSrvAddr)
	opt := httpreq.NewOpt().WithMethod(http.MethodGet)

	resp, err := cli.SendWithOpt("/get", opt)
	assert.NoErr(t, err)
	assert.True(t, httpreq.IsOK(resp.StatusCode))

	rr := testutil.ParseRespToReply(resp)
	assert.Eq(t, "GET", rr.Method)
}

func TestClient_SendRequest(t *testing.T) {
	cli := httpreq.New(testSrvAddr)
	req, err := http.NewRequest(http.MethodGet, testSrvAddr+"/get", nil)
	assert.NoErr(t, err)

	opt := httpreq.NewOpt()
	resp, err := cli.SendRequest(req, opt)
	assert.NoErr(t, err)
	assert.True(t, httpreq.IsOK(resp.StatusCode))
}

func TestClient_BaseURL(t *testing.T) {
	cli := httpreq.New()
	ret := cli.BaseURL("http://localhost:8080")

	assert.Eq(t, cli, ret)
}

func TestClient_DefaultMethod(t *testing.T) {
	cli := httpreq.New()

	ret := cli.DefaultMethod(http.MethodPost)
	assert.Eq(t, cli, ret)

	ret = cli.DefaultMethod("")
	assert.Eq(t, cli, ret)
}

func TestClient_ContentType(t *testing.T) {
	cli := httpreq.New()

	ret := cli.ContentType(httpctype.JSON)
	assert.Eq(t, cli, ret)
}
