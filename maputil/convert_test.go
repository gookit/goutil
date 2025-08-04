package maputil_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestKeyToLower(t *testing.T) {
	src := map[string]string{"A": "v0"}
	ret := maputil.KeyToLower(src)

	assert.Contains(t, ret, "a")
	assert.NotContains(t, ret, "A")

	// Test with multiple keys
	src2 := map[string]string{"Name": "test", "Age": "20", "Email": "test@example.com"}
	ret2 := maputil.KeyToLower(src2)
	assert.Contains(t, ret2, "name")
	assert.Contains(t, ret2, "age")
	assert.Contains(t, ret2, "email")
	assert.Eq(t, "test", ret2["name"])
	assert.Eq(t, "20", ret2["age"])
	assert.Eq(t, "test@example.com", ret2["email"])

	// Test with empty map
	src3 := map[string]string{}
	ret3 := maputil.KeyToLower(src3)
	assert.Len(t, ret3, 0)
}

func TestToStringMap(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}
	ret := maputil.ToStringMap(src)

	assert.Eq(t, ret["a"], "v0")
	assert.Eq(t, ret["b"], "23")

	keys := []string{"key0", "key1"}
	mp := maputil.CombineToSMap(keys, []string{"val0", "val1"})
	assert.Len(t, mp, 2)
	assert.Eq(t, "val0", mp.Str("key0"))

	// Test with different data types
	src2 := map[string]any{
		"name":     "John",
		"age":      30,
		"isActive": true,
		"score":    95.5,
		"nilValue": nil,
	}
	ret2 := maputil.ToStringMap(src2)
	assert.Eq(t, "John", ret2["name"])
	assert.Eq(t, "30", ret2["age"])
	assert.Eq(t, "true", ret2["isActive"])
	assert.Eq(t, "95.5", ret2["score"])
	assert.Eq(t, "", ret2["nilValue"])

	// Test with empty map
	src3 := map[string]any{}
	ret3 := maputil.ToStringMap(src3)
	assert.Len(t, ret3, 0)
}

func TestToL2StrMap(t *testing.T) {
	l2smp := maputil.ToL2StrMap(map[string]any{
		"g1": map[string]any{
			"a": "v0", "b": 23,
		},
		"g2": map[string]string{
			"key0": "val0",
			"key1": "val1",
		},
	})

	assert.Eq(t, "v0", l2smp["g1"]["a"])
	assert.Eq(t, "val0", l2smp["g2"]["key0"])

	// Test with mixed data types
	l2smp2 := maputil.ToL2StringMap(map[string]any{
		"user": map[string]any{
			"name": "John",
			"age":  30,
		},
		"config": map[string]string{
			"debug": "true",
			"port":  "8080",
		},
		"empty": map[string]any{},
	})

	assert.Eq(t, "John", l2smp2["user"]["name"])
	assert.Eq(t, "30", l2smp2["user"]["age"])
	assert.Eq(t, "true", l2smp2["config"]["debug"])
	assert.Eq(t, "8080", l2smp2["config"]["port"])
	assert.Len(t, l2smp2["empty"], 0)

	// Test with empty map
	l2smp3 := maputil.ToL2StringMap(map[string]any{})
	assert.Len(t, l2smp3, 0)
}

func TestToAnyMap(t *testing.T) {
	src := map[string]string{"a": "v0", "b": "23"}

	mp := maputil.ToAnyMap(src)
	assert.Len(t, mp, 2)
	assert.Eq(t, "v0", mp["a"])

	src1 := map[string]any{"a": "v0", "b": "23"}
	mp = maputil.ToAnyMap(src1)
	assert.Len(t, mp, 2)
	assert.Eq(t, "v0", mp["a"])

	_, err := maputil.TryAnyMap(123)
	assert.Err(t, err)

	// Test with different map types
	src2 := map[int]string{1: "one", 2: "two"}
	mp2 := maputil.ToAnyMap(src2)
	assert.Len(t, mp2, 2)
	assert.Eq(t, "one", mp2["1"])
	assert.Eq(t, "two", mp2["2"])

	// Test with empty map
	src3 := map[string]string{}
	mp3 := maputil.ToAnyMap(src3)
	assert.Len(t, mp3, 0)

	// Test TryAnyMap with valid maps
	mp4, err := maputil.TryAnyMap(src)
	assert.NoErr(t, err)
	assert.Len(t, mp4, 2)
	assert.Eq(t, "v0", mp4["a"])

	mp5, err := maputil.TryAnyMap(src1)
	assert.NoErr(t, err)
	assert.Len(t, mp5, 2)
	assert.Eq(t, "v0", mp5["a"])
}

