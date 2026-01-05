package syncs_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/gookit/goutil/syncs"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.NewEchoServer()
	defer s.Close()
	testSrvAddr = "http://" + s.Listener.Addr().String()
	fmt.Println("Test server listen on:", testSrvAddr)

	m.Run()
}

func TestWaitGroupGo(t *testing.T) {
	var wg syncs.WaitGroup
	var counter int32
	expected := int32(5)

	// 启动多个 goroutines
	for i := int32(0); i < expected; i++ {
		wg.Go(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	// 等待所有 goroutines 完成
	wg.Wait()
	assert.Eq(t, expected, counter)
}

func TestContextValue(t *testing.T) {
	assert.NotEmpty(t, syncs.ContextValue("key", "value"))
}

func TestSafeMap(t *testing.T) {
	t.Run("syncs.NewSafeMap", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		assert.NotNil(t, m)
		assert.Equal(t, 0, m.Len())
	})

	t.Run("SetAndGet", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		m.Set("key1", 100)

		value, ok := m.Get("key1")
		assert.True(t, ok)
		assert.Equal(t, 100, value)
	})

	t.Run("GetNonExistentKey", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		value, ok := m.Get("nonexistent")
		assert.False(t, ok)
		assert.Equal(t, 0, value)
	})

	t.Run("Delete", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		m.Set("key1", 100)

		// Verify key exists
		_, ok := m.Get("key1")
		assert.True(t, ok)

		// Delete the key
		m.Delete("key1")

		// Verify key is gone
		_, ok = m.Get("key1")
		assert.False(t, ok)
	})

	t.Run("ConcurrentSetAndGet", func(t *testing.T) {
		m := syncs.NewSafeMap[int, string]()
		const numGoroutines = 100

		// Concurrently set values
		for i := 0; i < numGoroutines; i++ {
			go m.Set(i, "value"+string(rune(i+'0')))
		}

		// Give some time for goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			go func(i int) {
				_, ok := m.Get(i)
				if !ok {
					t.Errorf("Expected key %d to exist", i)
				}
			}(i)
		}
	})

	t.Run("Range", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		m.Set("a", 1)
		m.Set("b", 2)
		m.Set("c", 3)

		count := 0
		m.Range(func(key string, value int) {
			count++
			assert.Contains(t, []string{"a", "b", "c"}, key)
			assert.Contains(t, []int{1, 2, 3}, value)
		})

		assert.Equal(t, 3, count)
	})

	t.Run("Clear", func(t *testing.T) {
		m := syncs.NewSafeMap[string, int]()
		m.Set("key1", 100)
		m.Set("key2", 200)

		// Verify map has items
		_, ok := m.Get("key1")
		assert.True(t, ok)

		// Clear the map
		m.Clear()

		// Verify map is empty
		_, ok = m.Get("key1")
		assert.False(t, ok)
		_, ok = m.Get("key2")
		assert.False(t, ok)
		assert.Equal(t, 0, m.Len())
	})

	t.Run("MultipleTypes", func(t *testing.T) {
		// Test with different key and value types
		m1 := syncs.NewSafeMap[int, string]()
		m1.Set(1, "one")
		value1, ok1 := m1.Get(1)
		assert.True(t, ok1)
		assert.Equal(t, "one", value1)

		m2 := syncs.NewSafeMap[string, bool]()
		m2.Set("enabled", true)
		value2, ok2 := m2.Get("enabled")
		assert.True(t, ok2)
		assert.Equal(t, true, value2)
	})
}