package byteutil_test

import (
	"errors"
	"io/fs"
	"testing"
	"time"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestStrOrErr(t *testing.T) {
	bs := []byte("hi, inhere")
	assert.Eq(t, "hi, inhere", byteutil.SafeString(bs, nil))
	assert.Eq(t, "", byteutil.SafeString(bs, errors.New("error")))

	str, err := byteutil.StrOrErr(bs, nil)
	assert.NoErr(t, err)
	assert.Eq(t, "hi, inhere", str)

	str, err = byteutil.StrOrErr(bs, errors.New("error"))
	assert.Err(t, err)
	assert.Eq(t, "", str)
}

func TestToString(t *testing.T) {
	assert.Eq(t, "123", byteutil.String([]byte("123")))
	assert.Eq(t, "123", byteutil.ToString([]byte("123")))
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		v   any
		exp []byte
		ok  bool
	}{
		{nil, nil, true},
		{123, []byte("123"), true},
		{int8(123), []byte("123"), true},
		{int16(123), []byte("123"), true},
		{int32(123), []byte("123"), true},
		{int64(123), []byte("123"), true},
		{uint(123), []byte("123"), true},
		{uint8(123), []byte("123"), true},
		{uint16(123), []byte("123"), true},
		{uint32(123), []byte("123"), true},
		{uint64(123), []byte("123"), true},
		{float32(123), []byte("123"), true},
		{float64(123), []byte("123"), true},
		{[]byte("123"), []byte("123"), true},
		{"123", []byte("123"), true},
		{true, []byte("true"), true},
		// special
		{time.Duration(123), []byte("123"), true},
		{fs.ModePerm, []byte("-rwxrwxrwx"), true},
		{errors.New("error msg"), []byte("error msg"), true},
		// failed
		{[]string{"123"}, nil, false},
	}

	for _, item := range tests {
		bs, err := byteutil.ToBytes(item.v)
		assert.Eq(t, item.ok, err == nil)
		assert.Eq(t, item.exp, bs, "real value: %v", item.v)
	}

	// SafeBytes
	bs := byteutil.SafeBytes([]string{"123"})
	assert.Eq(t, []byte("[123]"), bs)
}