func TestHTTPQueryString(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}
	str := maputil.HTTPQueryString(src)

	fmt.Println(str)
	assert.Contains(t, str, "b=23")
	assert.Contains(t, str, "a=v0")

	// Test with different data types
	src2 := map[string]any{
		"name":  "John Doe",
		"age":   30,
		"active": true,
		"score": 95.5,
	}
	str2 := maputil.HTTPQueryString(src2)
	assert.Contains(t, str2, "name=John Doe")
	assert.Contains(t, str2, "age=30")
	assert.Contains(t, str2, "active=true")
	assert.Contains(t, str2, "score=95.5")

	// Test with empty map
	src3 := map[string]any{}
	str3 := maputil.HTTPQueryString(src3)
	assert.Eq(t, "", str3)

	// Test with special characters
	src4 := map[string]any{
		"email": "test@example.com",
		"msg":   "hello world",
	}
	str4 := maputil.HTTPQueryString(src4)
	assert.Contains(t, str4, "email=test@example.com")
	assert.Contains(t, str4, "msg=hello world")
}

func TestToString2(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}

	s := maputil.ToString2(src)
	assert.Contains(t, s, "b:23")
	assert.Contains(t, s, "a:v0")
}

func TestToString(t *testing.T) {
	src := map[string]any{"a": "v0", "b": 23}

	s := maputil.ToString(src)
	dump.P(s)
	assert.Contains(t, s, "b:23")
	assert.Contains(t, s, "a:v0")

	s = maputil.ToString(nil)
	assert.Eq(t, "", s)

	s = maputil.ToString(map[string]any{})
	assert.Eq(t, "{}", s)

	s = maputil.ToString(map[string]any{"": nil})
	assert.Eq(t, "{:}", s)

	// Test with various data types
	src2 := map[string]any{
		"name":     "John",
		"age":      30,
		"isActive": true,
		"score":    95.5,
		"tags":     []string{"dev", "user"},
		"nilValue": nil,
		"emptyMap": map[string]any{},
	}
	s2 := maputil.ToString(src2)
	assert.Contains(t, s2, "name:John")
	assert.Contains(t, s2, "age:30")
	assert.Contains(t, s2, "isActive:true")
	assert.Contains(t, s2, "score:95.5")
	assert.Contains(t, s2, "tags:[dev user]")
	assert.Contains(t, s2, "nilValue:")
	assert.Contains(t, s2, "emptyMap:map[]")

	// Test with nested maps
	src3 := map[string]any{
		"user": map[string]any{
			"name": "John",
			"age":  30,
		},
	}
	s3 := maputil.ToString(src3)
	assert.Contains(t, s3, "user:map[name:John age:30]")
}

func TestFlatten(t *testing.T) {
	data := map[string]any{
		"name": "inhere",
		"age":  234,
		"top": map[string]any{
			"sub0": "val0",
			"sub1": []string{"val1-0", "val1-1"},
		},
	}

	mp := maputil.Flatten(data)
	assert.ContainsKeys(t, mp, []string{"age", "name", "top.sub0", "top.sub1[0]", "top.sub1[1]"})
	assert.Nil(t, maputil.Flatten(nil))

	fmp := make(map[string]string)
	maputil.FlatWithFunc(data, func(path string, val reflect.Value) {
		fmp[path] = fmt.Sprintf("%v", val.Interface())
	})
	dump.P(fmp)
	assert.Eq(t, "inhere", fmp["name"])
	assert.Eq(t, "234", fmp["age"])
	assert.Eq(t, "val0", fmp["top.sub0"])
	assert.Eq(t, "val1-0", fmp["top.sub1[0]"])

	assert.NotPanics(t, func() {
		maputil.FlatWithFunc(nil, nil)
	})

	// Test with more complex nested structure
	data2 := map[string]any{
		"user": map[string]any{
			"profile": map[string]any{
				"name": "John",
				"preferences": map[string]any{
					"theme": "dark",
					"lang":  "en",
				},
			},
			"settings": map[string]any{
				"notifications": true,
			},
		},
		"app": map[string]any{
			"version": "1.0.0",
		},
	}

	mp2 := maputil.Flatten(data2)
	assert.Eq(t, "John", mp2["user.profile.name"])
	assert.Eq(t, "dark", mp2["user.profile.preferences.theme"])
	assert.Eq(t, "en", mp2["user.profile.preferences.lang"])
	assert.Eq(t, true, mp2["user.settings.notifications"])
	assert.Eq(t, "1.0.0", mp2["app.version"])

	// Test with array values
	data3 := map[string]any{
		"tags": []string{"web", "api"},
		"numbers": []int{1, 2, 3},
	}
	mp3 := maputil.Flatten(data3)
	assert.Eq(t, "web", mp3["tags[0]"])
	assert.Eq(t, "api", mp3["tags[1]"])
	assert.Eq(t, 1, mp3["numbers[0]"])
	assert.Eq(t, 2, mp3["numbers[1]"])
	assert.Eq(t, 3, mp3["numbers[2]"])
}

