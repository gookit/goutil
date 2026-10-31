package maputil_test

import (
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/x/assert"
	"testing"
)

func TestMergeStringMapEmptyDestinationIgnoreCase(t *testing.T) {
	for _, alias := range []bool{false, true} {
		for _, nilDst := range []bool{false, true} {
			name := "direct_allocated"
			if alias {
				name = "alias_allocated"
			}
			if nilDst {
				name += "_nil"
			}
			t.Run(name, func(t *testing.T) {
				src := map[string]string{"HOST": "localhost", "Port": "8080"}
				dst := map[string]string{}
				if nilDst {
					dst = nil
				}
				var got map[string]string
				if alias {
					got = maputil.MergeSMap(src, dst, true)
				} else {
					got = maputil.MergeStringMap(src, dst, true)
				}
				assert.Eq(t, map[string]string{"host": "localhost", "port": "8080"}, got)
				assert.Eq(t, map[string]string{"HOST": "localhost", "Port": "8080"}, src)
			})
		}
	}
	t.Run("case sensitive", func(t *testing.T) {
		assert.Eq(t, map[string]string{"HOST": "localhost"}, maputil.MergeStringMap(map[string]string{"HOST": "localhost"}, nil, false))
	})
	t.Run("populated destination", func(t *testing.T) {
		assert.Eq(t, map[string]string{"host": "new", "other": "kept"}, maputil.MergeStringMap(map[string]string{"HOST": "new"}, map[string]string{"host": "old", "other": "kept"}, true))
	})
	t.Run("empty source", func(t *testing.T) {
		assert.Eq(t, map[string]string{"HOST": "kept"}, maputil.MergeStringMap(nil, map[string]string{"HOST": "kept"}, true))
	})
}
