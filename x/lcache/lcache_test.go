package lcache

import (
	"testing"
	"time"

	"github.com/gookit/goutil/testutil/assert"
)

func TestCache_SetAndGet(t *testing.T) {
	c := New()

	t.Run("basic", func(t *testing.T) {
		// Test basic set and get
		c.Set("key1", "value1", 5*time.Minute)
		val, found := c.Get("key1")
		assert.True(t, found)
		assert.Eq(t, "value1", val)

		// Test non-existent key
		_, found = c.Get("non-existent")
		assert.False(t, found)
		c.Clear()

		assert.Empty(t, c.Keys())
		assert.Len(t, c.Len(), 0)
	})

	t.Run("str", func(t *testing.T) {
		c.Set("str", "hello", 5*time.Minute)
		val, found := c.Get("str")
		assert.True(t, found)
		assert.Eq(t, "hello", val)

		// Test mismatch
		c.Set("int", 123, 5*time.Minute)
		_, found = c.Get("int")
		assert.False(t, found)
		c.Clear()
	})

	t.Run("int", func(t *testing.T) {
		c.Set("num", 42, 5*time.Minute)
		val, found := c.Get("num")
		assert.True(t, found)
		assert.Eq(t, 42, val)

		// Test mismatch
		c.Set("str", "hello", 5*time.Minute)
		_, found = c.Get("str")
		assert.False(t, found)
		c.Clear()
	})

	t.Run("bool", func(t *testing.T) {
		c.Set("flag", true, 5*time.Minute)
		val, found := c.Get("flag")
		assert.True(t, found)
		assert.True(t, val.(bool))

		// Test type mismatch
		c.Set("str", "hello", 5*time.Minute)
		_, found = c.Get("str")
		assert.False(t, found)
		c.Clear()
	})
}

func TestCache_Expiration(t *testing.T) {
	c := New()
	defer c.Clear()

	// Set with short TTL
	c.Set("short", "Val", 100*time.Millisecond)
	val, found := c.Get("short")
	assert.True(t, found)
	assert.Eq(t, "Val", val)

	// Wait for expiration
	time.Sleep(150 * time.Millisecond)
	_, found = c.Get("short")
	assert.False(t, found)
}

func TestCache_NoExpiration(t *testing.T) {
	c := New()
	defer c.Clear()

	// Set with duration <= 0 (never expires)
	c.Set("permanent", "Val", 0)
	val, found := c.Get("permanent")
	assert.True(t, found)
	assert.Eq(t, "Val", val)

	// Wait a bit and check again
	time.Sleep(100 * time.Millisecond)
	val, found = c.Get("permanent")
	assert.True(t, found)
	assert.Eq(t, "Val", val)
}

func TestCache_Delete(t *testing.T) {
	c := New()
	defer c.Clear()

	c.Set("key", "Val", 5*time.Minute)
	_, found := c.Get("key")
	assert.True(t, found)

	c.Delete("key")
	_, found = c.Get("key")
	assert.False(t, found)

	// Delete non-existent key should not panic
	c.Delete("non-existent")
}

func TestCache_Has(t *testing.T) {
	c := New()
	defer c.Clear()

	c.Set("key", "Val", 5*time.Minute)
	assert.True(t, c.Has("key"))
	assert.False(t, c.Has("non-existent"))
}

func TestCache_Keys(t *testing.T) {
	c := New()
	defer c.Clear()

	c.Set("key1", "value1", 5*time.Minute)
	c.Set("key2", "value2", 5*time.Minute)
	c.Set("key3", "value3", 5*time.Minute)

	keys := c.Keys()
	assert.Eq(t, 3, len(keys))
	assert.Contains(t, keys, "key1")
	assert.Contains(t, keys, "key2")
	assert.Contains(t, keys, "key3")
}

func TestCache_Len(t *testing.T) {
	c := New()
	defer c.Clear()

	assert.Eq(t, 0, c.Len())

	c.Set("key1", "value1", 5*time.Minute)
	assert.Eq(t, 1, c.Len())

	c.Set("key2", "value2", 5*time.Minute)
	assert.Eq(t, 2, c.Len())
}

func TestCache_Clear(t *testing.T) {
	c := New()
	defer c.Clear()

	c.Set("key1", "value1", 5*time.Minute)
	c.Set("key2", "value2", 5*time.Minute)
	assert.Eq(t, 2, c.Len())

	c.Clear()
	assert.Eq(t, 0, c.Len())
	assert.False(t, c.Has("key1"))
	assert.False(t, c.Has("key2"))
}

func TestCache_Concurrent(t *testing.T) {
	c := New()
	defer c.Clear()

	// Concurrent writes
	for i := 0; i < 100; i++ {
		go func(n int) {
			c.Set("key", n, 5*time.Minute)
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 100; i++ {
		go func() {
			c.Get("key")
		}()
	}

	// Wait a bit for goroutines to complete
	time.Sleep(100 * time.Millisecond)

	// Cache should still be in a valid state
	_, found := c.Get("key")
	assert.True(t, found)
}
