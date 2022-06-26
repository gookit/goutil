package httphelper_test

import (
	"net/http"
	"testing"

	"github.com/gookit/goutil/netutil/httphelper"
	"github.com/stretchr/testify/assert"
)

func TestHeaderToStringMap(t *testing.T) {
	assert.Nil(t, httphelper.HeaderToStringMap(nil))
	assert.Nil(t, httphelper.HeaderToStringMap(http.Header{}))

	want := map[string]string{"key": "value; more"}
	assert.Equal(t, want, httphelper.HeaderToStringMap(http.Header{
		"key": {"value", "more"},
	}))
}
