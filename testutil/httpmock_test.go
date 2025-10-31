package testutil_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/fakeobj"
)

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.MockHttpServer()
	defer s.Close()

	testSrvAddr = s.HTTPHost()
	s.PrintHttpHost()

	m.Run()
}

func TestNewHTTPRequest(t *testing.T) {
	r := testutil.NewHTTPRequest("GET", testSrvAddr+"/hello", &testutil.MD{
		Headers: map[string]string{
			"X-Test": "val",
		},
		BeforeSend: func(req *http.Request) {
			req.Header.Set("X-Test2", "val2")
		},
	})

	assert.Eq(t, "GET", r.Method)
	assert.Eq(t, testSrvAddr+"/hello", r.URL.String())
	assert.Eq(t, "val", r.Header.Get("X-Test"))
	assert.Eq(t, "val2", r.Header.Get("X-Test2"))

	assert.Panics(t, func() {
		testutil.NewHTTPRequest("invalid", "://", nil)
	})
}

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

func TestEchoReply(t *testing.T) {
	er := testutil.EchoReply{}

	// JSONMap
	assert.Nil(t, er.JSONMap())
	er.JSON = []any{"name", "inhere"}
	assert.Nil(t, er.JSONMap())

	// HeaderString
	assert.Empty(t, er.HeaderString("X-Test"))
	er.Headers = map[string]any{
		"X-Test":  "val",
		"X-Test2": []string{"val2", "ext1"},
	}
	assert.Eq(t, "val", er.HeaderString("X-Test"))
	assert.Eq(t, "[val2 ext1]", er.HeaderString("X-Test2"))
}

func TestNewEchoServer(t *testing.T) {
	r, err := http.Post(testSrvAddr, "text/plain", strings.NewReader("hello!"))
	assert.NoErr(t, err)

	rr := testutil.ParseRespToReply(r)
	// dump.P(rr)
	assert.Eq(t, "POST", rr.Method)
	assert.Eq(t, "text/plain", rr.ContentType())
	assert.Eq(t, "hello!", rr.Body)

	r, err = http.Post(testSrvAddr, "application/json", strings.NewReader(`{"name": "inhere", "age": 18}`))
	assert.NoErr(t, err)

	rr = testutil.ParseRespToReply(r)
	// dump.P(rr)
	assert.Eq(t, "POST", rr.Method)
	assert.Eq(t, "application/json", rr.ContentType())
	assert.Eq(t, "application/json", rr.HeaderString("Content-Type"))
	assert.Eq(t, `{"name": "inhere", "age": 18}`, rr.Body)
	dataMap := rr.JSONMap()
	assert.NotNil(t, dataMap)
	assert.Eq(t, "inhere", dataMap["name"])
	assert.Eq(t, float64(18), dataMap["age"])

	r, err = http.Head(testSrvAddr + "/head")
	assert.NoErr(t, err)
	rr = testutil.ParseRespToReply(r)
	assert.Eq(t, "HEAD", rr.Method)

	rr = testutil.ParseRespToReply(&http.Response{})
	assert.Empty(t, *rr)

	rr = testutil.ParseBodyToReply(nil)
	assert.Empty(t, *rr)

	assert.Panics(t, func() {
		tr := fakeobj.NewStrReader("invalid-json")
		testutil.ParseBodyToReply(tr)
	})

	t.Run("404", func(t *testing.T) {
		r, err = http.Get(testSrvAddr + "/404")
		assert.NoErr(t, err)
		assert.Eq(t, http.StatusNotFound, r.StatusCode)
	})
	t.Run("500", func(t *testing.T) {
		r, err = http.Get(testSrvAddr + "/500")
		assert.NoErr(t, err)
		assert.Eq(t, http.StatusInternalServerError, r.StatusCode)
	})
	t.Run("405", func(t *testing.T) {
		r, err = http.Get(testSrvAddr + "/post")
		assert.NoErr(t, err)
		assert.Eq(t, http.StatusMethodNotAllowed, r.StatusCode)

		// post /get
		r, err = http.Post(testSrvAddr+"/get", "text/plain", strings.NewReader("hello!"))
		assert.NoErr(t, err)
		assert.Eq(t, http.StatusMethodNotAllowed, r.StatusCode)
	})
	t.Run("custom-code", func(t *testing.T) {
		r, err = http.Get(testSrvAddr + "/status-302")
		assert.NoErr(t, err)
		assert.Eq(t, http.StatusFound, r.StatusCode)
	})
}
