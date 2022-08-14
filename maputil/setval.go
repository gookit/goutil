package maputil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/strutil"
)

// SetByPath set sub-map value by key path.
// Supports dot syntax to set deep values.
//
// For example:
//
//	SetByPath("name.first", "Mat")
func SetByPath(mp *map[string]any, path string, val any) error {
	return SetByKeys(mp, strings.Split(path, KeySepStr), val)
}

// SetByKeys set sub-map value by path keys.
// Supports dot syntax to set deep values.
//
// For example:
//
//	SetByKeys([]string{"name", "first"}, "Mat")
func SetByKeys(mp *map[string]any, keys []string, val any) (err error) {
	kln := len(keys)
	if kln == 0 {
		return nil
	}

	mpv := *mp
	if len(mpv) == 0 {
		*mp = MakeByKeys(keys, val)
		return nil
	}

	topK := keys[0]
	if kln == 1 {
		mpv[topK] = val
		return nil
	}

	if _, ok := mpv[topK]; !ok {
		mpv[topK] = MakeByKeys(keys[1:], val)
		return nil
	}

	rv := reflect.ValueOf(mp).Elem()
	return setMapByKeys(rv, keys, reflect.ValueOf(val))
}

func setMapByKeys(rv reflect.Value, keys []string, nv reflect.Value) (err error) {

	if rv.Kind() != reflect.Map {
		return fmt.Errorf("input rv value type must be Map, but was %s", rv.Kind())
	}

	maxI := len(keys) - 1
	for i, key := range keys {
		idx := -1
		isPtr := false
		isMap := rv.Kind() == reflect.Map
		isSlice := rv.Kind() == reflect.Slice
		isLast := i == len(keys)-1

		// slice index key must be ended on the keys.
		// eg: "top.arr[2]" -> "arr[2]"
		if pos := strings.IndexRune(key, '['); pos > 0 {
			idx, err = strconv.Atoi(key[pos+1 : len(key)-1])
			if err != nil {
				err = fmt.Errorf("invalid array index on key: %s", key)
				break
			}
			key = key[:pos]

			// update value
			if !isMap {
				err = fmt.Errorf(
					"current value#%s type is %s, cannot get sub-value by key: %s",
					strings.Join(keys[i:], "."),
					rv.Kind(),
					key,
				)
				break
			}

			rftK := reflect.ValueOf(key)
			tmpV := rv.MapIndex(rftK)
			if !tmpV.IsValid() {
				// deep make map by keys
				newVal := MakeByKeys(keys[i:], nv.Interface())
				rv.SetMapIndex(rftK, reflect.ValueOf(newVal))
				break
			}

			// get real type: any -> map
			if tmpV.Kind() == reflect.Interface {
				tmpV = tmpV.Elem()
			}

			if tmpV.Kind() != reflect.Slice {
				err = fmt.Errorf(
					"current value#%s type is %s, cannot set sub by index: %d",
					strings.Join(keys[i:], "."),
					tmpV.Kind(),
					idx,
				)
				break
			}

			wantLen := idx + 1
			sliLen := tmpV.Len()
			elemTyp := tmpV.Type().Elem()

			if wantLen > sliLen {
				newAdd := reflect.MakeSlice(tmpV.Type(), 0, wantLen-sliLen)

				for i := 0; i < wantLen-sliLen; i++ {
					newAdd = reflect.Append(newAdd, reflect.New(elemTyp).Elem())
				}

				tmpV = reflect.AppendSlice(tmpV, newAdd)
			}

			if !isLast {
				if elemTyp.Kind() == reflect.Map {
					err := setMapByKeys(tmpV.Index(idx), keys[i:], nv)
					if err != nil {
						return err
					}

					// tmpV.Index(idx).Set(elemV)
					rv.SetMapIndex(rftK, tmpV)
				} else {
					err = fmt.Errorf(
						"key %s[%d] elem must be map for set sub-value by remain path: %s",
						key,
						idx,
						strings.Join(keys[i:], "."),
					)
				}
			} else {
				// last - set value
				tmpV.Index(idx).Set(nv)
				rv.SetMapIndex(rftK, tmpV)
			}

			break
		}

		// set value on last key
		if isLast {
			if isMap {
				rv.SetMapIndex(reflect.ValueOf(key), nv)
				break
			}

			if isSlice {
				// key is slice index
				if strutil.IsNumeric(key) {
					idx, _ = strconv.Atoi(key)
				}

				if idx > -1 {
					wantLen := idx + 1
					sliLen := rv.Len()

					if wantLen > sliLen {
						elemTyp := rv.Type().Elem()
						newAdd := reflect.MakeSlice(rv.Type(), 0, wantLen-sliLen)

						for i := 0; i < wantLen-sliLen; i++ {
							newAdd = reflect.Append(newAdd, reflect.New(elemTyp).Elem())
						}

						if !rv.CanAddr() {
							err = fmt.Errorf("cannot set value to a cannot addr slice, key: %s", key)
							break
						}

						rv.Set(reflect.AppendSlice(rv, newAdd))
					}

					rv.Index(idx).Set(nv)
				} else {
					err = fmt.Errorf("cannot set slice value by named key %q", key)
				}
				break
			}

			err = fmt.Errorf(
				"cannot set sub-value for type %q(path %q, key %q)",
				rv.Kind(),
				strings.Join(keys[:i], "."),
				key,
			)
			break
		}

		if isMap {
			rftK := reflect.ValueOf(key)
			if tmpV := rv.MapIndex(rftK); tmpV.IsValid() {
				// get real type: any -> map
				tmpV, isPtr = getRealVal(tmpV)
				if tmpV.Kind() == reflect.Map {
					rv = tmpV
					continue
				}

				// sub is slice and is not ptr
				if tmpV.Kind() == reflect.Slice {
					if isPtr {
						rv = tmpV
						continue // to (E)
					}

					// next key is index number.
					nxtKey := keys[i+1]
					if strutil.IsNumeric(nxtKey) {
						idx, _ = strconv.Atoi(nxtKey)
						sliLen := tmpV.Len()
						wantLen := idx + 1

						if wantLen > sliLen {
							elemTyp := tmpV.Type().Elem()
							newAdd := reflect.MakeSlice(tmpV.Type(), 0, wantLen-sliLen)
							for i := 0; i < wantLen-sliLen; i++ {
								newAdd = reflect.Append(newAdd, reflect.New(elemTyp).Elem())
							}

							tmpV = reflect.AppendSlice(tmpV, newAdd)
						}

						// rv = tmpV.Index(idx) // TODO
						if i+1 == maxI {
							tmpV.Index(idx).Set(nv)
						} else {
							err := setMapByKeys(tmpV.Index(idx), keys[i+1:], nv)
							if err != nil {
								return err
							}
						}

						rv.SetMapIndex(rftK, tmpV)
						break

					} else {
						err = fmt.Errorf("cannot set slice value by named key %s(parent: %s)", nxtKey, key)
						break
					}

					// reflect.MakeSlice() // TODO
				} else {

					err = fmt.Errorf(
						"map item type is %s, cannot set sub-value by path %q",
						rv.Kind(),
						strings.Join(keys[i:], "."),
					)
					break
				}

				// rv = tmpV
			} else {
				// deep make map by keys
				newVal := MakeByKeys(keys[i:], nv.Interface())
				rv.SetMapIndex(rftK, reflect.ValueOf(newVal))
				break
			}
		} else if isSlice && strutil.IsNumeric(key) { // (E). slice from ptr slice
			idx, _ = strconv.Atoi(key)
			sliLen := rv.Len()
			wantLen := idx + 1

			if wantLen > sliLen {
				elemTyp := rv.Type().Elem()
				newAdd := reflect.MakeSlice(rv.Type(), 0, wantLen-sliLen)
				for i := 0; i < wantLen-sliLen; i++ {
					newAdd = reflect.Append(newAdd, reflect.New(elemTyp).Elem())
				}

				rv = reflect.AppendSlice(rv, newAdd)
			}

			rv = rv.Index(idx)
		} else {
			err = fmt.Errorf(
				"map item type is %s, cannot set sub-value by path %q",
				rv.Kind(),
				strings.Join(keys[i:], "."),
			)
			break
		}

		// TODO remove it
		// dump.P(key, isPtr, rv.CanAddr())
	}
	return
}

func getRealVal(rv reflect.Value) (reflect.Value, bool) {
	// get real type: any -> map
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	isPtr := false
	if rv.Kind() == reflect.Ptr {
		isPtr = true
		rv = rv.Elem()
	}

	return rv, isPtr
}

// "arr[2]" => "arr", 2, true
func parseArrKeyIndex(key string) (string, int, bool) {
	pos := strings.IndexRune(key, '[')
	if pos < 1 || !strings.HasSuffix(key, "]") {
		return key, 0, false
	}

	var idx int
	var err error

	idxStr := key[pos+1 : len(key)-1]
	if idxStr != "" {
		idx, err = strconv.Atoi(idxStr)
		if err != nil {
			return key, 0, false
		}
	}

	key = key[:pos]
	return key, idx, true
}
