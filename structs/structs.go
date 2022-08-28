package structs

import (
	"errors"
	"reflect"

	"github.com/gookit/goutil/reflects"
)

// MapStruct simple copy src struct value to dst struct
// func MapStruct(srcSt, dstSt interface{}) {
// 	// TODO
// }

const defaultInitTag = "default"

// InitOptions struct
type InitOptions struct {
	TagName string
}

// InitDefaults init struct default value by field "default" tag.
//
// Example:
//
//	type User1 struct {
//		Name string `default:"inhere"`
//		Age  int32  `default:"30"`
//	}
//
//	u1 := &User1{}
//	err = structs.InitDefaults(u1, nil)
//	fmt.Printf("%+v\n", u1)
//	// Output: {Name:inhere Age:30}
func InitDefaults(ptr interface{}, opt *InitOptions) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		return errors.New("must be provider an pointer value")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("must be provider an struct value")
	}

	if opt == nil {
		opt = &InitOptions{TagName: defaultInitTag}
	} else if opt.TagName == "" {
		opt.TagName = defaultInitTag
	}

	return initDefaults(rv, opt.TagName)
}

func initDefaults(rv reflect.Value, tagName string) error {
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		// skip don't exported field
		if ft.Name[0] >= 'a' && ft.Name[0] <= 'z' {
			continue
		}

		fv := rv.Field(i)
		if fv.Kind() == reflect.Struct {
			err := initDefaults(fv, tagName)
			if err != nil {
				return err
			}
			continue
		}

		tagVal, ok := ft.Tag.Lookup(tagName)
		if ok && tagVal != "" && fv.CanSet() {
			val, err := reflects.ValueByKind(tagVal, fv.Kind())
			if err != nil {
				return err
			}

			fv.Set(val)
		}
	}

	return nil
}
