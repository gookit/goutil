package maputil

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/strutil"
)

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
//     SetByKeys([]string{"name", "first"}, "Mat")
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
		mp[topK] = val
		return mp, nil
	}

	if _, ok := mp[topK]; !ok {
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
			}
		}
	}

	return mp, nil
}

// SetByKeys2 set sub-map value by sub-keys.
// Supports dot syntax to set deep values.
//
// For example:
//
//     SetByKeys2(mp, []string{"name", "first"}, "Mat")
func SetByKeys2(mp *map[string]any, keys []string, val any) (err error) {
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
				err = errorx.Wrapf(err, "invalid array index on key: %s", key)
				break
			}
			key = key[:pos]

			// update value
			if !isMap {
				err = errorx.Rawf(
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
				err = errorx.Rawf(
					"current value#%s type is %s, cannot set sub by index: %d",
					strings.Join(keys[i:], "."),
					tmpV.Kind(),
					idx,
				)
				break
			}

			if !isLast {
				err = errorx.Rawf(
					"key %s[%d] must be at end on set slice item. but was remaining path: %s",
					key,
					idx,
					strings.Join(keys[i:], "."),
				)
			} else {
				wantLen := idx + 1
				sliLen := tmpV.Len()

				if wantLen > sliLen {
					elemTyp := tmpV.Type().Elem()
					newAdd := reflect.MakeSlice(tmpV.Type(), 0, wantLen-sliLen)

					for i := 0; i < wantLen-sliLen; i++ {
						newAdd = reflect.Append(newAdd, reflect.New(elemTyp).Elem())
					}

					tmpV = reflect.AppendSlice(tmpV, newAdd)
				}

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

						rv.Set(reflect.AppendSlice(rv, newAdd))
					}

					rv.Index(idx).Set(nv)
				} else {
					err = errorx.Rawf("cannot set slice value by named key %q", key)
				}
				break
			}

			err = errorx.Rawf(
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
				rv, isPtr = getRealVal(tmpV)

				if !isPtr && rv.Kind() == reflect.Slice {
					// reflect.MakeSlice() // TODO
				}

				// rv = tmpV
			} else {
				// deep make map by keys
				newVal := MakeByKeys(keys[i:], nv.Interface())
				rv.SetMapIndex(rftK, reflect.ValueOf(newVal))
				break
			}
		} else if isSlice && strutil.IsNumeric(key) { // slice
			idx, _ = strconv.Atoi(key)
			sliLen := rv.Len()
			wantLen := idx + 1

			if wantLen > sliLen {
				elemTyp := rv.Type().Elem()
				newAdd := reflect.MakeSlice(rv.Type(), 0, wantLen-sliLen)
				for i := 0; i < wantLen-sliLen; i++ {
					newAdd = reflect.Append(newAdd, reflect.New(elemTyp))
				}

				rv = reflect.AppendSlice(rv, newAdd)
			}

			rv = rv.Index(idx)
		} else {
			err = errorx.Rawf(
				"map item type is %s, cannot set sub-value by path %q",
				rv.Kind(),
				strings.Join(keys[i:], "."),
			)
			break
		}

		// TODO remove it
		dump.P(key, isPtr, rv.CanAddr())
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
