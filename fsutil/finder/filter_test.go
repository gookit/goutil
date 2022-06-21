package finder_test

import (
	"testing"

	"github.com/gookit/goutil/fsutil/finder"
	"github.com/stretchr/testify/assert"
)

func TestFilterFunc(t *testing.T) {
	fn := finder.FilterFunc(func(filePath, filename string) bool {
		return false
	})

	assert.False(t, fn("path/some.txt", "some.txt"))
}
