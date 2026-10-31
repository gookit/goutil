package httpreq_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/x/assert"
)

func TestReadBodyRetainsTransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", "100")
		_, _ = fmt.Fprint(w, `{"ok":true}`)
	}))
	defer server.Close()
	get := func() *httpreq.RespX {
		response, err := http.Get(server.URL)
		if err != nil {
			t.Fatal(err)
		}
		return httpreq.NewResp(response)
	}
	t.Run("repeated reads", func(t *testing.T) {
		response := get()
		for i := 0; i < 2; i++ {
			if err := response.ReadBody(); !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("ReadBody #%d error = %v, want unexpected EOF", i+1, err)
			}
		}
	})
	t.Run("JSON cannot hide transport error", func(t *testing.T) {
		response := get()
		for i := 0; i < 2; i++ {
			var value map[string]bool
			if err := response.BindJSON(&value); !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("BindJSON #%d error = %v, want unexpected EOF", i+1, err)
			}
			if value != nil {
				t.Fatalf("decoded truncated response: %v", value)
			}
		}
	})
	t.Run("buffer access retains failure", func(t *testing.T) {
		response := get()
		if err := response.ReadBody(); !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatal(err)
		}
		defer func() {
			err, ok := recover().(error)
			if !ok || !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("BodyBuffer panic = %v, want unexpected EOF", err)
			}
		}()
		response.BodyBuffer()
	})
	t.Run("clear buffer resets failure", func(t *testing.T) {
		response := get()
		if err := response.ReadBody(); !errors.Is(err, io.ErrUnexpectedEOF) {
			t.Fatal(err)
		}
		response.CloseBuffer()
		good := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = fmt.Fprint(w, `{"ok":true}`) }))
		defer good.Close()
		fresh, err := http.Get(good.URL)
		assert.NoErr(t, err)
		response.Response = fresh
		var value map[string]bool
		assert.NoErr(t, response.BindJSON(&value))
		if !value["ok"] {
			t.Fatalf("value = %v", value)
		}
	})
}
