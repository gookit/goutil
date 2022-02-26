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

// ParseTags TODO for parse struct tags.
func ParseTags(v interface{}) error {
	rv := reflect.ValueOf(v)
	return ParseReflectTags(rv)
}

// ParseReflectTags value
func ParseReflectTags(v reflect.Value) error {
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	t := v.Type()
	if t.Kind() != reflect.Struct {
		return errNotAnStruct
	}

	tagName := "xxx"
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		ft := t.Field(i).Type

		// skip don't exported field
		name := sf.Name
		if name[0] >= 'a' && name[0] <= 'z' {
			continue
		}

		// eg: "name=int0;shorts=i;required=true;desc=int option message"
		str := sf.Tag.Get(tagName)
		if str == "" {
			continue
		}

		fv := v.Field(i)
		if ft.Kind() == reflect.Ptr {
			// isPtr = true
			ft = ft.Elem()
			// fv = fv.Elem()
		}

		fmt.Println(fv.String())
	}
	return nil
}

// ParseTagValue string.
func ParseTagValue(str string) maputil.SMap {
	return make(maputil.SMap) // TODO
}

// ParseTagValueINI tag value string. is like INI format data
// eg: "name=int0;shorts=i;required=true;desc=int option message"
func ParseTagValueINI(field, str string) (mp maputil.SMap, err error) {
	str = strings.Trim(str, "; ")
	ss := strutil.Split(str, ";")
	if len(ss) == 0 {
		return
	}

	mp = make(maputil.SMap, len(ss))
	for _, s := range ss {
		if strings.ContainsRune(s, '=') == false {
			err = fmt.Errorf("parse tag error on field '%s': item must match `KEY=VAL`", field)
			return
		}

		kvNodes := strings.SplitN(s, "=", 2)
		key, val := kvNodes[0], kvNodes[1]
		// if !flagTagKeys.Has(key) {
		// 	panicf("parse tag error on field '%s': invalid key name '%s'", name, key)
		// }

		mp[key] = val
	}
	return
}
