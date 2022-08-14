package maputil

import (
	"reflect"
	"strings"

	"github.com/gookit/goutil/arrutil"
)

// Key, value sep char consts
const (
	ValSepStr  = ","
	ValSepChar = ','
	KeySepStr  = "."
	KeySepChar = '.'
)

// MergeSMap simple merge two string map. merge src to dst map
func MergeSMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	return MergeStringMap(src, dst, ignoreCase)
}

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

// MakeByPath build new value by key names
//
// Example:
//
//	"site.info"
//	->
//	map[string]interface{} {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	"site.tags[1]"
//	->
//	map[string]interface{} {
//		site: {tags: [val]}
//	}
func MakeByPath(path string, val interface{}) (mp map[string]interface{}) {
	return MakeByKeys(strings.Split(path, KeySepStr), val)
}

// MakeByKeys build new value by key names
//
// Example:
//
//	// case 1:
//	[]string{"site", "info"}
//	->
//	map[string]interface{} {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	[]string{"site", "tags[1]"}
//	->
//	map[string]interface{} {
//		site: {tags: [val]}
//	}
func MakeByKeys(keys []string, val any) (mp map[string]interface{}) {
	size := len(keys)

	// if last key contains slice index, make slice wrap the val
	lastKey := keys[size-1]
	if newK, idx, ok := parseArrKeyIndex(lastKey); ok {
		// valTyp := reflect.TypeOf(val)
		sliTyp := reflect.SliceOf(reflect.TypeOf(val))
		sliVal := reflect.MakeSlice(sliTyp, idx+1, idx+1)
		sliVal.Index(idx).Set(reflect.ValueOf(val))

		// update val and last key
		val = sliVal.Interface()
		keys[size-1] = newK
	}

	if size == 1 {
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
