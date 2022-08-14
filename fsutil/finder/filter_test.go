package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFilterFunc(t *testing.T) {
	fn := finder.FilterFunc(func(filePath, filename string) bool {
		return false
	})

	assert.False(t, fn("path/some.txt", "some.txt"))
}
