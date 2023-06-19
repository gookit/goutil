package testutil_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMockRequest(t *testing.T) {
	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello!"))

		if r.Body != nil {
			bs, _ := io.ReadAll(r.Body)
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

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.NewEchoServer()
	defer s.Close()

	testSrvAddr = "http://" + s.Listener.Addr().String()
	fmt.Println("server addr:", testSrvAddr)

	m.Run()
}

func TestNewEchoServer(t *testing.T) {
	r, err := http.Post(testSrvAddr, "text/plain", strings.NewReader("hello!"))
	assert.NoErr(t, err)

	rpl := testutil.ParseRespToReply(r)
	dump.P(rpl)
	assert.Eq(t, "POST", rpl.Method)
	assert.Eq(t, "text/plain", rpl.ContentType())
	assert.Eq(t, "hello!", rpl.Body)

	r, err = http.Post(testSrvAddr, "application/json", strings.NewReader(`{"name": "inhere", "age": 18}`))
	assert.NoErr(t, err)

	rpl = testutil.ParseRespToReply(r)
	dump.P(rpl)
	assert.Eq(t, "POST", rpl.Method)
	assert.Eq(t, "application/json", rpl.ContentType())
	assert.Eq(t, `{"name": "inhere", "age": 18}`, rpl.Body)
}
