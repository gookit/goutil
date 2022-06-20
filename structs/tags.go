package structs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
)

var errNotAnStruct = errors.New("must input an struct")

// TagParser struct
type TagParser struct {
	TagNames string

	Func func(tagVal string) map[string]string
}

// ParseTags for parse struct tags.
func ParseTags(v interface{}, tagNames []string) (map[string]maputil.SMap, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		rv = rv.Elem()
	}

	return ParseReflectTags(rv.Type(), tagNames)
}

// ParseReflectTags parse struct tags info.
func ParseReflectTags(rt reflect.Type, tagNames []string) (map[string]maputil.SMap, error) {
	if rt.Kind() != reflect.Struct {
		return nil, errNotAnStruct
	}

	// key is field name.
	result := make(map[string]maputil.SMap)

	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)

		// skip don't exported field
		name := sf.Name
		if name[0] >= 'a' && name[0] <= 'z' {
			continue
		}

		smp := make(maputil.SMap)
		for _, tagName := range tagNames {
			// eg: "name=int0;shorts=i;required=true;desc=int option message"
			tagVal := sf.Tag.Get(tagName)
			if tagVal == "" {
				continue
			}

			smp[tagName] = tagVal
		}

		result[name] = smp

		// TODO field is struct.
		// fv := v.Field(i)
		// ft := t.Field(i).Type
		// if ft.Kind() == reflect.Ptr {
		// 	// isPtr = true
		// 	ft = ft.Elem()
		// 	// fv = fv.Elem()
		// }
	}
	return result, nil
}

// ParseTagValueINI tag value string. is like INI format data
//
// eg: "name=int0;shorts=i;required=true;desc=int option message"
func ParseTagValueINI(field, tagStr string) (mp maputil.SMap, err error) {
	tagStr = strings.Trim(tagStr, "; ")
	ss := strutil.Split(tagStr, ";")
	if len(ss) == 0 {
		return
	}

	mp = make(maputil.SMap, len(ss))
	for _, s := range ss {
		if !strings.ContainsRune(s, '=') {
			err = fmt.Errorf("parse tag error on field '%s': item must match `KEY=VAL`", field)
			return
		}

		kvNodes := strings.SplitN(s, "=", 2)
		key, val := kvNodes[0], strings.TrimSpace(kvNodes[1])
		// if !flagTagKeys.Has(key) {
		// 	panicf("parse tag error on field '%s': invalid key name '%s'", name, key)
		// }

		mp[key] = val
	}
	return
}
