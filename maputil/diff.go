package maputil

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

// Diff get a map record the difference between the two maps.
// The calculation is based on a map change to b map
// if deleted some field in a map then the field will be returned in result field same key but is nil value.
// a map {"a":"a","b":"b"}
// b map {"b":"b","c":"c"}
// return map {"a":nil,"c":"c"}
func Diff(a, b map[string]interface{}) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	memo := make(map[string]struct{}, len(a))

	for k := range a {
		memo[k] = struct{}{}
	}

	for k, vb := range b {
		va, ok := a[k]
		if !ok {
			res[k] = vb
			continue
		}

		delete(memo, k)
		rbv := reflect.ValueOf(vb)
		rav := reflect.ValueOf(va)

	reswitch:
		if !rbv.IsValid() && !rav.IsValid() {
			continue
		}

		if (!rbv.IsValid() && rav.IsValid()) ||
			(rbv.IsValid() && !rav.IsValid()) {
			res[k] = vb
			continue
		}

		if rbv.Kind() != rav.Kind() {
			res[k] = vb
			continue
		}

		jsonifya, err := json.Marshal(va)
		if err != nil {
			return res, errors.Wrap(err, "json.Marshal")
		}
		jsonifyb, err := json.Marshal(vb)
		if err != nil {
			return res, errors.Wrap(err, "json.Marshal")
		}

		switch rbv.Kind() {
		case reflect.Slice, reflect.Array:
			var (
				tmpa []interface{}
				tmpb []interface{}
			)
			if err := json.Unmarshal(jsonifya, &tmpa); err != nil {
				return nil, errors.Wrap(err, "json unmarshal error")
			}
			if err := json.Unmarshal(jsonifyb, &tmpb); err != nil {
				return nil, errors.Wrap(err, "json unmarshal error")
			}

			result, err := compareArr(tmpa, tmpb)
			if err != nil {
				return res, errors.WithMessage(err, "compareArr error")
			}
			// if len(result) == 0 && result != nil  means no changes in this field.
			if len(result) != 0 || result == nil {
				res[k] = result
			}

		case reflect.Map, reflect.Struct:
			tmpa := make(map[string]interface{})
			tmpb := make(map[string]interface{})
			if err := json.Unmarshal(jsonifya, &tmpa); err != nil {
				return nil, errors.Wrap(err, "json unmarshal error")
			}
			if err := json.Unmarshal(jsonifyb, &tmpb); err != nil {
				return nil, errors.Wrap(err, "json unmarshal error")
			}

			if val, err := Diff(tmpa, tmpb); err != nil {
				return nil, errors.WithMessage(err, "map diff error")
			} else if len(val) > 0 {
				res[k] = val
			}
		case reflect.Ptr:
			rbv = rbv.Elem()
			rav = rav.Elem()
			goto reswitch

		default:
			if !rav.IsValid() && !rbv.IsValid() {
				continue
			}
			isComparable := rbv.Type().Comparable()
			if !isComparable {
				return nil, errors.New("not comparable")
			}
			if va != vb {
				res[k] = vb
			}

		}
	}
	// record the deleted key and get it null value
	for k := range memo {
		res[k] = nil
	}
	return res, nil
}

func compareArr(a, b []interface{}) ([]interface{}, error) {
	res := make([]interface{}, 0)
	if a == nil && b == nil {
		return res, nil
	}
	for i, v := range b {
		if i >= len(a) {
			res = append(res, v)
			continue
		}

		if reflect.TypeOf(v).Kind() == reflect.Map &&
			reflect.TypeOf(a[i]).Kind() == reflect.Map {
			amap, ok := a[i].(map[string]interface{})
			if !ok {
				continue
			}
			bmap, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			if val, err := Diff(amap, bmap); err != nil {
				return res, nil
			} else if len(val) > 0 {
				res = append(res, val)
				continue
			}
			continue
		}

		if reflect.TypeOf(v).Kind() == reflect.Struct &&
			reflect.TypeOf(a[i]).Kind() == reflect.Struct {
			var (
				amap map[string]interface{}
				bmap map[string]interface{}
			)
			c, err := json.Marshal(v)
			if err != nil {
				return res, errors.Wrap(err, "json.Marshal")
			}
			err = json.Unmarshal(c, &bmap)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal")
			}
			c, _ = json.Marshal(a[i])
			err = json.Unmarshal(c, &amap)
			if err != nil {
				return nil, errors.Wrap(err, "json.Unmarshal")
			}
			if val, err := Diff(amap, bmap); err != nil {
				return res, errors.Wrap(err, "map diff error")
			} else if len(val) > 0 {
				res = append(res, val)
				continue
			}
		}
		res = append(res, v)
	}

	return res, nil
}
