package httpreq_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/netutil/httpctype"
	"github.com/gookit/goutil/netutil/httpreq"
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
	s = rx.BodyString()
	// fmt.Println(s)
	assert.StrContains(t, s, `"hi"`)
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
}
