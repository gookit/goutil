package maputil

import (
	"strings"

	"github.com/gookit/goutil/arrutil"
)

// Key sep char consts
const (
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
