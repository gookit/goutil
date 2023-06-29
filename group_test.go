package goutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewErrGroup(t *testing.T) {
	httpreq.SetTimeout(3000)

	eg := goutil.NewErrGroup()
	eg.Add(func() error {
		resp, err := httpreq.Get(testSrvAddr + "/get")
		if err != nil {
			return err
		}

		fmt.Println(testutil.ParseBodyToReply(resp.Body))
		return nil
	}, func() error {
		resp := httpreq.MustResp(httpreq.Post(testSrvAddr+"/post", "hi"))
		fmt.Println(testutil.ParseBodyToReply(resp.Body))
		return nil
	})

	err := eg.Wait()
	assert.NoErr(t, err)
}

func TestQuickRun_methods(t *testing.T) {
	qr := goutil.NewQuickRun()
	qr.Add(func(ctx *structs.Data) error {
		resp := httpreq.MustResp(httpreq.Get(testSrvAddr + "/get"))
		rr := testutil.ParseBodyToReply(resp.Body)
		assert.Eq(t, "GET", rr.Method)
		return nil
	})

	assert.NoErr(t, qr.Run())
}
