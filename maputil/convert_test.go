package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestKeyToLower(t *testing.T) {
	src := map[string]string{"A": "v0"}
	ret := maputil.KeyToLower(src)

	assert.Contains(t, ret, "a")
	assert.NotContains(t, ret, "A")
}

func TestToStringMap(t *testing.T) {
	src := map[string]interface{}{"a": "v0", "b": 23}
	ret := maputil.ToStringMap(src)

	assert.Equal(t, ret["a"], "v0")
	assert.Equal(t, ret["b"], "23")
}
