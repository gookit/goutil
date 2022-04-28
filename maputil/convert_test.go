package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
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

func TestToString(t *testing.T) {
	src := map[string]interface{}{"a": "v0", "b": 23}

	s := maputil.ToString(src)
	dump.P(s)
	assert.Contains(t, s, "b:23")
	assert.Contains(t, s, "a:v0")

	s = maputil.ToString(nil)
	assert.Equal(t, "", s)

	s = maputil.ToString(map[string]interface{}{})
	assert.Equal(t, "{}", s)

	s = maputil.ToString(map[string]interface{}{"": nil})
	assert.Equal(t, "{:}", s)
}
