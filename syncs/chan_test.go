package syncs_test

import (
	"testing"

	"github.com/gookit/goutil/syncs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGo(t *testing.T) {
	err := syncs.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}
