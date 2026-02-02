package lcache_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gookit/goutil/x/lcache"
)

func ExampleCache() {
	cache := lcache.New()
	defer cache.Clear()

	// Set values
	cache.Set("name", "John", 5*time.Minute)
	cache.Set("age", 30, 0) // Never expires

	// Get values
	name, _ := cache.Get("name")
	age, _ := cache.Get("age")

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Output:
	// Name: John, Age: 30
}

func ExampleCache_withOptions() {
	var buf bytes.Buffer

	// Create cache with custom settings
	cache := lcache.New(
		lcache.WithCapacity(10),
		lcache.WithSerializer("gob"),
		lcache.WithOnEvictFn(func(key string, value any) {
			buf.WriteString(fmt.Sprintf("Evicting %s: %v\n", key, value))
		}),
	)
	defer cache.Clear()

	cache.Set("key", "Val", 1*time.Hour)

	value, found := cache.Get("key")
	if found {
		fmt.Println(value)
	}

	// Output:
	// Val
}
