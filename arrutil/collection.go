package arrutil

import (
	"errors"
	"reflect"
)

const ErrElementNotFound = "element not found"

type Comparer func(a, b interface{}) int
type Predicate func(a interface{}) bool

var (
	StringEqualsComparer Comparer = func(a, b interface{}) int {
		typeOfA := reflect.TypeOf(a)
		if typeOfA.Kind() == reflect.Ptr {
			typeOfA = typeOfA.Elem()
		}

		typeOfB := reflect.TypeOf(b)
		if typeOfB.Kind() == reflect.Ptr {
			typeOfB = typeOfB.Elem()
		}

		if typeOfA != typeOfB {
			return -1
		}

		strA := ""
		strB := ""

		if val, ok := a.(string); ok {
			strA = val
		} else if val, ok := a.(*string); ok {
			strA = *val
		} else {
			return -1
		}

		if val, ok := b.(string); ok {
			strB = val
		} else if val, ok := b.(*string); ok {
			strB = *val
		} else {
			return -1
		}

		if strA == strB {
			return 0
		}
		return -1
	}

	ReferenceEqualsComparer Comparer = func(a, b interface{}) int {
		if a == b {
			return 0
		}
		return -1
	}

	ElemTypeEqualsComparer Comparer = func(a, b interface{}) int {
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
		return -1
	}
)

func TwowaySearch(data interface{}, item interface{}, fn Comparer) (int, error) {
	if data == nil {
		return -1, errors.New("collections.TwowaySearch: data is nil")
	}
	if fn == nil {
		return -1, errors.New("collections.TwowaySearch: fn is nil")
	}

	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Slice {
		return -1, errors.New("collections.TwowaySearch: data is not a slice")
	}

	dataVal := reflect.ValueOf(data)
	if dataVal.Len() == 0 {
		return -1, errors.New("collections.TwowaySearch: data is empty")
	}
	itemType := dataType.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}

	if itemType != dataVal.Index(0).Type() {
		return -1, errors.New("collections.TwowaySearch: item type is not the same as data type")
	}

	forward := 0
	backward := dataVal.Len() - 1

	for forward <= backward {
		forwardVal := dataVal.Index(forward).Interface()
		if fn(forwardVal, item) == 0 {
			return forward, nil
		}

		backwardVal := dataVal.Index(backward).Interface()
		if fn(backwardVal, item) == 0 {
			return backward, nil
		}

		forward++
		backward--
	}

	return -1, errors.New(ErrElementNotFound)
}

func MakeEmptySlice(itemType reflect.Type) interface{} {
	ret := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0).Interface()
	return ret
}

func CloneSlice(data interface{}) interface{} {
	typeOfData := reflect.TypeOf(data)
	if typeOfData.Kind() != reflect.Slice {
		panic("collections.CloneSlice: data must be a slice")
	}
	return reflect.AppendSlice(reflect.New(reflect.SliceOf(typeOfData.Elem())).Elem(), reflect.ValueOf(data)).Interface()
}

func Excepts(first interface{}, second interface{}, fn Comparer) interface{} {
	typeOfFirst := reflect.TypeOf(first)
	if typeOfFirst.Kind() != reflect.Slice {
		panic("collections.Excepts: first must be a slice")
	}
	valOfFirst := reflect.ValueOf(first)
	if valOfFirst.Len() == 0 {
		return MakeEmptySlice(typeOfFirst.Elem())
	}

	typeOfSecond := reflect.TypeOf(second)
	if typeOfSecond.Kind() != reflect.Slice {
		panic("collections.Excepts: second must be a slice")
	}

	valOfSecond := reflect.ValueOf(second)
	if valOfSecond.Len() == 0 {
		return CloneSlice(first)
	}

	result := reflect.New(reflect.SliceOf(typeOfFirst.Elem())).Elem()
	for i := 0; i < valOfFirst.Len(); i++ {
		s := valOfFirst.Index(i).Interface()
		if i, _ := TwowaySearch(second, s, fn); i < 0 {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}

	return result.Interface()
}

func Intersects(first interface{}, second interface{}, fn Comparer) interface{} {
	typeOfFirst := reflect.TypeOf(first)
	if typeOfFirst.Kind() != reflect.Slice {
		panic("collections.Intersects: first must be a slice")
	}
	valOfFirst := reflect.ValueOf(first)
	if valOfFirst.Len() == 0 {
		return MakeEmptySlice(typeOfFirst.Elem())
	}

	typeOfSecond := reflect.TypeOf(second)
	if typeOfSecond.Kind() != reflect.Slice {
		panic("collections.Intersects: second must be a slice")
	}

	valOfSecond := reflect.ValueOf(second)
	if valOfSecond.Len() == 0 {
		return MakeEmptySlice(typeOfFirst.Elem())
	}

	result := reflect.New(reflect.SliceOf(typeOfFirst.Elem())).Elem()
	for i := 0; i < valOfFirst.Len(); i++ {
		s := valOfFirst.Index(i).Interface()
		if i, _ := TwowaySearch(second, s, fn); i >= 0 {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}

	return result.Interface()
}

func Union(first interface{}, second interface{}, fn Comparer) interface{} {
	excepts := Excepts(second, first, fn)

	typeOfFirst := reflect.TypeOf(first)
	if typeOfFirst.Kind() != reflect.Slice {
		panic("collections.Intersects: first must be a slice")
	}
	valOfFirst := reflect.ValueOf(first)
	if valOfFirst.Len() == 0 {
		return CloneSlice(second)
	}

	result := reflect.AppendSlice(reflect.New(reflect.SliceOf(typeOfFirst.Elem())).Elem(), valOfFirst)
	result = reflect.AppendSlice(result, reflect.ValueOf(excepts))
	return result.Interface()
}

func Find(source interface{}, fn Predicate) (interface{}, error) {
	aType := reflect.TypeOf(source)
	if aType.Kind() != reflect.Slice {
		panic("collections.Find: source must be a slice")
	}

	sourceVal := reflect.ValueOf(source)
	if sourceVal.Len() == 0 {
		return nil, errors.New(ErrElementNotFound)
	}

	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if fn(s) {
			return s, nil
		}
	}
	return nil, errors.New(ErrElementNotFound)
}

func FindOrDefault(source interface{}, fn Predicate, defaultValue interface{}) interface{} {
	item, err := Find(source, fn)
	if err != nil {
		if err.Error() == ErrElementNotFound {
			return defaultValue
		}
	}
	return item
}

func TakeWhile(data interface{}, fn Predicate) interface{} {
	aType := reflect.TypeOf(data)
	if aType.Kind() != reflect.Slice {
		panic("collections.TakeWhile: data must be a slice")
	}
	sourceVal := reflect.ValueOf(data)
	if sourceVal.Len() == 0 {
		return MakeEmptySlice(aType.Elem())
	}
	result := reflect.New(reflect.SliceOf(aType.Elem())).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if fn(s) {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}
	return result.Interface()
}

func ExceptWhile(data interface{}, fn Predicate) interface{} {
	aType := reflect.TypeOf(data)
	if aType.Kind() != reflect.Slice {
		panic("collections.ExceptWhile: data must be a slice")
	}
	sourceVal := reflect.ValueOf(data)
	if sourceVal.Len() == 0 {
		return MakeEmptySlice(aType.Elem())
	}
	result := reflect.New(reflect.SliceOf(aType.Elem())).Elem()
	for i := 0; i < sourceVal.Len(); i++ {
		s := sourceVal.Index(i).Interface()
		if !fn(s) {
			result = reflect.Append(result, reflect.ValueOf(s))
		}
	}
	return result.Interface()
}