func TestStringsMapToAnyMap(t *testing.T) {
	assert.Nil(t, maputil.StringsMapToAnyMap(nil))

	hh := http.Header{
		"key0": []string{"val0", "val1"},
		"key1": []string{"val2"},
	}

	mp := maputil.StringsMapToAnyMap(hh)
	assert.Contains(t, mp, "key0")
	assert.Contains(t, mp, "key1")
	assert.Len(t, mp["key0"], 2)

	dm := maputil.Data(mp)
	assert.Eq(t, "val0", dm.Str("key0.0"))
	assert.Eq(t, "val2", dm.Str("key1"))

	// Test with empty map
	emptyMap := map[string][]string{}
	result := maputil.StringsMapToAnyMap(emptyMap)
	assert.Nil(t, result)

	// Test with multiple single values
	singleValues := map[string][]string{
		"key1": {"value1"},
		"key2": {"value2"},
	}
	result2 := maputil.StringsMapToAnyMap(singleValues)
	assert.Eq(t, "value1", result2["key1"])
	assert.Eq(t, "value2", result2["key2"])

	// Test with mixed single and multiple values
	mixedValues := map[string][]string{
		"single": {"value"},
		"multiple": {"value1", "value2", "value3"},
	}
	result3 := maputil.StringsMapToAnyMap(mixedValues)
	assert.Eq(t, "value", result3["single"])
	assert.Len(t, result3["multiple"], 3)
}

func TestCombineToMap(t *testing.T) {
	keys := []string{"key0", "key1"}

	mp := maputil.CombineToMap(keys, []int{1, 2})
	assert.Len(t, mp, 2)
	assert.Eq(t, 1, mp["key0"])
	assert.Eq(t, 2, mp["key1"])

	// Test with different types
	strKeys := []string{"name", "city"}
	strValues := []string{"John", "Beijing"}
	strMap := maputil.CombineToMap(strKeys, strValues)
	assert.Len(t, strMap, 2)
	assert.Eq(t, "John", strMap["name"])
	assert.Eq(t, "Beijing", strMap["city"])

	// Test with mismatched lengths (values shorter)
	intKeys := []int{1, 2, 3}
	intValues := []string{"one", "two"}
	intMap := maputil.CombineToMap(intKeys, intValues)
	assert.Len(t, intMap, 3)
	assert.Eq(t, "one", intMap[1])
	assert.Eq(t, "two", intMap[2])
	assert.Eq(t, "", intMap[3]) // zero value

	// Test with mismatched lengths (keys shorter)
	floatKeys := []float64{1.1, 2.2}
	floatValues := []bool{true, false, true}
	floatMap := maputil.CombineToMap(floatKeys, floatValues)
	assert.Len(t, floatMap, 2)
	assert.Eq(t, true, floatMap[1.1])
	assert.Eq(t, false, floatMap[2.2])

	// Test with empty slices
	emptyKeys := []string{}
	emptyValues := []int{}
	emptyMap := maputil.CombineToMap(emptyKeys, emptyValues)
	assert.Len(t, emptyMap, 0)
}

func TestCombineToSMap(t *testing.T) {
	// Basic test case from TestToStringMap
	keys := []string{"key0", "key1"}
	mp := maputil.CombineToSMap(keys, []string{"val0", "val1"})
	assert.Len(t, mp, 2)
	assert.Eq(t, "val0", mp.Str("key0"))

	// Test with different lengths
	keys2 := []string{"a", "b", "c"}
	values2 := []string{"1", "2"}
	mp2 := maputil.CombineToSMap(keys2, values2)
	assert.Len(t, mp2, 3)
	assert.Eq(t, "1", mp2.Str("a"))
	assert.Eq(t, "2", mp2.Str("b"))
	assert.Eq(t, "", mp2.Str("c")) // empty string for missing value

	// Test with more values than keys
	keys3 := []string{"x"}
	values3 := []string{"1", "2", "3"}
	mp3 := maputil.CombineToSMap(keys3, values3)
	assert.Len(t, mp3, 1)
	assert.Eq(t, "1", mp3.Str("x"))

	// Test with empty slices
	keys4 := []string{}
	values4 := []string{}
	mp4 := maputil.CombineToSMap(keys4, values4)
	assert.Len(t, mp4, 0)

	// Test with nil slices
	mp5 := maputil.CombineToSMap(nil, nil)
	assert.Len(t, mp5, 0)
}

