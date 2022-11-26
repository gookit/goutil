package httpreq_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBuildBasicAuth(t *testing.T) {
	val := httpreq.BuildBasicAuth("inhere", "abcd&123")

	assert.Contains(t, val, "Basic ")
}

func TestAddHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "inhere.xyz", nil)
	assert.NoErr(t, err)

	httpreq.AddHeaders(req, http.Header{
		"key0": []string{"val0"},
	})

	assert.Eq(t, "val0", req.Header.Get("key0"))
}

func TestHeaderToStringMap(t *testing.T) {
	assert.Nil(t, httpreq.HeaderToStringMap(nil))
	assert.Nil(t, httpreq.HeaderToStringMap(http.Header{}))

	want := map[string]string{"key": "value; more"}
	assert.Eq(t, want, httpreq.HeaderToStringMap(http.Header{
		"key": {"value", "more"},
	}))
}

func TestToQueryValues(t *testing.T) {
	vs := httpreq.ToQueryValues(map[string]string{"field1": "value1", "field2": "value2"})
	assert.StrContains(t, vs.Encode(), "field=value1")

	vs = httpreq.ToQueryValues(map[string]any{"field1": 234, "field2": "value2"})
	assert.StrContains(t, vs.Encode(), "field=234")
}

func TestRequestToString(t *testing.T) {
	req, err := http.NewRequest("GET", "inhere.xyz", nil)
	assert.NoErr(t, err)

	httpreq.AddHeaders(req, http.Header{
		"custom-key0": []string{"val0"},
	})

	vs := httpreq.ToQueryValues(map[string]string{"field1": "value1", "field2": "value2"})

	req.Body = io.NopCloser(strings.NewReader(vs.Encode()))

	str := httpreq.RequestToString(req)
	dump.P(str)

	assert.StrContains(t, str, "GET inhere.xyz")
	assert.StrContains(t, str, "Custom-Key0: val0")
	assert.StrContains(t, str, "field1=value1")
}

func TestResponseToString(t *testing.T) {
	res := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: 50,
		Header: http.Header{
			"Foo": []string{"Bar"},
		},
		Body: io.NopCloser(strings.NewReader("foo...bar")),
	}

	str := httpreq.ResponseToString(res)
	dump.P(str)

	assert.StrContains(t, str, "HTTP/1.1 200 OK")
	assert.StrContains(t, str, "Foo: Bar")
	assert.StrContains(t, str, "foo...bar")
}
