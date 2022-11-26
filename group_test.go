package goutil_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewErrGroup(t *testing.T) {
	httpreq.ConfigStd(func(hc *http.Client) {
		hc.Timeout = 3 * time.Second
	})

	eg := goutil.NewErrGroup()
	eg.Add(func() error {
		resp, err := httpreq.Get("https://httpbin.org/get", nil)
		if err != nil {
			return err
		}

		fmt.Println(resp.Body)
		return nil
	}, func() error {
		resp, err := httpreq.Post("https://httpbin.org/post", "hi", nil)
		if err != nil {
			return err
		}

		fmt.Println(resp.Body)
		return nil
	})

	err := eg.Wait()
	assert.NoErr(t, err)
}
