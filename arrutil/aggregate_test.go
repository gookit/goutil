package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/stretchr/testify/assert"
)

func TestAggregate_FnReturnsValue_ShouldPass(t *testing.T) {
	data := []int{
		10, 20, 30,
	}
	results := arrutil.Aggregate(data, func(acc, val interface{}) interface{} {
		ret := 0
		for _, v := range acc.([]int) {
			ret += v
		}
		ret += val.(int)
		return ret
	})
	// seed is empty, sum(seed) = 0, result is 10
	assert.Equal(t, 10, results.([]int)[0])
	// seed is 10, sum(seed) = 10, result is 30
	assert.Equal(t, 30, results.([]int)[1])
	// seed is 10, 30, sum(seed) = 40, result is 70
	assert.Equal(t, 70, results.([]int)[2])
}

func TestAggregate_FnReturnsSlice_ShouldPass(t *testing.T) {
	data := []int{
		10,
	}
	results := arrutil.Aggregate(data, func(acc, val interface{}) interface{} {
		ret := 0
		for _, v := range acc.([]int) {
			ret += v
		}
		ret += val.(int)
		retSlice := append(acc.([]int), ret)
		return retSlice
	})
	// seed is empty, sum(seed) = 0, result is 0
	assert.Equal(t, 10, results.([]int)[0])
}
