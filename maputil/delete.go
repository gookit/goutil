package maputil

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/reflects"
)

// DeleteByPath delete value by key path from a map(map[string]any).eg "top" "top.sub"
//
// Example:
//
//	mp := map[string]any{
//		"top": map[string]any{
//			"sub": "value",
//		},
//	}
//	ok := DeleteByPath(mp, "top.sub") // return true
func DeleteByPath(mp map[string]any, path string) bool {
	if _, ok := mp[path]; ok {
		delete(mp, path)
		return true
	}

	_, ok := deleteByPathKeys(mp, mp, strings.Split(path, KeySepStr))
	return ok
}

// deleteByPathKeys delete value by key path from a map(map[string]any).eg "top" "top.sub"
//
// Example:
//
//	mp := map[string]any{
//		"top": map[string]any{
//			"sub": "value",
//		},
//	}
//	val, ok := deleteByPathKeys(mp, mp, []string{"top", "sub"}) // return "value", true
func deleteByPathKeys(parent, child any, keys []string) (any, bool) {
	var (
		prevLevel, currLevel any
		ok                   bool
		prevKey              string
	)

	kl := len(keys)

	prevLevel = parent
	currLevel = child

	for i, k := range keys {
		switch tData := currLevel.(type) {
		case map[string]string:
			if _, ok = tData[k]; !ok {
				return nil, false
			}
			prevLevel = currLevel
			currLevel = tData[k]
			prevKey = k
			if kl == i+1 {
				delete(tData, k)
				return currLevel, true
			}
		case map[string]any:
			if _, ok = tData[k]; !ok {
				return nil, false
			}
			prevLevel = currLevel
			currLevel = tData[k]
			prevKey = k
			if kl == i+1 {
				delete(tData, k)
				return currLevel, true
			}
		case map[any]any:
			// Check if the key exists in the map
			if val, ok := tData[k]; ok {
				prevLevel = currLevel
				currLevel = val
				prevKey = k

				// If it's the last key, delete it and return the current level
				if kl == i+1 {
					delete(tData, k)
					return currLevel, true
				}
			} else {
				// Try converting the key to an integer
				if idx, err := strconv.Atoi(k); err == nil {
					prevLevel = currLevel
					currLevel = tData[idx]
					prevKey = k

					// If it's the last key, delete it and return the current level
					if kl == i+1 {
						delete(tData, idx)
						return currLevel, true
					}
				}
			}
		case map[string]int:
			if _, ok = tData[k]; !ok {
				return nil, false
			}
			prevLevel = currLevel
			currLevel = tData[k]
			prevKey = k
			if kl == i+1 {
				delete(tData, k)
				return currLevel, true
			}
		default:
			rv := reflect.ValueOf(tData)
			// check is slice
			if rv.Kind() == reflect.Slice {
				if k == Wildcard {
					if kl == i+1 { // * is last key
						rv = rv.Slice(0, 0)
						reflect.ValueOf(prevLevel).SetMapIndex(reflect.ValueOf(prevKey), rv)
						return currLevel, true
					}

					isDeleted := false
					for si := 0; si < rv.Len(); si++ {
						el := reflects.Indirect(rv.Index(si))
						if el.Kind() != reflect.Map {
							return nil, false
						}

						// el is map value.
						if _, ok := deleteByPathKeys(prevLevel, el.Interface(), keys[i+1:]); ok {
							if reflects.IsEmpty(el) {
								if rv.Len() > 1 {
									rv = reflect.AppendSlice(rv.Slice(0, si), rv.Slice(si+1, rv.Len()))
									si--
								} else {
									rv = reflect.MakeSlice(rv.Type(), 0, 0)
									break
								}
							}
							isDeleted = true
						}
					}

					currLevel = rv.Interface()
					reflect.ValueOf(prevLevel).SetMapIndex(reflect.ValueOf(prevKey), rv)
					if rv.Len() > 0 {
						return currLevel, true
					}

					return nil, isDeleted
				}

				ii, err := strconv.Atoi(k)
				if err != nil || ii >= rv.Len() {
					return nil, false
				}

				currLevel = rv.Index(ii).Interface()

				if kl == i+1 {
					rv = reflect.AppendSlice(rv.Slice(0, ii), rv.Slice(ii+1, rv.Len()))
					reflect.ValueOf(prevLevel).SetMapIndex(reflect.ValueOf(prevKey), rv)
					return currLevel, true
				}
				continue
			}
		}
	}

	return nil, false
}
