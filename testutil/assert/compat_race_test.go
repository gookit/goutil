package assert_test

import (
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

// Run with -race: assertions in parallel tests must not write shared config.
func TestCompatParallelAssertionsDoNotRace(t *testing.T) {
	for i := 0; i < 8; i++ {
		t.Run("parallel", func(t *testing.T) {
			t.Parallel()
			for j := 0; j < 100; j++ {
				assert.Eq(t, j, j)
				assert.True(t, true)
				assert.NoErr(t, nil)
			}
		})
	}
}
