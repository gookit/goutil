package jsonutil_test

import (
	"testing"

	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/testutil/assert"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var testUser = user{"inhere", 200}
var invalid = jsonutil.Encode

func TestPretty(t *testing.T) {
	tests := []any{
		map[string]int{"a": 1},
		struct {
			A int `json:"a"`
		}{1},
	}
	want := `{
    "a": 1
}`
	for _, sample := range tests {
		got, err := jsonutil.Pretty(sample)
		assert.NoErr(t, err)
		assert.Eq(t, want, got)
	}

	bts, err := jsonutil.EncodePretty(map[string]int{"a": 1})
	assert.NoErr(t, err)
	assert.Eq(t, want, string(bts))
}

func TestMustPretty(t *testing.T) {
	type T1 struct {
		A1 int    `json:"a1"`
		B1 string `json:"b1"`
	}

	type T2 struct {
		T1        // embedded without json tag
		C  T1     `json:"c"` // as field and with json tag
		A  int    `json:"a"`
		B  string `json:"b"`
	}

	t1 := T1{A1: 1, B1: "b1-at-T1"}
	t2 := T2{T1: t1, A: 1, B: "b-at-T2", C: t1}

	str := jsonutil.MustPretty(t2)
	assert.StrContains(t, str, `"c": {`)
	assert.NotContains(t, str, `"T1": {`)
	// fmt.Println(str)
	/*Output:
	{
	    "a1": 1,
	    "b1": "b1-at-T1",
	    "c": {
	        "a1": 1,
	        "b1": "b1-at-T1"
	    },
	    "a": 1,
	    "b": "b-at-T2"
	}
	*/

	type T3 struct {
		T1 `json:"t1"` // with json tag
		A  int         `json:"a"`
		B  string      `json:"b"`
	}

	t3 := T3{T1: t1, A: 1, B: "b-at-T3"}
	str = jsonutil.MustPretty(t3)
	assert.StrContains(t, str, `"t1": {`)
	// fmt.Println(str)
	/*Output:
	{
	    "t1": {
	        "a1": 1,
	        "b1": "b1-at-T1"
	    },
	    "a": 1,
	    "b": "b-at-T3"
	}
	*/

	assert.Panics(t, func() {
		jsonutil.MustPretty(jsonutil.Encode)
	})
}

func TestMapping(t *testing.T) {
	mp := map[string]any{"name": "inhere", "age": 200}

	usr := &user{}
	err := jsonutil.Mapping(mp, usr)
	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)

	assert.Err(t, jsonutil.Mapping(invalid, usr))
}

func TestWriteReadFile(t *testing.T) {
	usr := user{"inhere", 200}

	err := jsonutil.WriteFile("testdata/test.json", &usr)
	assert.NoErr(t, err)
	assert.Err(t, jsonutil.WriteFile("testdata/test.json", jsonutil.Encode))

	err = jsonutil.WritePretty("testdata/test2.json", &usr)
	assert.NoErr(t, err)
	assert.Err(t, jsonutil.WritePretty("testdata/test2.json", jsonutil.Encode))

	err = jsonutil.ReadFile("testdata/test.json", &usr)
	assert.NoErr(t, err)

	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)

	err = jsonutil.ReadFile("testdata/not-exist.json", &usr)
	assert.Err(t, err)
}

func TestIsJsonFast(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"single character", "a", false},
		{"two characters object", "{}", true},
		{"two characters slice", "[]", true},
		{"invalid json", "{a}", false},
		{"valid json object", `{"a": 1}`, true},
		{"valid json array", `[1]`, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Eq(t, tc.expected, jsonutil.IsJSON(tc.input))

			if jsonutil.IsJSONFast(tc.input) != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, !tc.expected)
			}
		})
	}

	// test IsArray
	t.Run("IsArray", func(t *testing.T) {
		assert.True(t, jsonutil.IsArray(`[]`))
		assert.True(t, jsonutil.IsArray(`[1]`))
		assert.False(t, jsonutil.IsArray(`{"a": 1}`))
		assert.False(t, jsonutil.IsArray(`a`))
	})

	// test IsObject
	t.Run("IsObject", func(t *testing.T) {
		assert.True(t, jsonutil.IsObject(`{}`))
		assert.True(t, jsonutil.IsObject(`{"a": 1}`))
		assert.False(t, jsonutil.IsObject(`[1]`))
		assert.False(t, jsonutil.IsObject(``))
		assert.False(t, jsonutil.IsObject(`a`))
	})
}

func TestStripComments(t *testing.T) {
	is := assert.New(t)

	str := jsonutil.StripComments(`{"name":"app"}`)
	is.Eq(`{"name":"app"}`, str)

	givens := []string{
		// single line comments
		`{
"name":"app" // comments
}`,
		`{
// comments
"name":"app" 
}`,
		`{"name":"app"} // comments
`,
		// multi line comments
		`{"name":"app"} /* comments */
`,
		`/* comments */
{"name":"app"}`,
		`/* 
comments 
*/
{"name":"app"}`,
		`/** 
comments 
*/
{"name":"app"}`,
		`/** 
comments 
**/
{"name":"app"}`,
		`/** 
* comments 
**/
{"name":"app"}`,
		`/** 
/* comments 
**/
{"name":"app"}`,
		`/** 
/* comments *
**/
{"name":"app"}`,
		`{"name": /*comments*/"app"}`,
		`{/*comments*/"name": "app"}`,
	}
	wants := []string{
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		// multi line comments
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name":"app"}`,
		`{"name": "app"}`,
		`{"name": "app"}`,
	}

	for i, s := range givens {
		is.Eq(wants[i], jsonutil.StripComments(s))
	}

	str = jsonutil.StripComments(`{"name":"app"} // comments`)
	is.Eq(`{"name":"app"}`, str)

	// fix https://github.com/gookit/config/issues/2
	str = jsonutil.StripComments(`{"name":"http://abc.com"} // comments`)
	is.Eq(`{"name":"http://abc.com"}`, str)

	str = jsonutil.StripComments(`{
"address": [
	"http://192.168.1.XXX:2379"
]
} // comments`)
	is.Eq(`{"address":["http://192.168.1.XXX:2379"]}`, str)

	s := `{"name":"http://abc.com"} // comments`
	s = jsonutil.StripComments(s)
	assert.Eq(t, `{"name":"http://abc.com"}`, s)

	s = `
{// comments
    "name": "app", // comments
/*comments*/
    "debug": false,
    "baseKey": "value", // comments
	/* comments */
    "age": 123,
    "envKey1": "${NotExist|defValue}",
    "map1": { // comments
        "key": "val",
        "key1": "val1",
        "key2": "val2"
    },
    "arr1": [ // comments
        "val",
        "val1", // comments
		/* comments */
        "val2",
		"http://a.com"
    ],
	/* 
		comments 
*/
    "lang": {
		/** 
 		 * comments 
 		 */
        "dir": "res/lang",
        "allowed": {
            "en": "val",
            "zh-CN": "val2"
        }
    }
}`
	ep := `{"name":"app","debug":false,"baseKey":"value","age":123,"envKey1":"${NotExist|defValue}","map1":{"key":"val","key1":"val1","key2":"val2"},"arr1":["val","val1","val2","http://a.com"],"lang":{"dir":"res/lang","allowed":{"en":"val","zh-CN":"val2"}}}`
	s = jsonutil.StripComments(s)
	assert.Eq(t, ep, s)
}
