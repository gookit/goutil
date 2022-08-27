package structs

import (
	"errors"
	"reflect"
)

// ToMap simple convert structs to map by reflect
func ToMap(st interface{}) map[string]interface{} {
	mp, _ := StructToMap(st, nil)
	return mp
}

// TryToMap simple convert structs to map by reflect
func TryToMap(st interface{}) (map[string]interface{}, error) {
	return StructToMap(st, nil)
}

// MustToMap alis of TryToMap, but will panic on error
func MustToMap(st interface{}) map[string]interface{} {
	mp, err := StructToMap(st, nil)
	if err != nil {
		panic(err)
	}
	return mp
}

const defaultFieldTag = "json"

// MapOptions struct
type MapOptions struct {
	TagName string
}

// StructToMap simple convert structs to map by reflect
func StructToMap(st interface{}, opt *MapOptions) (map[string]interface{}, error) {
	mp := make(map[string]interface{})
	if st == nil {
		return mp, nil
	}

	obj := reflect.ValueOf(st)
	if obj.Kind() == reflect.Ptr {
		obj = obj.Elem()
	}

	if obj.Kind() != reflect.Struct {
		return mp, errors.New("must be an struct")
	}

	if opt == nil {
		opt = &MapOptions{TagName: defaultFieldTag}
	} else if opt.TagName == "" {
		opt.TagName = defaultFieldTag
	}

	mp, err := structToMap(obj, opt.TagName)
	return mp, err
}

func structToMap(obj reflect.Value, tagName string) (map[string]interface{}, error) {
	refType := obj.Type()
	mp := make(map[string]interface{})

	for i := 0; i < obj.NumField(); i++ {
		ft := refType.Field(i)
		name := ft.Name
		// skip don't exported field
		if name[0] >= 'a' && name[0] <= 'z' {
			continue
		}

		tagVal, ok := ft.Tag.Lookup(tagName)
		if ok && tagVal != "" {
			sMap, err := ParseTagValueDefault(name, tagVal)
			if err != nil {
				return nil, err
			}

			name = sMap.Default("name", name)
			// un-exported field
			if name == "" {
				continue
			}
		}

		field := obj.Field(i)
		if field.Kind() == reflect.Struct {
			sub, err := structToMap(field, tagName)
			if err != nil {
				return nil, err
			}
			mp[name] = sub
			continue
		}

		if field.CanInterface() {
			mp[name] = field.Interface()
		}
	}

	return mp, nil
}
