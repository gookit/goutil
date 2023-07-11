package httpreq_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestOptions_build(t *testing.T) {
	opt := httpreq.OptOrNew(nil)
	opt.BytesBody([]byte("name=inhere"))
	assert.Equal(t, []byte("name=inhere"), fsutil.ReadAll(opt.Body))

	opt.StringBody("name=inhere")
	assert.Equal(t, []byte("name=inhere"), fsutil.ReadAll(opt.Body))

	// FormBody
	opt.FormBody(map[string]string{"name": "inhere"})
	assert.Equal(t, []byte("name=inhere"), fsutil.ReadAll(opt.Body))

	// WithJSON
	opt.WithJSON(map[string]string{"name": "inhere"})
	assert.Equal(t, httpctype.JSON, opt.ContentType)
	assert.Equal(t, `{"name":"inhere"}
`, fsutil.ReadString(opt.Body))
}

func TestOption_Send(t *testing.T) {
	opt := httpreq.OptOrNew(nil)
	assert.NotNil(t, opt)

	opt = httpreq.NewOpt(httpreq.WithJSONType)
	assert.NotNil(t, opt)
	assert.Eq(t, httpctype.JSON, opt.ContentType)

	resp, err := httpreq.New(testSrvAddr).
		StringBody(`{"name": "inhere"}`).
		WithContentType(httpctype.JSON).
		WithHeaderMap(map[string]string{"coustom1": "value1"}).
		Send("POST", "/json")

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

func TestOptions_REST(t *testing.T) {
	opt := httpreq.New(testSrvAddr).
		WithOption().
		WithContentType(httpctype.Form).
		WithHeader("custom1", "value1").
		WithHeaderMap(map[string]string{
			"custom2": "value2",
		})

	t.Run("Get", func(t *testing.T) {
		resp, err := opt.Copy().Get("/get", httpreq.WithData("name=inhere&age=18"))
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "GET", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.Equal(t, "value2", rr.Headers["Custom2"])
	})

	t.Run("Post", func(t *testing.T) {
		resp, err := opt.Copy().Post("/post", nil, httpreq.WithData("name=inhere&age=18"))
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		assert.Equal(t, "POST", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.Equal(t, "value2", rr.Headers["Custom2"])
		dump.P(rr)
	})

	t.Run("Put", func(t *testing.T) {
		resp, err := opt.Copy().WithData("name=inhere&age=18").Put("/put", nil)
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))

		rr := testutil.ParseRespToReply(resp)
		// dump.P(rr)
		assert.Equal(t, "PUT", rr.Method)
		assert.Equal(t, "value1", rr.Headers["Custom1"])
		assert.Equal(t, "value2", rr.Headers["Custom2"])
		assert.NotEmpty(t, rr.Form)
		assert.NotEmpty(t, rr.Body)
	})

	t.Run("Delete", func(t *testing.T) {
		resp, err := opt.Copy().Delete("/delete", httpreq.WithData("name=inhere&age=18"))
		assert.NoErr(t, err)
		sc := resp.StatusCode
		assert.True(t, httpreq.IsOK(sc))
		assert.True(t, httpreq.IsSuccessful(sc))
	})
}
