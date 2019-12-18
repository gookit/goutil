package testutil_test

import (
	"net/http"
	"testing"

	"github.com/gookit/goutil/testutil"
	"github.com/stretchr/testify/assert"
)

func TestMockRequest(t *testing.T) {
	r := http.NewServeMux()
	r.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello"))
	}))

	w := testutil.MockRequest(r, "GET", "/", nil)
	assert.Equal(t, "hello", w.Body.String())
}
