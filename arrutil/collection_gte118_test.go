package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMap(t *testing.T) {
	list1 := []map[string]any{
		{"name": "tom", "age": 23},
		{"name": "john", "age": 34},
	}

	flatArr := arrutil.Column(list1, func(obj map[string]any) (val any, find bool) {
		return obj["age"], true
	})

	assert.NotEmpty(t, flatArr)
	assert.Contains(t, flatArr, 23)
	assert.Len(t, flatArr, 2)
	assert.Eq(t, 34, flatArr[1])
}
