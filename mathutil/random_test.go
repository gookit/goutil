package mathutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	min, max := 1000, 9999

	for i := 0; i < 5; i++ {
		val := RandomInt(min, max)
		fmt.Println(val)
		assert.True(t, val >= min)
		assert.True(t, val <= max)

		seed := time.Now().UnixNano()
		val = RandomIntWithSeed(min, max, seed)
		assert.True(t, val >= min)
	}
}
