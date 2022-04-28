package maputil

import (
	"reflect"
	"strconv"
	"strings"
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

// GetByPath get value from a map[string]interface{}. eg "top" "top.sub"
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool) {
	if val, ok := mp[key]; ok {
		return val, true
	}

	// has sub key? eg. "top.sub"
	if !strings.ContainsRune(key, '.') {
		return nil, false
	}

	keys := strings.Split(key, ".")
	topK := keys[0]

	// find top item data based on top key
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
		case map[string]interface{}: // is map(decode from toml/json)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[interface{}]interface{}: // is map(decode from yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		case []interface{}: // is an slice
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

// Keys get all keys of the given map.
func Keys(mp interface{}) (keys []string) {
	rftVal := reflect.ValueOf(mp)
	if rftVal.Type().Kind() == reflect.Ptr {
		rftVal = rftVal.Elem()
	}

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
func Values(mp interface{}) (values []interface{}) {
	rftTyp := reflect.TypeOf(mp)
	if rftTyp.Kind() == reflect.Ptr {
		rftTyp = rftTyp.Elem()
	}

	if rftTyp.Kind() != reflect.Map {
		return
	}

	rftVal := reflect.ValueOf(mp)
	values = make([]interface{}, 0, rftVal.Len())
	for _, key := range rftVal.MapKeys() {
		values = append(values, rftVal.MapIndex(key).Interface())
	}
	return
}

func getBySlice(k string, slice []interface{}) (val interface{}, ok bool) {
	i, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		return nil, false
	}
	if size := int64(len(slice)); i >= size {
		return nil, false
	}
	return slice[i], true
}
