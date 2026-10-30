package httpreq_test

import (
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/x/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestBindJSONKeepsCachedBody(t *testing.T) {
	const body = `{"name":"Ada","count":2}`
	for _, status := range []int{http.StatusOK, http.StatusCreated} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			r := httpreq.NewResp(&http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(body)), ContentLength: int64(len(body))})
			var first map[string]any
			assert.NoErr(t, r.BindJSONOnOk(&first))
			assert.Eq(t, "Ada", first["name"])
			assert.Eq(t, body, r.BodyString())
			var second struct {
				Name  string
				Count int
			}
			assert.NoErr(t, r.BindJSON(&second))
			assert.Eq(t, "Ada", second.Name)
			assert.Eq(t, 2, second.Count)
			assert.Eq(t, body, r.BodyString())
		})
	}
	t.Run("first-value decoder semantics", func(t *testing.T) {
		body := `{"first":1} {"second":2}`
		r := httpreq.NewResp(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))})
		var got map[string]int
		assert.NoErr(t, r.BindJSON(&got))
		assert.Eq(t, map[string]int{"first": 1}, got)
		assert.Eq(t, body, r.BodyString())
	})
	t.Run("nil destination retains body", func(t *testing.T) {
		r := httpreq.NewResp(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(body))})
		assert.NoErr(t, r.BindJSON(nil))
		assert.Eq(t, body, r.BodyString())
	})
	t.Run("invalid JSON retains bytes", func(t *testing.T) {
		r := httpreq.NewResp(&http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader("invalid"))})
		var got any
		assert.Err(t, r.BindJSON(&got))
		assert.Eq(t, "invalid", r.BodyString())
	})
}
