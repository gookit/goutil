package byteutil_test

import (
	"errors"
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestFirstLine(t *testing.T) {
	bs := []byte("hi\ninhere")
	assert.Eq(t, []byte("hi"), byteutil.FirstLine(bs))
	assert.Eq(t, []byte("hi"), byteutil.FirstLine([]byte("hi")))
}

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

func TestMd5(t *testing.T) {
	assert.NotEmpty(t, byteutil.Md5("abc"))
	assert.NotEmpty(t, byteutil.Md5([]int{12, 34}))
}

func TestAppendAny(t *testing.T) {
	assert.Eq(t, []byte("123"), byteutil.AppendAny(nil, 123))
	assert.Eq(t, []byte("123"), byteutil.AppendAny([]byte{}, 123))
	assert.Eq(t, []byte("123"), byteutil.AppendAny([]byte("1"), 23))
	assert.Eq(t, []byte("1<nil>"), byteutil.AppendAny([]byte("1"), nil))
	assert.Eq(t, "3600000000000", string(byteutil.AppendAny([]byte{}, timex.OneHour)))
}

func TestCut(t *testing.T) {
	// test for byteutil.Cut()
	b, a, ok := byteutil.Cut([]byte("age=123"), '=')
	assert.True(t, ok)
	assert.Eq(t, []byte("age"), b)
	assert.Eq(t, []byte("123"), a)

	b, a, ok = byteutil.Cut([]byte("age=123"), 'x')
	assert.False(t, ok)
	assert.Eq(t, []byte("age=123"), b)
	assert.Empty(t, a)
}
