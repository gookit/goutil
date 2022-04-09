package jsonutil_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/jsonutil"
	"github.com/stretchr/testify/assert"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var testUser = user{"inhere", 200}

func TestPretty(t *testing.T) {
	tests := []interface{}{
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
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	}
}

func TestEncode(t *testing.T) {
	bts, err := jsonutil.Encode(testUser)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"inhere","age":200}`, string(bts))

	bts, err = jsonutil.Encode(&testUser)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"inhere","age":200}`, string(bts))
}

func TestEncodeUnescapeHTML(t *testing.T) {
	bts, err := jsonutil.Encode(&testUser)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"inhere","age":200}`, string(bts))
}

func TestEncodeToWriter(t *testing.T) {
	buf := &bytes.Buffer{}

	err := jsonutil.EncodeToWriter(testUser, buf)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"inhere","age":200}
`, buf.String())
}

func TestDecode(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.Decode([]byte(str), usr)

	assert.NoError(t, err)
	assert.Equal(t, "inhere", usr.Name)
	assert.Equal(t, 200, usr.Age)
}

func TestDecodeString(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.DecodeString(str, usr)

	assert.NoError(t, err)
	assert.Equal(t, "inhere", usr.Name)
	assert.Equal(t, 200, usr.Age)
}

func TestWriteReadFile(t *testing.T) {
	user := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"inhere", 200}

	err := jsonutil.WriteFile("testdata/test.json", &user)
	assert.NoError(t, err)

	err = jsonutil.ReadFile("testdata/test.json", &user)
	assert.NoError(t, err)

	assert.Equal(t, "inhere", user.Name)
	assert.Equal(t, 200, user.Age)
}

func TestStripComments(t *testing.T) {
	is := assert.New(t)

	str := jsonutil.StripComments(`{"name":"app"}`)
	is.Equal(`{"name":"app"}`, str)

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
		is.Equal(wants[i], jsonutil.StripComments(s))
	}

	str = jsonutil.StripComments(`{"name":"app"} // comments`)
	is.Equal(`{"name":"app"}`, str)

	// fix https://github.com/gookit/config/issues/2
	str = jsonutil.StripComments(`{"name":"http://abc.com"} // comments`)
	is.Equal(`{"name":"http://abc.com"}`, str)

	str = jsonutil.StripComments(`{
"address": [
	"http://192.168.1.XXX:2379"
]
} // comments`)
	is.Equal(`{"address":["http://192.168.1.XXX:2379"]}`, str)

	s := `{"name":"http://abc.com"} // comments`
	s = jsonutil.StripComments(s)
	assert.Equal(t, `{"name":"http://abc.com"}`, s)

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
	assert.Equal(t, ep, s)
}
