package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/x/assert"
)

func TestToByteSizeUnitWithoutNumber(t *testing.T) {
	for _, input := range []string{"B", "b", " B "} {
		t.Run(input, func(t *testing.T) {
			_, err := strutil.ToByteSize(input)
			assert.Error(t, err)
		})
	}
}
