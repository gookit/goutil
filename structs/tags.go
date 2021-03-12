package structs

import (
	"errors"
	"fmt"
	"reflect"
)

var errNotAnStruct = errors.New("must input an struct")

// TagParser struct
type TagParser struct {
	TagName string

	Func func(tagVal string) map[string]string
}

// TODO for parse struct tags.
func ParseTags(v interface{}) error {
	rv := reflect.ValueOf(v)
	return ParseReflectTags(rv)
}

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
