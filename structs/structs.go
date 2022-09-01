package structs

import (
	"errors"
	"reflect"

	"github.com/gookit/goutil/envutil"
	"github.com/gookit/goutil/reflects"
)

// MapStruct simple copy src struct value to dst struct
// func MapStruct(srcSt, dstSt interface{}) {
// 	// TODO
// }

const defaultInitTag = "default"

// InitOptFunc define
type InitOptFunc func(opt *InitOptions)

// InitOptions struct
type InitOptions struct {
	TagName  string
	ParseEnv bool
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
//	err = structs.InitDefaults(u1)
//	fmt.Printf("%+v\n", u1)
//	// Output: {Name:inhere Age:30}
func InitDefaults(ptr interface{}, optFns ...InitOptFunc) error {
	rv := reflect.ValueOf(ptr)
	if rv.Kind() != reflect.Ptr {
		return errors.New("must be provider an pointer value")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New("must be provider an struct value")
	}

	opt := &InitOptions{TagName: defaultInitTag}
	for _, fn := range optFns {
		fn(opt)
	}

	return initDefaults(rv, opt)
}

func initDefaults(rv reflect.Value, opt *InitOptions) error {
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		ft := rt.Field(i)
		// skip don't exported field
		if ft.Name[0] >= 'a' && ft.Name[0] <= 'z' {
			continue
		}

		fv := rv.Field(i)
		if fv.Kind() == reflect.Struct {
			if err := initDefaults(fv, opt); err != nil {
				return err
			}
			continue
		}

		// skip set on field has value
		if !fv.IsZero() {
			continue
		}

		tagVal, ok := ft.Tag.Lookup(opt.TagName)
		if !ok || tagVal == "" || !fv.CanSet() {
			continue
		}

		// get real type of the ptr value
		if fv.Kind() == reflect.Ptr {
			elemTyp := fv.Type().Elem()
			fv.Set(reflect.New(elemTyp))
			// use elem for set value
			fv = reflect.Indirect(fv)
		}

		if opt.ParseEnv {
			tagVal = envutil.ParseValue(tagVal)
		}

		val, err := reflects.ValueByKind(tagVal, fv.Kind())
		if err != nil {
			return err
		}

		fv.Set(val)
	}

	return nil
}
