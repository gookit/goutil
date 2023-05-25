package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFilterFunc(t *testing.T) {
	fn := finder.FilterFunc(func(el finder.Elem) bool {
		return false
	})

	assert.False(t, fn(finder.NewElem("path/some.txt", &testutil.DirEnt{})))
}
