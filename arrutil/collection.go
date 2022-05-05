package arrutil

import (
	"errors"
	"reflect"
)

var (
	// 对比字符串
	StringEqualCompareFunc = func(a, b interface{}) int {
		if reflect.TypeOf(a) != reflect.TypeOf(b) {
			return -1
		}
		strA := ""
		strB := ""

		at := reflect.TypeOf(a)
		if at.Kind() == reflect.String {
			strA = a.(string)
		} else if at.Kind() == reflect.Ptr && at.Elem().Kind() == reflect.String {
			strA = *a.(*string)
		} else {
			return -1
		}

		bt := reflect.TypeOf(b)
		if bt.Kind() == reflect.String {
			strB = b.(string)
		} else if bt.Kind() == reflect.Ptr && bt.Elem().Kind() == reflect.String {
			strB = *b.(*string)
		} else {
			return -1
		}

		if strA == strB {
			return 0
		}
		return 1
	}

	// 对比指针地址，如果相同则返回0，否则返回1
	EqualCompareFunc = func(a, b interface{}) int {
		if a == b {
			return 0
		}
		return 1
	}

	// 对比对象的类型，如果相同则返回0，否则返回1
	ElemTypeEqualCompareFunc = func(a, b interface{}) int {
		at := reflect.TypeOf(a)
		bt := reflect.TypeOf(b)
		if at.Kind() == reflect.Ptr {
			at = at.Elem()
		}

		if bt.Kind() == reflect.Ptr {
			bt = bt.Elem()
		}

		if at == bt {
			return 0
		}
		return 1

	}
)

// 检查数组中是否包含指定的元素
// data: 数组
// item: 指定的元素
// comparer: 比较器
// return: 包含返回true, 否则返回false
func ContainsItem(data interface{}, item interface{}, comparer func(a interface{}, b interface{}) int) bool {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Slice && dataType.Kind() != reflect.Array {
		panic("data must be a slice or array")
	}
	s := reflect.ValueOf(data)
	for i := 0; i < s.Len(); i++ {
		if comparer(s.Index(i).Interface(), item) == 0 {
			return true
		}
	}
	return false
}

// 在两个数组之间取差集
// a: 数组a
// b: 数组b
// comparer: 比较器
// return: 差集
func Excepts(source interface{}, data interface{}, comparer func(a interface{}, b interface{}) int) interface{} {
	aType := reflect.TypeOf(source)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.Excepts: source must be a slice or array")
	}
	var itemType reflect.Type
	sourceVal := reflect.ValueOf(source)
	if sourceVal.Len() == 0 {
		return data
	} else {
		itemType = sourceVal.Index(0).Type()
	}

	bType := reflect.TypeOf(data)
	if bType.Kind() != reflect.Slice && bType.Kind() != reflect.Array {
		panic("collections.Excepts: data must be a slice or array")
	}

	bVal := reflect.ValueOf(data)
	if bVal.Len() == 0 {
		return source
	}

	result := reflect.New(reflect.ArrayOf(0, itemType)).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		for j := 0; j < bVal.Len(); j++ {
			s := sourceVal.Index(i).Interface()
			d := bVal.Index(j).Interface()
			if comparer(s, d) != 0 {
				result = reflect.Append(result, reflect.ValueOf(s))
			}
		}
	}

	return result
}

// 在两个数组之间取交集
// a: 数组a
// b: 数组b
// comparer: 比较器
// return: 交集
func Intersects(source interface{}, data interface{}, comparer func(a interface{}, b interface{}) int) interface{} {
	aType := reflect.TypeOf(source)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.Excepts: source must be a slice or array")
	}
	var itemType reflect.Type
	sourceVal := reflect.ValueOf(source)
	if sourceVal.Len() == 0 {
		return data
	} else {
		itemType = sourceVal.Index(0).Type()
	}

	bType := reflect.TypeOf(data)
	if bType.Kind() != reflect.Slice && bType.Kind() != reflect.Array {
		panic("collections.Excepts: data must be a slice or array")
	}

	bVal := reflect.ValueOf(data)
	if bVal.Len() == 0 {
		return source
	}

	result := reflect.New(reflect.ArrayOf(0, itemType)).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		for j := 0; j < bVal.Len(); j++ {
			s := sourceVal.Index(i).Interface()
			d := bVal.Index(j).Interface()
			if comparer(s, d) == 0 {
				result = reflect.Append(result, reflect.ValueOf(s))
			}
		}
	}

	return result
}

