package testutil_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMockRequest(t *testing.T) {
	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello!"))

		if r.Body != nil {
			bs, _ := ioutil.ReadAll(r.Body)
			_, _ = w.Write(bs)
		}
	}))

	w := testutil.MockRequest(r, "GET", "/", nil)
	assert.Eq(t, "hello!", w.Body.String())

	w = testutil.MockRequest(r, "POST", "/", &testutil.MD{BodyString: "body"})
	assert.Eq(t, "hello!body", w.Body.String())

	w = testutil.MockRequest(r, "POST", "/", &testutil.MD{Body: strings.NewReader("BODY")})
	assert.Eq(t, "hello!BODY", w.Body.String())
}
