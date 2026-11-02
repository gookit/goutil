package httpreq_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gookit/goutil/netutil/httpreq"
)

func TestOptionCopyPreservesSource(t *testing.T) {
	t.Run("body", func(t *testing.T) {
		body := strings.NewReader("source payload")
		source := httpreq.NewOpt().WithBody(body)
		copied := source.Copy()
		if source.Body != body {
			t.Error("Copy cleared the source body")
		}
		if copied.Body != nil {
			t.Error("copy retained a consumed request body")
		}
	})
	t.Run("copy headers", func(t *testing.T) {
		source := httpreq.NewOpt().WithHeader("X-Trace", "source")
		copied := source.Copy().WithHeader("X-Trace", "copy").WithHeader("X-New", "copy")
		if source.HeaderMap["X-Trace"] != "source" || source.HeaderMap["X-New"] != "" {
			t.Errorf("copy changed source headers: %v", source.HeaderMap)
		}
		if copied.HeaderMap["X-Trace"] != "copy" {
			t.Error("copy header was not set")
		}
	})
	t.Run("source headers", func(t *testing.T) {
		source := httpreq.NewOpt().WithHeader("X-Trace", "source")
		copied := source.Copy()
		source.WithHeader("X-Trace", "changed")
		if copied.HeaderMap["X-Trace"] != "source" {
			t.Error("source mutation changed copied headers")
		}
	})
	t.Run("nil headers", func(t *testing.T) {
		source := httpreq.NewOpt()
		copied := source.Copy()
		if copied.HeaderMap != nil {
			t.Error("nil headers changed")
		}
		copied.WithHeader("X-New", "copy")
		if source.HeaderMap != nil {
			t.Error("new copy header changed source")
		}
	})
	t.Run("metadata", func(t *testing.T) {
		ctx := context.Background()
		source := &httpreq.Option{Context: ctx, Method: http.MethodPost, Timeout: 1234, Data: "data", ContentType: "text/plain"}
		copied := source.Copy()
		if copied.Context != ctx || copied.Method != source.Method || copied.Timeout != source.Timeout || copied.Data != source.Data || copied.ContentType != source.ContentType {
			t.Error("copied metadata changed")
		}
	})
}

func TestOptionCopyUsesIndependentRequests(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		_, _ = io.WriteString(w, r.Header.Get("X-Trace")+":"+string(body))
	}))
	defer srv.Close()
	source := httpreq.New(srv.URL).StringBody("original").WithHeader("X-Trace", "source")
	copied := source.Copy().WithHeader("X-Trace", "copy")
	response, err := source.Send(http.MethodPost, "")
	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil || string(got) != "source:original" {
		t.Errorf("source request = %q, %v", got, err)
	}
	response, err = copied.Get("")
	if err != nil {
		t.Fatal(err)
	}
	got, err = io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil || string(got) != "copy:" {
		t.Errorf("copy request = %q, %v", got, err)
	}
}
