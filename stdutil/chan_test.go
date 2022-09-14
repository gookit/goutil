package stdutil_test

import (
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGo(t *testing.T) {
	err := stdutil.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}
