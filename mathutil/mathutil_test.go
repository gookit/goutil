package mathutil_test

import (
	"testing"

	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestDiv(t *testing.T) {
	assert.Eq(t, float64(20), mathutil.Div(27, 1.35))
	assert.Eq(t, 14, mathutil.DivInt(27, 2))
	assert.Eq(t, 20, mathutil.DivF2i(27, 1.35))
}

func TestMul(t *testing.T) {
	assert.Eq(t, float64(5), mathutil.Mul(2, 2.35))
	assert.Eq(t, 36, mathutil.MulF2i(27, 1.35))
}

// TestRange tests the Range function with different input cases
func TestRange(t *testing.T) {
	t.Run("single number", func(t *testing.T) {
		var results []int
		err := mathutil.Range("5", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{5}, results)
	})

	t.Run("range expression", func(t *testing.T) {
		var results []int
		err := mathutil.Range("1-3", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, results)
	})

	t.Run("multiple numbers and ranges", func(t *testing.T) {
		var results []int
		err := mathutil.Range("1,3-5,7", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 3, 4, 5, 7}, results)
	})

	t.Run("with empty items", func(t *testing.T) {
		var results []int
		err := mathutil.Range("1, ,3", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 3}, results)
	})

	t.Run("invalid single number", func(t *testing.T) {
		err := mathutil.Range("abc", func(val int) {
			// This should not be called
		})
		assert.ErrMsg(t, err, "invalid integer value: abc")
	})

	t.Run("invalid range expression", func(t *testing.T) {
		err := mathutil.Range("a-b", func(val int) {
			// This should not be called
		})
		assert.Error(t, err)
	})

	t.Run("negative numbers range", func(t *testing.T) {
		var results []int
		err := mathutil.Range("-2-1", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{-2, -1, 0, 1}, results)
	})

	t.Run("single number with negative", func(t *testing.T) {
		var results []int
		err := mathutil.Range("-5", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{-5}, results)
	})

	t.Run("empty string input", func(t *testing.T) {
		var results []int
		err := mathutil.Range("", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{}, results)
	})

	t.Run("only empty items", func(t *testing.T) {
		var results []int
		err := mathutil.Range(", ,", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{}, results)
	})

	t.Run("zero range", func(t *testing.T) {
		var results []int
		err := mathutil.Range("0-0", func(val int) {
			results = append(results, val)
		})
		assert.NoError(t, err)
		assert.Equal(t, []int{0}, results)
	})

	t.Run("large range", func(t *testing.T) {
		count := 0
		err := mathutil.Range("1-3", func(val int) {
			count++
		})
		assert.NoError(t, err)
		assert.Equal(t, 3, count)
	})
}

func TestExpand(t *testing.T) {
	// Test normal cases
	t.Run("simple number list", func(t *testing.T) {
		result, err := mathutil.Expand("1,2,3")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("number range", func(t *testing.T) {
		result, err := mathutil.Expand("1-3")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	t.Run("mixed numbers and ranges", func(t *testing.T) {
		result, err := mathutil.Expand("1,2-4,5")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("with empty items", func(t *testing.T) {
		result, err := mathutil.Expand("1,,2,3")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, result)
	})

	// Test edge cases
	var emptyInts []int
	t.Run("empty string", func(t *testing.T) {
		result, err := mathutil.Expand("")
		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("only commas", func(t *testing.T) {
		result, err := mathutil.Expand(",")
		assert.NoError(t, err)
		assert.Equal(t, emptyInts, result)
	})

	t.Run("multiple commas", func(t *testing.T) {
		result, err := mathutil.Expand(",,,")
		assert.NoError(t, err)
		assert.Equal(t, emptyInts, result)
	})

	t.Run("single number", func(t *testing.T) {
		result, err := mathutil.Expand("42")
		assert.NoError(t, err)
		assert.Equal(t, []int{42}, result)
	})

	t.Run("single range", func(t *testing.T) {
		result, err := mathutil.Expand("1-5")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("range with same start and end", func(t *testing.T) {
		result, err := mathutil.Expand("5-5")
		assert.NoError(t, err)
		assert.Equal(t, []int{5}, result)
	})

	// Test error cases
	t.Run("error cases", func(t *testing.T) {
		t.Run("invalid number", func(t *testing.T) {
			result, err := mathutil.Expand("abc")
			assert.Nil(t, result)
			assert.ErrMsg(t, err, "invalid integer value: abc")
		})

		t.Run("invalid number in list", func(t *testing.T) {
			result, err := mathutil.Expand("1,abc,3")
			assert.ErrMsg(t, err, "invalid integer value: abc")
			assert.Nil(t, result)
		})

		t.Run("invalid range format", func(t *testing.T) {
			result, err := mathutil.Expand("1-2-3")
			assert.ErrHasMsg(t, err, "invalid range end value")
			assert.Nil(t, result)
		})

		t.Run("range with invalid numbers", func(t *testing.T) {
			result, err := mathutil.Expand("a-b")
			assert.ErrMsg(t, err, "invalid range start value: a")
			assert.Nil(t, result)
		})

		t.Run("range with non-numeric start", func(t *testing.T) {
			result, err := mathutil.Expand("abc-5")
			assert.ErrSubMsg(t, err, "invalid range start value")
			assert.Nil(t, result)
		})

		t.Run("range with non-numeric end", func(t *testing.T) {
			result, err := mathutil.Expand("1-xyz")
			assert.ErrMsg(t, err, "invalid range end value: xyz")
			assert.Nil(t, result)
		})
	})

	// Test range edge cases
	t.Run("descending range", func(t *testing.T) {
		result, err := mathutil.Expand("5-1")
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
	})

	t.Run("negative numbers", func(t *testing.T) {
		// This should be "-2--1" which means -2 to -1
		result, err := mathutil.Expand("-2--1")
		assert.NoError(t, err)
		assert.Equal(t, []int{-2, -1}, result)
	})

	t.Run("negative to positive", func(t *testing.T) {
		result, err := mathutil.Expand("-1-1")
		assert.NoError(t, err)
		assert.Equal(t, []int{-1, 0, 1}, result)
	})
}
