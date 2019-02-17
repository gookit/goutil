package jsonutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripComments(t *testing.T) {
	is := assert.New(t)

	str := StripComments(`{"name":"app"}`)
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
		is.Equal(wants[i], StripComments(s))
	}

	str = StripComments(`{"name":"app"} // comments`)
	is.Equal(`{"name":"app"}`, str)

	// fix https://github.com/gookit/config/issues/2
	str = StripComments(`{"name":"http://abc.com"} // comments`)
	is.Equal(`{"name":"http://abc.com"}`, str)

	str = StripComments(`{
"address": [
	"http://192.168.1.XXX:2379"
]
} // comments`)
	is.Equal(`{"address":["http://192.168.1.XXX:2379"]}`, str)

	s := `{"name":"http://abc.com"} // comments`
	s = StripComments(s)
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
	s = StripComments(s)
	assert.Equal(
		t,
		`{"name":"app","debug":false,"baseKey":"value","age":123,"envKey1":"${NotExist|defValue}","map1":{"key":"val","key1":"val1","key2":"val2"},"arr1":["val","val1","val2","http://a.com"],"lang":{"dir":"res/lang","allowed":{"en":"val","zh-CN":"val2"}}}`,
		s,
	)
}