// 在两个数组之间取并集
// a: 数组a
// b: 数组b
// comparer: 比较器
func Union(source interface{}, data interface{}, comparer func(a interface{}, b interface{}) int) interface{} {
	aType := reflect.TypeOf(source)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.Excepts: source must be a slice or array")
	}
	var itemType reflect.Type
	sourceVal := reflect.ValueOf(source)
	if sourceVal.Len() == 0 {
		return data
	} else {
		itemType = sourceVal.Index(0).Type()
	}

	bType := reflect.TypeOf(data)
	if bType.Kind() != reflect.Slice && bType.Kind() != reflect.Array {
		panic("collections.Excepts: data must be a slice or array")
	}

	bVal := reflect.ValueOf(data)
	if bVal.Len() == 0 {
		return source
	}

	result := reflect.New(reflect.ArrayOf(0, itemType)).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		for j := 0; j < result.Len(); j++ {
			s := sourceVal.Index(i).Interface()
			d := result.Index(j).Interface()
			if comparer(s, d) != 0 {
				result = reflect.Append(result, reflect.ValueOf(s))
			}
		}
	}
	for i := 0; i < bVal.Len(); i++ {
		for j := 0; j < result.Len(); j++ {
			s := bVal.Index(i).Interface()
			d := result.Index(j).Interface()
			if comparer(s, d) != 0 {
				result = reflect.Append(result, reflect.ValueOf(s))
			}
		}
	}
	return result
}

var ErrElementNotFound = errors.New("element not found")

// 在数组中查找符合条件的元素, 如果没有找到, 则返回 ErrElementNotFound
// source: 数组
// predicate: 条件
// return: 符合条件的元素
func Find(source interface{}, predicate func(a interface{}) bool) (interface{}, error) {
	aType := reflect.TypeOf(source)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.Find: source must be a slice or array")
	}

	sourceVal := reflect.ValueOf(source)
	if sourceVal.Len() == 0 {
		return nil, ErrElementNotFound
	}

	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if predicate(s) {
			return s, nil
		}
	}
	return nil, ErrElementNotFound
}

// 在数组中查找符合条件的元素, 如果没有找到, 则返回默认值
// data: 数组
// predicate: 条件
// defaultValue: 默认值
// return: 符合条件的元素
func FindOrDefault(source interface{}, predicate func(a interface{}) bool, defaultValue interface{}) interface{} {
	item, err := Find(source, predicate)
	if err != nil {
		if err == ErrElementNotFound {
			return defaultValue
		}
	}
	return item
}

// 在数组中查找符合条件的元素, 如果没有找到, 则返回空数组
// data: 数组
// predicate: 条件
// return: 符合条件的元素
func TakeWhile(data interface{}, predicate func(a interface{}) bool) interface{} {
	aType := reflect.TypeOf(data)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.TakeWhile: data must be a slice or array")
	}
	sourceVal := reflect.ValueOf(data)
	if sourceVal.Len() == 0 {
		return data
	}
	itemType := sourceVal.Index(0).Type()
	result := reflect.New(reflect.ArrayOf(0, itemType)).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if predicate(s) {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}
	return result
}

// 在数组中排除符合条件的元素, 如果没有找到, 则返回空数组
// data: 数组
// predicate: 条件
// return: 符合条件的元素
func ExceptWhile(data interface{}, predicate func(a interface{}) bool) interface{} {
	aType := reflect.TypeOf(data)
	if aType.Kind() != reflect.Slice && aType.Kind() != reflect.Array {
		panic("collections.ExceptWhile: data must be a slice or array")
	}
	sourceVal := reflect.ValueOf(data)
	if sourceVal.Len() == 0 {
		return data
	}
	itemType := sourceVal.Index(0).Type()
	result := reflect.New(reflect.ArrayOf(0, itemType)).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if !predicate(s) {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}
	return result
}
