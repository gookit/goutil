package arrutil

import "reflect"

func Aggregate(data interface{}, fn func(seed, element interface{}) interface{}) interface{} {
	if data == nil {
		return nil
	}
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Slice {
		panic("arrutil.Aggregate: source must be a slice")
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Len() == 0 {
		return nil
	}
	itemType := dataValue.Index(0).Type()
	seed := reflect.New(reflect.SliceOf(itemType)).Elem()
	for i := 0; i < dataValue.Len(); i++ {
		v := dataValue.Index(i).Interface()
		ret := fn(seed.Interface(), v)
		retType := reflect.TypeOf(ret)
		if retType.Kind() == reflect.Slice {
			retSliceVal := reflect.ValueOf(ret)
			for j := 0; j < retSliceVal.Len(); j++ {
				seed = reflect.Append(seed, retSliceVal.Index(j))
			}
		} else {
			seed = reflect.Append(seed, reflect.ValueOf(ret))
		}
	}
	return seed.Interface()
}
