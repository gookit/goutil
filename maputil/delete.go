package maputil

import (
	"fmt"
	"reflect"
	"strings"
)

// DeleteByPath deletes the value at the given key path.
//
// It supports the dot syntax and the slice index syntax of SetByPath, so
// "list[1]" and "list.1" both delete the element at index 1 and shrink the slice.
// Deleting a path that does not exist is not an error.
//
// Example:
//
//	DeleteByPath(&mp, "user.name")
func DeleteByPath(mp *map[string]any, path string) error {
	return DeleteByKeys(mp, strings.Split(path, KeySepStr))
}

// DeleteByKeys deletes the value at the given path keys.
//
// It is the path form of DeleteByPath, see the example there.
func DeleteByKeys(mp *map[string]any, keys []string) error {
	if mp == nil || *mp == nil || len(keys) == 0 {
		return nil
	}

	rv := reflect.ValueOf(mp).Elem()
	return deleteByKeys(rv, expandPathKeys(keys))
}

// deleteByKeys deletes the first key from the map value rv, then goes on with the rest.
func deleteByKeys(rv reflect.Value, keys []string) error {
	key, rest := keys[0], keys[1:]

	kv, ok := mapKeyValue(rv.Type().Key(), key)
	if !ok {
		return fmt.Errorf("cannot delete key %q from a %s", key, rv.Type())
	}

	child := rv.MapIndex(kv)
	if !child.IsValid() { // the key does not exist
		return nil
	}

	if len(rest) == 0 {
		rv.SetMapIndex(kv, reflect.Value{}) // delete the key
		return nil
	}

	newChild, replace, err := deleteFrom(child, rest)
	if err != nil || !replace {
		return err
	}
	rv.SetMapIndex(kv, newChild)
	return nil
}

// deleteFrom deletes the keys from the child value, which may be an interface value.
// When the removal shrinks a slice, replace is true and newChild is the value to store
// back to the parent. A map is changed in place, so nothing has to be stored back.
func deleteFrom(child reflect.Value, keys []string) (newChild reflect.Value, replace bool, err error) {
	if child.Kind() == reflect.Interface {
		if child.IsNil() {
			return child, false, nil
		}

		newV, replace, err := deleteFrom(child.Elem(), keys)
		if err != nil || !replace {
			return child, false, err
		}
		return newV, true, nil
	}

	switch child.Kind() {
	case reflect.Map:
		return child, false, deleteByKeys(child, keys)
	case reflect.Slice:
		idx, err := parseIndexKey(keys[0])
		if err != nil {
			return child, false, fmt.Errorf("cannot delete key %q from a %s", keys[0], child.Type())
		}
		if idx >= child.Len() {
			return child, false, nil
		}

		// delete the element and shrink the slice
		if len(keys) == 1 {
			newSli := reflect.AppendSlice(child.Slice(0, idx), child.Slice(idx+1, child.Len()))
			return newSli, true, nil
		}

		// delete a value of the element. eg: "list[0].name"
		newElem, elemReplace, err := deleteFrom(child.Index(idx), keys[1:])
		if err != nil {
			return child, false, err
		}
		if elemReplace {
			child.Index(idx).Set(newElem)
		}
		return child, false, nil
	}

	return child, false, fmt.Errorf("cannot delete %q from a %s value", strings.Join(keys, "."), child.Type())
}

// mapKeyValue converts the string key to the key type of the map.
// It reports false when the map is not keyed by a string kind type.
func mapKeyValue(keyTyp reflect.Type, key string) (reflect.Value, bool) {
	if keyTyp.Kind() != reflect.String {
		return reflect.Value{}, false
	}

	kv := reflect.ValueOf(key)
	if keyTyp != kv.Type() {
		kv = kv.Convert(keyTyp)
	}
	return kv, true
}
