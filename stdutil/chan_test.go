package stdutil_test

import (
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestGo(t *testing.T) {
	err := goutil.Go(func() error {
		return nil
	})
	assert.NoErr(t, err)
}
