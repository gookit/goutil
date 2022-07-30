package maputil

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/arrutil"
)

// Key sep char consts
const (
	KeySepStr  = "."
	KeySepChar = '.'
)

// MergeStringMap simple merge two string map. merge src to dst map
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}

		dst[k] = v
	}
	return dst
}

// GetByPath get value by key path from a map(map[string]any). eg "top" "top.sub"
func GetByPath(path string, mp map[string]any) (val interface{}, ok bool) {
	if val, ok := mp[path]; ok {
		return val, true
	}

	// no sub key
	if len(mp) == 0 || !strings.ContainsRune(path, '.') {
		return nil, false
	}

	// has sub key. eg. "top.sub"
	keys := strings.Split(path, ".")
	topK := keys[0]

	// find top item data use top key
	var item interface{}
	if item, ok = mp[topK]; !ok {
		return
	}

	for _, k := range keys[1:] {
		switch tData := item.(type) {
		case map[string]string: // is simple map
			if item, ok = tData[k]; !ok {
				return
			}
		case map[string]any: // is map(decode from toml/json)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[any]any: // is map(decode from yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		case []any: // is a slice
			if item, ok = getBySlice(k, tData); !ok {
				return
			}
		case []string, []int, []float32, []float64, []bool, []rune:
			slice := reflect.ValueOf(tData)
			sData := make([]interface{}, slice.Len())
			for i := 0; i < slice.Len(); i++ {
				sData[i] = slice.Index(i).Interface()
			}
			if item, ok = getBySlice(k, sData); !ok {
				return
			}
		default: // error
			return nil, false
		}
	}

	return item, true
}

// SetByPath set sub-map value by key path.
// Supports dot syntax to set deep values.
//
// For example:
//
//     SetByPath("name.first", "Mat")
func SetByPath(mp map[string]any, path string, val any) (map[string]interface{}, error) {
	if len(mp) == 0 {
		return MakeByPath(path, val), nil
	}

	_, ok := mp[path]
	// is top key OR no sub key
	if ok || !strings.ContainsRune(path, KeySepChar) {
		mp[path] = val
		return mp, nil
	}

	return SetByKeys(mp, strings.Split(path, KeySepStr), val)
}

// SetByKeys set sub-map value by sub-keys.
// Supports dot syntax to set deep values.
//
// For example:
//
//     SetByPath([]string{"name", "first"}, "Mat")
func SetByKeys(mp map[string]any, keys []string, val any) (map[string]interface{}, error) {
	kln := len(keys)
	if kln == 0 {
		return mp, nil
	}

	if len(mp) == 0 {
		return MakeByKeys(keys, val), nil
	}

	topK := keys[0]
	if kln == 1 {
		mp[topK] = MakeByKeys(keys[1:], val)
		return mp, nil
	}

	var ok bool
	// var item interface{}
	// find top item data use top key
	if _, ok = mp[topK]; !ok {
		mp[topK] = MakeByKeys(keys[1:], val)
		return mp, nil
	}

	// reflect.MapOf()

	obj := mp
	max := len(keys) - 1

	for index, field := range keys {
		if index == max {
			obj[field] = val
		}

		if _, exists := obj[field]; !exists {
			obj[field] = make(Data)
			obj = obj[field].(Data)
		} else {
			switch typData := obj[field].(type) {
			case Data:
				// obj = obj[field].(Data)
				obj = typData
			case map[string]any:
				// obj = obj[field].(map[string]any)
				obj = typData
				// case map[string]string:
				// 	obj = typData
			}
		}
	}

	return mp, nil
}

// MakeByPath build new value by key names
//
// Example:
// 	"site.info"
// 	->
// 	map[string]interface{} {
//		site: {info: val}
// 	}
func MakeByPath(path string, val interface{}) (mp map[string]interface{}) {
	return MakeByKeys(strings.Split(path, KeySepStr), val)
}

// MakeByKeys build new value by key names
//
// Example:
// 	"site.info"
// 	->
// 	map[string]interface{} {
//		site: {info: val}
// 	}
func MakeByKeys(keys []string, val any) (mp map[string]interface{}) {
	if len(keys) == 1 {
		return map[string]any{keys[0]: val}
	}

	// multi nodes
	arrutil.Reverse(keys)
	for _, p := range keys {
		if mp == nil {
			mp = map[string]any{p: val}
		} else {
			mp = map[string]any{p: mp}
		}
	}
	return
}

// Keys get all keys of the given map.
func Keys(mp any) (keys []string) {
	rftVal := reflect.Indirect(reflect.ValueOf(mp))
	if rftVal.Kind() != reflect.Map {
		return
	}

	keys = make([]string, 0, rftVal.Len())
	for _, key := range rftVal.MapKeys() {
		keys = append(keys, key.String())
	}
	return
}

// Values get all values from the given map.
func Values(mp any) (values []interface{}) {
	rftVal := reflect.Indirect(reflect.ValueOf(mp))
	if rftVal.Kind() != reflect.Map {
		return
	}

	values = make([]any, 0, rftVal.Len())
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}

func getBySlice(k string, slice []any) (val interface{}, ok bool) {
	i, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return nil, false
	}
	if size := int64(len(slice)); i >= size {
		return nil, false
	}
	return slice[i], true
}
