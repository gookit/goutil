package arrutil

import (
	"reflect"

	"github.com/gookit/goutil/comdef"
)

// Reverse any T slice.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"0", "info", "user", "site"}
func Reverse[T any](ls []T) {
	ln := len(ls)
	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		ls[i], ls[li] = ls[li], ls[i]
	}
}

// Remove give element from slice []T.
//
// eg: []string{"site", "user", "info", "0"} -> []string{"site", "user", "info"}
func Remove[T comdef.Compared](ls []T, val T) []T {
	return Filter(ls, func(el T) bool {
		return el != val
	})
}

// Filter given slice, default will filter zero value.
//
// Usage:
//
//	// output: [a, b]
//	ss := arrutil.Filter([]string{"a", "", "b", ""})
func Filter[T any](ls []T, filter ...comdef.MatchFunc[T]) []T {
	var fn comdef.MatchFunc[T]
	if len(filter) > 0 && filter[0] != nil {
		fn = filter[0]
	} else {
		fn = func(el T) bool {
			// if el == nil { // Filter nil value
			// 	return false
			// }
			return !reflect.ValueOf(el).IsZero()
		}
	}

	newLs := make([]T, 0, len(ls))
	for _, el := range ls {
		if fn(el) {
			newLs = append(newLs, el)
		}
	}
	return newLs
}

// Map a list to new list with map filter function.
//
// eg: mapping [object0{},object1{},...] to flatten list [object0.someKey, object1.someKey, ...]
func Map[T, V any](list []T, mapFilter func(input T) (target V, ok bool)) []V {
	if len(list) == 0 {
		return nil
	}

	flatArr := make([]V, 0, len(list))
	for _, obj := range list {
		if target, ok := mapFilter(obj); ok {
			flatArr = append(flatArr, target)
		}
	}
	return flatArr
}

// Map1 a list to new list with map function.
//
// eg: mapping [object0{},object1{},...] to flatten list [object0.someKey, object1.someKey, ...]
func Map1[T, R any](list []T, mapFn func(t T) R) []R {
	if len(list) == 0 {
		return nil
	}

	ret := make([]R, len(list))
	for i := range list {
		ret[i] = mapFn(list[i])
	}
	return ret
}

// Column collect sub elements from list. alias of Map func
//
// Example:
//   list := []map[string]any{
//     {"id": 1, "name": "one", "age": 23},
//     {"id": 2, "name": "two", "age": 23},
//     {"id": 3, "name": "three", "age": 23},
//   }
//   names := arrutil.Column(list, func(el map[string]any) string {
//     return el["name"].(string)
//   })
func Column[T any, V any](list []T, mapFn func(obj T) (val V, find bool)) []V {
	return Map(list, mapFn)
}

// Unique value in the given slice data.
func Unique[T comdef.NumberOrString](list []T) []T {
	if len(list) < 2 {
		return list
	}

	valMap := make(map[T]struct{}, len(list))
	uniArr := make([]T, 0, len(list))

	for _, t := range list {
		if _, ok := valMap[t]; !ok {
			valMap[t] = struct{}{}
			uniArr = append(uniArr, t)
		}
	}
	return uniArr
}

// IndexOf value in given slice.
func IndexOf[T comdef.NumberOrString](val T, list []T) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

// FirstOr get first value of slice, if slice is empty, return the default value.
func FirstOr[T any](list []T, defVal ...T) T {
	if len(list) > 0 {
		return list[0]
	}

	if len(defVal) > 0 {
		return defVal[0]
	}
	var zero T
	return zero
}

// Chunk split slice to chunks by size.
//
// eg: [1,2,3,4,5,6,7,8,9,10] -> [[1,2,3,4], [5,6,7,8], [9,10]]
func Chunk[T any](list []T, size int) [][]T {
	if size <= 0 {
		return nil
	}

	ln := len(list)
	if ln == 0 {
		return nil
	}

	chunks := make([][]T, 0, ln/size+1)

	for i := 0; i < ln; i += size {
		end := i + size
		if end > ln {
			end = ln
		}
		chunks = append(chunks, list[i:end])
	}

	return chunks
}

// ChunkBy split slice to chunks by size, and with custom chunk function.
//
// Example:
//   list := []map[string]any{
//     {"id": 1, "name": "one", "age": 23},
//     {"id": 2, "name": "two", "age": 23},
//     {"id": 3, "name": "three", "age": 23},
//   }
//   chunks := arrutil.ChunkBy(list, 2, func(el map[string]any) map[string]any {
//     return map[string]any{
//       "id": el["id"],
//       "name": el["name"],
//     }
//   })
// 	Output: [
// 		[{"id": 1, "name": "one"}, {"id": 2, "name": "two"}],
// 		[{"id": 3, "name": "three"}]
// 	]
func ChunkBy[T, R any](list []T, size int, mapFn func(el T) R) [][]R {
	if size <= 0 {
		return nil
	}

	ln := len(list)
	if ln == 0 {
		return nil
	}

	// 计算需要的块数量
	numChunks := ln/size + 1
	if ln%size == 0 {
		numChunks = ln / size
	}

	chunks := make([][]R, 0, numChunks)

	for i := 0; i < ln; i += size {
		end := i + size
		if end > ln {
			end = ln
		}

		// 创建当前块的切片
		currentChunk := make([]R, 0, size)
		for j := i; j < end; j++ {
			currentChunk = append(currentChunk, mapFn(list[j]))
		}
		chunks = append(chunks, currentChunk)
	}

	return chunks
}
