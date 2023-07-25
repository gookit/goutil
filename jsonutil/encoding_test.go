package jsonutil_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/gookit/goutil/jsonutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestEncode(t *testing.T) {
	bts, err := jsonutil.Encode(testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))

	bts, err = jsonutil.Encode(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))

	assert.Eq(t, `{"name":"inhere","age":200}`, jsonutil.MustString(testUser))

	assert.Panics(t, func() {
		jsonutil.MustString(invalid)
	})

	bts, err = jsonutil.Encode(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}`, string(bts))
}

func TestEncodeUnescapeHTML(t *testing.T) {
	bts, err := jsonutil.EncodeUnescapeHTML(&testUser)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}
`, string(bts))

	_, err = jsonutil.EncodeUnescapeHTML(invalid)
	assert.Err(t, err)
}

func TestEncodeToWriter(t *testing.T) {
	buf := &bytes.Buffer{}

	err := jsonutil.EncodeToWriter(testUser, buf)
	assert.NoErr(t, err)
	assert.Eq(t, `{"name":"inhere","age":200}
`, buf.String())
}

func TestDecode(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.Decode([]byte(str), usr)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)
}

func TestDecodeString(t *testing.T) {
	str := `{"name":"inhere","age":200}`
	usr := &user{}
	err := jsonutil.DecodeString(str, usr)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)

	// DecodeReader
	usr = &user{}
	err = jsonutil.DecodeReader(strings.NewReader(str), usr)

	assert.NoErr(t, err)
	assert.Eq(t, "inhere", usr.Name)
	assert.Eq(t, 200, usr.Age)
}
