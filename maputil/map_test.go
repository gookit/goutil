package maputil_test

import (
	"testing"

	"github.com/gookit/goutil/maputil"
	"github.com/stretchr/testify/assert"
)

func TestMergeStringMap(t *testing.T) {
	ret := maputil.MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"A": "v1"}, false)
	assert.Equal(t, map[string]string{"A": "v0"}, ret)

	ret = maputil.MergeStringMap(map[string]string{"A": "v0"}, map[string]string{"a": "v1"}, true)
	assert.Equal(t, map[string]string{"a": "v0"}, ret)
}
