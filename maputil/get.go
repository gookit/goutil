package maputil

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/reflects"
)

// some consts for separators
const (
	Wildcard = "*"
	PathSep  = "."
)

// DeepGet value by key path. eg "top" "top.sub"
func DeepGet(mp map[string]any, path string) (val any) {
	val, _ = GetByPath(path, mp)
	return
}

// QuietGet value by key path. eg "top" "top.sub"
func QuietGet(mp map[string]any, path string) (val any) {
	val, _ = GetByPath(path, mp)
	return
}

// GetFromAny get value by key path from any(map,slice) data. eg "top" "top.sub"
func GetFromAny(path string, data any) (val any, ok bool) {
	// empty data
	if data == nil {
		return nil, false
	}
	if len(path) == 0 {
		return data, true
	}

	keys := strings.Split(path, ".")
	if val, ok = getByPathKeys(data, keys); ok {
		return
	}

	// retry with the slice indexes expanded in the key segments. eg: "arr[2]"
	if hasArrIndexIn(keys) {
		if exp := expandPathKeys(keys); len(exp) > len(keys) {
			return getByPathKeys(data, exp)
		}
	}
	return nil, false
}

// GetByPath get value by key path from a map(map[string]any). eg "top" "top.sub"
func GetByPath(path string, mp map[string]any) (val any, ok bool) {
	if len(path) == 0 {
		return mp, true
	}
	if val, ok := mp[path]; ok {
		return val, true
	}

	// no sub key. eg: "top", "[0]"
	if len(mp) == 0 || (strings.IndexByte(path, '.') < 1 && !hasArrIndex(path)) {
		return nil, false
	}

	// key is path. eg: "top.sub", "top.arr[2]"
	return GetByPathKeys(mp, strings.Split(path, "."))
}

// GetByPathKeys get value by path keys from a map(map[string]any). eg "top" "top.sub"
//
// Example:
//
//	mp := map[string]any{
//		"top": map[string]any{
//			"sub": "value",
//		},
//	}
//	val, ok := GetByPathKeys(mp, []string{"top", "sub"}) // return "value", true
func GetByPathKeys(mp map[string]any, keys []string) (val any, ok bool) {
	kl := len(keys)
	if kl == 0 {
		return mp, true
	}

	if val, ok = getByPathKeysOfMap(mp, keys); ok {
		return
	}

	// retry with the slice indexes expanded in the key segments. eg: "arr[2]"
	// NOTICE: the literal key is matched first, so a key that really contains
	// brackets (eg. "child[x]", "a[0]") keeps working as before.
	if hasArrIndexIn(keys) {
		if exp := expandPathKeys(keys); len(exp) > kl {
			return getByPathKeysOfMap(mp, exp)
		}
	}
	return nil, false
}

// getByPathKeysOfMap get value from a map by the plain key path, no index expansion.
func getByPathKeysOfMap(mp map[string]any, keys []string) (val any, ok bool) {
	// find top item data use top key
	var item any
	if item, ok = mp[keys[0]]; !ok {
		return
	}

	// find sub item data use sub key
	return getByPathKeys(item, keys[1:])
}

// hasArrIndexIn check if any key segment contains a slice index. eg: "arr[2]"
func hasArrIndexIn(keys []string) bool {
	for _, k := range keys {
		if hasArrIndex(k) {
			return true
		}
	}
	return false
}

// expandPathKeys expands the trailing slice indexes of the key segments into
// separate numeric keys. eg:
//
//	["top", "arr[2]"]  -> ["top", "arr", "2"]
//	["top", "m[0][1]"] -> ["top", "m", "0", "1"]
//
// A segment without a valid trailing index is kept as-is.
func expandPathKeys(keys []string) (expanded []string) {
	expanded = make([]string, 0, len(keys))
	for _, seg := range keys {
		expanded = append(expanded, expandSegIndex(seg)...)
	}
	return
}

// expandSegIndex splits the trailing "[N]" indexes of a key segment.
// eg: "arr[2]" -> ["arr", "2"], "matrix[0][1]" -> ["matrix", "0", "1"].
// The segment is returned unchanged when it has no valid trailing index, so
// a key with a literal bracket (eg. "child[x]", "arr[-1]") is kept as-is.
func expandSegIndex(seg string) []string {
	var idxs []string
	for hasArrIndex(seg) {
		pos := strings.LastIndexByte(seg, '[')
		if pos <= 0 { // eg: "[0]", the index must follow a key name
			break
		}

		idx := seg[pos+1 : len(seg)-1]
		if i, err := strconv.Atoi(idx); err != nil || i < 0 {
			break
		}

		idxs = append([]string{idx}, idxs...)
		seg = seg[:pos]
	}

	if len(idxs) == 0 {
		return []string{seg}
	}
	return append([]string{seg}, idxs...)
}

