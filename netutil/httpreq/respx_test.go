package httpreq_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRespX_String(t *testing.T) {
	rx := httpreq.MustRespX(httpreq.Get(testSrvAddr + "/get"))
	assert.NotNil(t, rx)
	assert.Equal(t, 200, rx.StatusCode)
	assert.True(t, rx.IsOk())
	assert.True(t, rx.IsSuccessful())
	assert.False(t, rx.IsFail())
	assert.False(t, rx.IsEmptyBody())
	assert.Equal(t, httpctype.JSON, rx.ContentType())

	s := rx.String()
	fmt.Println(s)
	assert.StrContains(t, s, "GET")

	rx = httpreq.MustRespX(httpreq.Post(testSrvAddr+"/post", "hi"))
	assert.NotNil(t, rx)
	assert.True(t, rx.IsOk())

	// BodyString
	s = rx.BodyString()
	assert.NoErr(t, rx.CloseBody())
	// fmt.Println(s)
	assert.StrContains(t, s, `"hi"`)

	// BindJSONOnOk
	bd := &testutil.EchoReply{}
	assert.NoError(t, rx.BindJSONOnOk(bd))
	assert.Eq(t, "/post", bd.URL)
	assert.Eq(t, "hi", bd.Body)

	assert.NoError(t, rx.BindJSON(nil))
	assert.NoError(t, rx.BindJSONOnOk(nil))

	rx.CloseBuffer()
}

func TestWrapResp(t *testing.T) {
	rx, err := httpreq.WrapResp(httpreq.Get(testSrvAddr + "/get"))
	assert.NoErr(t, err)
	assert.NotNil(t, rx)
	assert.Equal(t, 200, rx.StatusCode)
	assert.True(t, rx.IsOk())
	assert.True(t, rx.IsSuccessful())
	assert.False(t, rx.IsFail())
	assert.False(t, rx.IsEmptyBody())
	assert.Equal(t, httpctype.JSON, rx.ContentType())

	s := rx.String()
	fmt.Println(s)
	assert.StrContains(t, s, "GET")

	rx, err = httpreq.WrapResp(&http.Response{}, errors.New("a error"))
	assert.Nil(t, rx)
	assert.Err(t, err)
}
