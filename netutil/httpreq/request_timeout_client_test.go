package httpreq_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gookit/goutil/netutil/httpreq"
)

func TestRequestTimeoutPreservesHTTPClient(t *testing.T) {
	t.Run("TLS transport and original timeout", func(t *testing.T) {
		srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, "trusted") }))
		defer srv.Close()
		hc := srv.Client()
		hc.Timeout = 2 * time.Second
		cli := httpreq.NewWithDoer(hc)
		response, err := cli.Get(srv.URL, func(o *httpreq.Option) { o.Timeout = 1000 })
		if err != nil {
			t.Fatal(err)
		}
		data, err := io.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil || string(data) != "trusted" {
			t.Errorf("response %q %v", data, err)
		}
		if hc.Timeout != 2*time.Second || cli.Doer() != hc {
			t.Error("original HTTP client mutated")
		}
	})
	t.Run("redirect policy", func(t *testing.T) {
		stop := errors.New("stop redirect")
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				http.Redirect(w, r, "/target", http.StatusFound)
			} else {
				_, _ = io.WriteString(w, "redirected")
			}
		}))
		defer srv.Close()
		hc := &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return stop }}
		response, err := httpreq.NewWithDoer(hc).Get(srv.URL, func(o *httpreq.Option) { o.Timeout = 1000 })
		if response != nil && response.Body != nil {
			_ = response.Body.Close()
		}
		if !errors.Is(err, stop) {
			t.Errorf("redirect policy lost: %v", err)
		}
	})
	t.Run("cookie jar", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/set" {
				http.SetCookie(w, &http.Cookie{Name: "session", Value: "present", Path: "/"})
			} else {
				c, e := r.Cookie("session")
				if e != nil {
					http.Error(w, "missing", 400)
					return
				}
				_, _ = io.WriteString(w, c.Value)
			}
		}))
		defer srv.Close()
		jar, _ := cookiejar.New(nil)
		hc := &http.Client{Jar: jar}
		cli := httpreq.NewWithDoer(hc)
		response, err := cli.Get(srv.URL+"/set", func(o *httpreq.Option) { o.Timeout = 1000 })
		if err != nil {
			t.Fatal(err)
		}
		_ = response.Body.Close()
		response, err = cli.Get(srv.URL+"/read", func(o *httpreq.Option) { o.Timeout = 1000 })
		if err != nil {
			t.Fatal(err)
		}
		data, err := io.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil || string(data) != "present" {
			t.Errorf("jar lost: %q %v", data, err)
		}
	})
	t.Run("timeout enforced", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() }))
		defer srv.Close()
		response, err := httpreq.NewWithDoer(srv.Client()).Get(srv.URL, func(o *httpreq.Option) { o.Timeout = 20 })
		if response != nil {
			_ = response.Body.Close()
		}
		if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
			t.Errorf("expected timeout: %v", err)
		}
	})
	t.Run("zero override", func(t *testing.T) {
		srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, "original") }))
		defer srv.Close()
		response, err := httpreq.NewWithDoer(srv.Client()).Get(srv.URL)
		if err != nil {
			t.Fatal(err)
		}
		data, err := io.ReadAll(response.Body)
		_ = response.Body.Close()
		if err != nil || string(data) != "original" {
			t.Errorf("response %q %v", data, err)
		}
	})
}
