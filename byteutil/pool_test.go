package byteutil_test

import (
	"sync"
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewChanPool(t *testing.T) {
	p := byteutil.NewChanPool(10, 8, 8)

	assert.Equal(t, 8, p.Width())
	assert.Equal(t, 8, p.WidthCap())

	p.Put([]byte("abc"))
	assert.Equal(t, []byte("abc"), p.Get())

	// test concurrent get and put
	t.Run("concurrent", func(t *testing.T) {
		p := byteutil.NewChanPool(10, 8, 8)
		wg := sync.WaitGroup{}

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				p.Put([]byte("abc"))
				assert.Equal(t, []byte("abc"), p.Get())
				wg.Done()
			}(i)
		}

		p.Put([]byte("abc"))
		assert.Equal(t, []byte("abc"), p.Get())
	})
}
