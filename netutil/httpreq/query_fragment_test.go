package httpreq_test

import (
	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/x/assert"
	"net/url"
	"testing"
)

func TestAppendQueryBeforeFragment(t *testing.T) {
	for _, tt := range []struct{ name, in, want string }{
		{"fragment", "https://example.test/path#section", "https://example.test/path?lang=go#section"},
		{"existing query", "https://example.test/path?old=1#section", "https://example.test/path?old=1&lang=go#section"},
		{"question in fragment", "https://example.test/path#section?inside=1", "https://example.test/path?lang=go#section?inside=1"},
		{"empty fragment", "https://example.test/path#", "https://example.test/path?lang=go#"},
		{"relative", "path#section", "path?lang=go#section"},
		{"escaped hash", "https://example.test/path%23part#section", "https://example.test/path%23part?lang=go#section"},
		{"ordinary", "https://example.test/path", "https://example.test/path?lang=go"},
		{"ordinary query", "https://example.test/path?old=1", "https://example.test/path?old=1&lang=go"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Eq(t, tt.want, httpreq.AppendQueryToURLString(tt.in, url.Values{"lang": {"go"}}))
		})
	}
	t.Run("no new query", func(t *testing.T) {
		in := "https://example.test/path#section"
		assert.Eq(t, in, httpreq.AppendQueryToURLString(in, nil))
	})
}
