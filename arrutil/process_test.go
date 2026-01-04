package arrutil_test

import (
	"testing"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestReverse(t *testing.T) {
	ss := []string{"a", "b", "c"}
	arrutil.Reverse(ss)
	assert.Eq(t, []string{"c", "b", "a"}, ss)

	ints := []int{1, 2, 3}
	arrutil.Reverse(ints)
	assert.Eq(t, []int{3, 2, 1}, ints)
}

func TestRemove(t *testing.T) {
	ss := []string{"a", "b", "c"}
	ns := arrutil.Remove(ss, "b")
	assert.Eq(t, []string{"a", "c"}, ns)

	ints := []int{1, 2, 3}
	ni := arrutil.Remove(ints, 2)
	assert.Eq(t, []int{1, 3}, ni)
}

func TestFilter(t *testing.T) {
	is := assert.New(t)
	ss := arrutil.Filter([]string{"a", "", "b", ""})
	is.Eq([]string{"a", "b"}, ss)
}

func TestFirstOr(t *testing.T) {
	is := assert.New(t)
	is.Eq("a", arrutil.FirstOr([]string{"a", "b"}, "c"))
	is.Eq("c", arrutil.FirstOr([]string{}, "c"))
	is.Eq("c", arrutil.FirstOr(nil, "c"))
	is.Eq("", arrutil.FirstOr([]string{}))
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

	t.Run("empty", func(t *testing.T) {
		ss := arrutil.Map([]string{}, func(s string) (string, bool) {
			return s, true
		})
		assert.Empty(t, ss)
	})

	t.Run("Map1", func(t *testing.T) {
		names := arrutil.Map1(list1, func(obj map[string]any) string {
			return obj["name"].(string)
		})
		assert.NotEmpty(t, names)
		assert.Contains(t, names, "tom")
		assert.Len(t, names, 2)
		assert.Eq(t, "john", names[1])

		ss := arrutil.Map1([]string{}, func(s string) string {
			return s
		})
		assert.Empty(t, ss)
	})
}

func TestChunk(t *testing.T) {
	// 测试 size <= 0 的情况
	t.Run("SizeLessThanOrEqualZero", func(t *testing.T) {
		result := arrutil.Chunk([]int{1, 2, 3}, 0)
		assert.Equal(t, [][]int(nil), result)

		result = arrutil.Chunk([]int{1, 2, 3}, -1)
		assert.Equal(t, [][]int(nil), result)
	})

	// 测试空切片的情况
	t.Run("EmptySlice", func(t *testing.T) {
		result := arrutil.Chunk([]int{}, 3)
		assert.Equal(t, [][]int(nil), result)
	})

	// 测试 size = 1 的情况
	t.Run("SizeOne", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expected := [][]int{{1}, {2}, {3}, {4}}
		result := arrutil.Chunk(input, 1)
		assert.Equal(t, expected, result)
	})

	// 测试 size 等于切片长度的情况
	t.Run("SizeEqualsSliceLength", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{{1, 2, 3}}
		result := arrutil.Chunk(input, 3)
		assert.Equal(t, expected, result)
	})

	// 测试 size 大于切片长度的情况
	t.Run("SizeGreaterThanSliceLength", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{{1, 2, 3}}
		result := arrutil.Chunk(input, 5)
		assert.Equal(t, expected, result)
	})

	// 测试能整除的情况
	t.Run("SizeDividesEvenly", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6}
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		result := arrutil.Chunk(input, 2)
		assert.Equal(t, expected, result)
	})

	// 测试不能整除的情况
	t.Run("SizeDoesNotDivideEvenly", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		result := arrutil.Chunk(input, 2)
		assert.Equal(t, expected, result)
	})

	// 测试字符串切片
	t.Run("StringSlice", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e"}
		expected := [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
		result := arrutil.Chunk(input, 2)
		assert.Equal(t, expected, result)
	})

	// 测试自定义结构体切片
	t.Run("StructSlice", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		input := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
		}
		expected := [][]Person{
			{{"Alice", 25}, {"Bob", 30}},
			{{"Charlie", 35}},
		}
		result := arrutil.Chunk(input, 2)
		assert.Equal(t, expected, result)
	})
}

func TestChunkBy(t *testing.T) {
	t.Run("Size lte 0", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2, 3}, 0, func(el int) int { return el })
		assert.Equal(t, [][]int(nil), result)

		result = arrutil.ChunkBy([]int{1, 2, 3}, -1, func(el int) int { return el })
		assert.Equal(t, [][]int(nil), result)
	})

	t.Run("EmptyList", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{}, 2, func(el int) int { return el })
		assert.Equal(t, [][]int(nil), result)
	})

	t.Run("SizeLargerThanList", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2}, 5, func(el int) int { return el })
		expected := [][]int{{1, 2}}
		assert.Equal(t, expected, result)
	})

	t.Run("SizeEqualsList", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2, 3, 4}, 4, func(el int) int { return el })
		expected := [][]int{{1, 2, 3, 4}}
		assert.Equal(t, expected, result)
	})

	t.Run("ExactDivision", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2, 3, 4, 5, 6}, 2, func(el int) int { return el })
		expected := [][]int{{1, 2}, {3, 4}, {5, 6}}
		assert.Equal(t, expected, result)
	})

	t.Run("NotExactDivision", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2, 3, 4, 5}, 2, func(el int) int { return el })
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		assert.Equal(t, expected, result)
	})

	t.Run("StringToIntMapping", func(t *testing.T) {
		result := arrutil.ChunkBy([]string{"1", "2", "3", "4", "5"}, 2, func(el string) int {
			// 简单的字符串转整数示例（实际项目中应使用strconv.Atoi）
			switch el {
			case "1":
				return 1
			case "2":
				return 2
			case "3":
				return 3
			case "4":
				return 4
			case "5":
				return 5
			default:
				return 0
			}
		})
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		assert.Equal(t, expected, result)
	})

	t.Run("StructMapping", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		people := []Person{
			{"Alice", 25},
			{"Bob", 30},
			{"Charlie", 35},
			{"David", 40},
		}

		result := arrutil.ChunkBy(people, 2, func(p Person) string {
			return p.Name
		})
		expected := [][]string{{"Alice", "Bob"}, {"Charlie", "David"}}
		assert.Equal(t, expected, result)
	})

	t.Run("SingleElement", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{42}, 1, func(el int) int { return el * 2 })
		expected := [][]int{{84}}
		assert.Equal(t, expected, result)
	})

	t.Run("LargeSize", func(t *testing.T) {
		result := arrutil.ChunkBy([]int{1, 2, 3}, 100, func(el int) int { return el })
		expected := [][]int{{1, 2, 3}}
		assert.Equal(t, expected, result)
	})
}