func getByPathKeys(item any, keys []string) (val any, ok bool) {
	kl := len(keys)

	for i, k := range keys {
		switch tData := item.(type) {
		case map[string]string: // is string map
			if item, ok = tData[k]; !ok {
				return
			}
		case map[string]any: // is map(decode from toml/json/yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[any]any: // is map(decode from yaml.v2)
			if item, ok = tData[k]; !ok {
				return
			}
		case []map[string]any: // is an any-map slice
			if k == Wildcard {
				if kl == i+1 { // * is last key
					return tData, true
				}

				// * is not last key, find sub item data
				sl := make([]any, 0, len(tData))
				for _, v := range tData {
					if val, ok = getByPathKeys(v, keys[i+1:]); ok {
						sl = append(sl, val)
					}
				}

				if len(sl) > 0 {
					return sl, true
				}
				return nil, false
			}

			// k is index number
			idx, err := strconv.Atoi(k)
			if err != nil || idx < 0 || idx >= len(tData) {
				return nil, false
			}
			item = tData[idx]
		default:
			if k == Wildcard && kl == i+1 { // * is last key
				return tData, true
			}

			rv := reflect.ValueOf(tData)

			// a typed map that the cases above do not cover. eg: map[string]int
			if rv.Kind() == reflect.Map {
				if v := mapValueByStringKey(rv, k); v.IsValid() {
					item = v.Interface()
					continue
				}
				return nil, false
			}

			// check is slice
			if rv.Kind() == reflect.Slice {
				if k == Wildcard {
					// * is not last key, find sub item data
					sl := make([]any, 0, rv.Len())
					for si := 0; si < rv.Len(); si++ {
						el := reflects.Indirect(rv.Index(si))
						if el.Kind() != reflect.Map {
							return nil, false
						}

						// el is map value.
						if val, ok = getByPathKeys(el.Interface(), keys[i+1:]); ok {
							sl = append(sl, val)
						}
					}

					if len(sl) > 0 {
						return sl, true
					}
					return nil, false
				}

				// check k is index number
				ii, err := strconv.Atoi(k)
				if err != nil || ii < 0 || ii >= rv.Len() {
					return nil, false
				}

				item = rv.Index(ii).Interface()
				continue
			}

			// as error
			return nil, false
		}

		// next is last key and it is *
		if kl == i+2 && keys[i+1] == Wildcard {
			return item, true
		}
	}

	return item, true
}

// mapValueByStringKey returns the value of a string path segment in a map whose key
// type is a string kind, so a typed map like map[string]int resolves too. The zero
// Value means the key is not found (or the map is keyed by something else).
func mapValueByStringKey(rv reflect.Value, segment string) reflect.Value {
	keyTyp := rv.Type().Key()
	if keyTyp.Kind() != reflect.String {
		return reflect.Value{}
	}

	key := reflect.ValueOf(segment)
	if !key.Type().AssignableTo(keyTyp) {
		key = key.Convert(keyTyp)
	}
	return rv.MapIndex(key)
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

// TypedKeys get all keys of the given typed map.
func TypedKeys[K comdef.SimpleType, V any](mp map[K]V) (keys []K) {
	for key := range mp {
		keys = append(keys, key)
	}
	return
}

// FirstKey returns the first key of the given map.
func FirstKey[T any](mp map[string]T) string {
	for key := range mp {
		return key
	}
	return ""
}

// Values get all values from the given map.
func Values(mp any) (values []any) {
	rv := reflect.Indirect(reflect.ValueOf(mp))
	if rv.Kind() != reflect.Map {
		return
	}

	values = make([]any, 0, rv.Len())
	for _, key := range rv.MapKeys() {
		values = append(values, rv.MapIndex(key).Interface())
	}
	return
}

// TypedValues get all values from the given typed map.
func TypedValues[K comdef.SimpleType, V any](mp map[K]V) (values []V) {
	for _, val := range mp {
		values = append(values, val)
	}
	return
}

// EachAnyMap iterates the given map and calls the given function for each item.
func EachAnyMap(mp any, fn func(key string, val any)) {
	rv := reflect.Indirect(reflect.ValueOf(mp))
	if rv.Kind() != reflect.Map {
		panic("not a map value")
	}

	for _, key := range rv.MapKeys() {
		fn(key.String(), rv.MapIndex(key).Interface())
	}
}

// EachTypedMap iterates the given map and calls the given function for each item.
func EachTypedMap[K comdef.SimpleType, V any](mp map[K]V, fn func(key K, val V)) {
	for key, val := range mp {
		fn(key, val)
	}
}
