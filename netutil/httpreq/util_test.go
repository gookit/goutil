package httpreq_test

import (
	"net/http"
	"testing"

	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBuildBasicAuth(t *testing.T) {
	val := httpreq.BuildBasicAuth("inhere", "abcd&123")

	assert.Contains(t, val, "Basic ")
}

func TestAddHeadersToRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "abc.xyz", nil)
	assert.NoErr(t, err)

	httpreq.AddHeadersToRequest(req, http.Header{
		"key0": []string{"val0"},
	})

	assert.Eq(t, "val0", req.Header.Get("key0"))
}
