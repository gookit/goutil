package httpreq_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gookit/goutil/netutil/httpreq"
)

type fixedRoundTripper struct{}

func (fixedRoundTripper) RoundTrip(*http.Request) (*http.Response, error) { return nil, nil }

func TestSetMaxIdleConnsTransport(t *testing.T) {
	t.Run("default client", func(t *testing.T) {
		global := http.DefaultTransport.(*http.Transport)
		max, perHost := global.MaxIdleConns, global.MaxIdleConnsPerHost
		cli := httpreq.New()
		cli.SetMaxIdleConns(7, 3)
		transport, ok := cli.Doer().(*http.Client).Transport.(*http.Transport)
		if !ok || transport == global {
			t.Fatal("default transport was not cloned")
		}
		defer transport.CloseIdleConnections()
		if transport.MaxIdleConns != 7 || transport.MaxIdleConnsPerHost != 3 {
			t.Error("limits not configured")
		}
		if global.MaxIdleConns != max || global.MaxIdleConnsPerHost != perHost {
			t.Error("global default mutated")
		}
	})
	t.Run("custom transport", func(t *testing.T) {
		transport := &http.Transport{}
		cli := httpreq.NewWithDoer(&http.Client{Transport: transport})
		cli.SetMaxIdleConns(8, 4)
		if cli.Doer().(*http.Client).Transport != transport || transport.MaxIdleConns != 8 || transport.MaxIdleConnsPerHost != 4 {
			t.Error("custom transport not preserved/configured")
		}
	})
	t.Run("other round tripper", func(t *testing.T) {
		rt := fixedRoundTripper{}
		cli := httpreq.NewWithDoer(&http.Client{Transport: rt})
		cli.SetMaxIdleConns(8, 4)
		if cli.Doer().(*http.Client).Transport != rt {
			t.Error("custom RoundTripper changed")
		}
	})
	t.Run("other doer", func(t *testing.T) {
		cli := httpreq.NewWithDoer(httpreq.DoerFunc(func(*http.Request) (*http.Response, error) { return nil, nil }))
		cli.SetMaxIdleConns(8, 4)
	})
	t.Run("real requests", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, "ok") }))
		defer srv.Close()
		cli := httpreq.New()
		cli.SetMaxIdleConns(5, 2)
		transport := cli.Doer().(*http.Client).Transport.(*http.Transport)
		defer transport.CloseIdleConnections()
		for i := 0; i < 3; i++ {
			response, err := cli.Get(srv.URL)
			if err != nil {
				t.Fatal(err)
			}
			data, err := io.ReadAll(response.Body)
			_ = response.Body.Close()
			if err != nil || string(data) != "ok" {
				t.Fatalf("response %q %v", data, err)
			}
		}
	})
}
