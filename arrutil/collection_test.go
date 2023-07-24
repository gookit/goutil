package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

// StringEqualComparer tests
func TestStringEqualsComparer(t *testing.T) {
	assert.Eq(t, 0, arrutil.StringEqualsComparer("a", "a"))
	assert.Eq(t, -1, arrutil.StringEqualsComparer("a", "b"))
}

func TestValueEqualsComparer(t *testing.T) {
	assert.Eq(t, 0, arrutil.ValueEqualsComparer("1", "1"))
	assert.Eq(t, -1, arrutil.ValueEqualsComparer(1, 2))
}

// ReflectEqualsComparer tests
func TestReflectEqualsComparer(t *testing.T) {
	assert.Eq(t, 0, arrutil.ReflectEqualsComparer(1, 1))
	assert.Eq(t, -1, arrutil.ReflectEqualsComparer(1, 2))
}

// ElemTypeEqualCompareFunc
func TestElemTypeEqualCompareFuncShouldEquals(t *testing.T) {
	var c = 1
	assert.Eq(t, 0, arrutil.ElemTypeEqualsComparer(c, c))
	assert.Eq(t, 0, arrutil.ElemTypeEqualsComparer(1, 1))

	var a, b any
	a = 1
	b = "2"
	assert.Eq(t, -1, arrutil.ElemTypeEqualsComparer(a, b))
}

func TestDifferencesShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}
	result := arrutil.Differences[string](data, []string{"a", "b"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"c"}, result)

	result = arrutil.Differences[string]([]string{"a", "b"}, data, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"c"}, result)

	result = arrutil.Diff([]string{"a", "b", "d"}, data, arrutil.ReflectEqualsComparer[string])
	assert.Eq(t, 2, len(result))
}

func TestExceptsShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}
	result := arrutil.Excepts(data, []string{"a", "b"}, arrutil.ValueEqualsComparer[string])
	assert.Eq(t, []string{"c"}, result)
}

func TestExceptsFirstEmptyShouldReturnsEmpty(t *testing.T) {
	data := []string{}
	result := arrutil.Excepts(data, []string{"a", "b"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

func TestExceptsSecondEmptyShouldReturnsFirst(t *testing.T) {
	data := []string{"a", "b"}
	result := arrutil.Excepts(data, []string{}, arrutil.StringEqualsComparer)
	assert.Eq(t, data, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Intersects tests
func TestIntersectsShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}
	result := arrutil.Intersects(data, []string{"a", "b"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"a", "b"}, result)
}

func TestIntersectsFirstEmptyShouldReturnsEmpty(t *testing.T) {
	var data []string
	second := []string{"a", "b"}
	result := arrutil.Intersects(data, second, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &second, &result, "should always returns new slice")
}

func TestIntersectsSecondEmptyShouldReturnsEmpty(t *testing.T) {
	data := []string{"a", "b"}
	second := []string{}
	result := arrutil.Intersects(data, second, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Union tests

func TestUnionShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}
	result := arrutil.Union(data, []string{"a", "b", "d"}, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"a", "b", "c", "d"}, result)
}

func TestUnionFirstEmptyShouldReturnsSecond(t *testing.T) {
	data := []string{}
	second := []string{"a", "b"}
	result := arrutil.Union(data, second, arrutil.StringEqualsComparer)
	assert.Eq(t, []string{"a", "b"}, result)
	assert.NotSame(t, &second, &result, "should always returns new slice")
}

func TestUnionSecondEmptyShouldReturnsFirst(t *testing.T) {
	data := []string{"a", "b"}
	second := []string{}
	result := arrutil.Union(data, second, arrutil.StringEqualsComparer)
	assert.Eq(t, data, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// Find tests
func TestFindShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}

	result, err := arrutil.Find(data, func(a string) bool { return a == "b" })
	assert.Nil(t, err)
	assert.Eq(t, "b", result)

	_, err = arrutil.Find(data, func(a string) bool { return a == "d" })
	assert.NotNil(t, err)
	assert.Eq(t, arrutil.ErrElementNotFound, err)

}

func TestFindEmptyReturnsErrElementNotFound(t *testing.T) {
	data := []string{}
	_, err := arrutil.Find(data, func(a string) bool { return a == "b" })
	assert.NotNil(t, err)
	assert.Eq(t, arrutil.ErrElementNotFound, err)
}

// FindOrDefault tests
func TestFindOrDefaultShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}

	result := arrutil.FindOrDefault(data, func(a string) bool { return a == "b" }, "d")
	assert.Eq(t, "b", result)

	result = arrutil.FindOrDefault(data, func(a string) bool { return a == "d" }, "d")
	assert.Eq(t, "d", result)
}

// TakeWhile tests
func TestTakeWhileShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}

	result := arrutil.TakeWhile(data, func(a string) bool { return a == "b" || a == "c" })
	assert.Eq(t, []string{"b", "c"}, result)
}

func TestTakeWhileEmptyReturnsEmpty(t *testing.T) {
	var data []string
	result := arrutil.TakeWhile(data, func(a string) bool { return a == "b" || a == "c" })
	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

// ExceptWhile tests

func TestExceptWhileShouldPassed(t *testing.T) {
	data := []string{"a", "b", "c"}

	result := arrutil.ExceptWhile(data, func(a string) bool { return a == "b" || a == "c" })
	assert.Eq(t, []string{"a"}, result)
}

func TestExceptWhileEmptyReturnsEmpty(t *testing.T) {
	var data []string
	result := arrutil.ExceptWhile(data, func(a string) bool { return a == "b" || a == "c" })

	assert.Eq(t, []string{}, result)
	assert.NotSame(t, &data, &result, "should always returns new slice")
}

func TestMap(t *testing.T) {
	list1 := []map[string]any{
		{"name": "tom", "age": 23},
		{"name": "john", "age": 34},
	}

	flatArr := arrutil.Column(list1, func(obj map[string]any) (val any, find bool) {
		return obj["age"], true
	})

	assert.NotEmpty(t, flatArr)
	assert.Contains(t, flatArr, 23)
	assert.Len(t, flatArr, 2)
	assert.Eq(t, 34, flatArr[1])
}