// AnyToStrMap tests
func TestAnyToStrMap(t *testing.T) {
	// Test with nil input
	assert.Nil(t, maputil.AnyToStrMap(nil))

	// Test with map[string]string input
	src1 := map[string]string{"key1": "val1", "key2": "val2"}
	result1 := maputil.AnyToStrMap(src1)
	assert.Eq(t, src1, result1)

	// Test with map[string]any input
	src2 := map[string]any{"key1": "val1", "key2": 123, "key3": true}
	result2 := maputil.AnyToStrMap(src2)
	assert.Eq(t, "val1", result2["key1"])
	assert.Eq(t, "123", result2["key2"])
	assert.Eq(t, "true", result2["key3"])

	// Test with unsupported type
	var src3 interface{} = "not a map"
	result3 := maputil.AnyToStrMap(src3)
	assert.Nil(t, result3)
}

func TestSliceToSMap(t *testing.T) {
	// Test with valid key-value pairs
	result := maputil.SliceToSMap("k1", "v1", "k2", "v2")
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Eq(t, "v1", result["k1"])
	assert.Eq(t, "v2", result["k2"])

	// Test with empty input
	assert.Nil(t, maputil.SliceToSMap())

	// Test with odd number of arguments (should return nil)
	nilResult := maputil.SliceToSMap("k1", "v1", "k2")
	assert.Nil(t, nilResult)

	// Test with a single pair
	singleResult := maputil.SliceToSMap("key", "value")
	assert.NotNil(t, singleResult)
	assert.Len(t, singleResult, 1)
	assert.Eq(t, "value", singleResult["key"])
}

func TestSliceToMap(t *testing.T) {
	// Test with valid key-value pairs of different types
	result := maputil.SliceToMap("name", "John", "age", 30, "active", true)
	assert.NotNil(t, result)
	assert.Len(t, result, 3)
	assert.Eq(t, "John", result["name"])
	assert.Eq(t, 30, result["age"])
	assert.Eq(t, true, result["active"])

	// Test with empty input
	assert.Nil(t, maputil.SliceToMap())

	// Test with odd number of arguments (should return nil)
	nilResult := maputil.SliceToMap("k1", "v1", "k2")
	assert.Nil(t, nilResult)

	// Test with mixed types including slice and map
	mixedResult := maputil.SliceToMap("string", "value", "int", 42, "slice", []int{1, 2, 3}, "map", map[string]string{"k": "v"})
	assert.NotNil(t, mixedResult)
	assert.Len(t, mixedResult, 4)
	assert.Eq(t, "value", mixedResult["string"])
	assert.Eq(t, 42, mixedResult["int"])
	assert.Eq(t, []int{1, 2, 3}, mixedResult["slice"])
	assert.Eq(t, map[string]string{"k": "v"}, mixedResult["map"])
}

func TestSliceToTypeMap(t *testing.T) {
	// Test with valid key-value pairs and a conversion function
	result := maputil.SliceToTypeMap(
		func(val any) string {
			return fmt.Sprintf("%v", val)
		},
		"name", "John", "age", 30, "active", true)

	assert.NotNil(t, result)
	assert.Len(t, result, 3)
	assert.Eq(t, "John", result["name"])
	assert.Eq(t, "30", result["age"])
	assert.Eq(t, "true", result["active"])

	// Test with empty input
	nilResult1 := maputil.SliceToTypeMap(func(val any) string {
		return fmt.Sprintf("%v", val)
	})
	assert.Nil(t, nilResult1)

	// Test with odd number of arguments (should return nil)
	nilResult := maputil.SliceToTypeMap(func(val any) string {
		return fmt.Sprintf("%v", val)
	}, "k1", "v1", "k2")
	assert.Nil(t, nilResult)

	// Test with int conversion function
	intResult := maputil.SliceToTypeMap(
		func(val any) int {
			if v, ok := val.(int); ok {
				return v
			}
			return 0
		},
		"first", 10, "second", 20, "third", 30)

	assert.NotNil(t, intResult)
	assert.Len(t, intResult, 3)
	assert.Eq(t, 10, intResult["first"])
	assert.Eq(t, 20, intResult["second"])
	assert.Eq(t, 30, intResult["third"])
}
