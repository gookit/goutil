package comfunc

import (
	"errors"
	"reflect"
)

// TryStructToMap simple convert structs to map by reflect
func TryStructToMap(st interface{}) (map[string]interface{}, error) {
	mp := make(map[string]interface{})
	if st == nil {
		return mp, nil
	}

	obj := reflect.ValueOf(st)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	refType := obj.Type()
	if refType.Kind() != reflect.Struct {
		return mp, errors.New("must be an struct")
	}

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		if field.CanInterface() {
			mp[refType.Field(i).Name] = field.Interface()
		}
	}

	return mp, nil
}
