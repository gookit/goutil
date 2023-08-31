package byteutil_test

import (
	"errors"
	"io/fs"
	"testing"
	"time"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/timex"
)

func TestRandom(t *testing.T) {
	bs, err := byteutil.Random(10)
	assert.NoError(t, err)
	assert.Len(t, bs, 10)

	bs, err = byteutil.Random(0)
	assert.NoError(t, err)
	assert.Len(t, bs, 0)
}

func TestFirstLine(t *testing.T) {
	bs := []byte("hi\ninhere")
	assert.Eq(t, []byte("hi"), byteutil.FirstLine(bs))
	assert.Eq(t, []byte("hi"), byteutil.FirstLine([]byte("hi")))
}

func TestMd5(t *testing.T) {
	assert.NotEmpty(t, byteutil.Md5("abc"))
	assert.NotEmpty(t, byteutil.Md5([]int{12, 34}))

	assert.Eq(t, "202cb962ac59075b964b07152d234b70", string(byteutil.Md5("123")))
	assert.Eq(t, "900150983cd24fb0d6963f7d28e17f72", string(byteutil.Md5("abc")))

	// short md5
	assert.Eq(t, "ac59075b964b0715", string(byteutil.ShortMd5("123")))
	assert.Eq(t, "3cd24fb0d6963f7d", string(byteutil.ShortMd5("abc")))
}

func TestAppendAny(t *testing.T) {
	assert.Eq(t, []byte("123"), byteutil.AppendAny(nil, 123))
	assert.Eq(t, []byte("123"), byteutil.AppendAny([]byte{}, 123))
	assert.Eq(t, []byte("123"), byteutil.AppendAny([]byte("1"), 23))
	assert.Eq(t, []byte("1<nil>"), byteutil.AppendAny([]byte("1"), nil))
	assert.Eq(t, "3600000000000", string(byteutil.AppendAny([]byte{}, timex.OneHour)))

	tests := []struct {
		dst []byte
		v   any
		exp []byte
	}{
		{nil, 123, []byte("123")},
		{[]byte{}, 123, []byte("123")},
		{[]byte("1"), 23, []byte("123")},
		{[]byte("1"), nil, []byte("1<nil>")},
		{[]byte{}, timex.OneHour, []byte("3600000000000")},
		{[]byte{}, int8(123), []byte("123")},
		{[]byte{}, int16(123), []byte("123")},
		{[]byte{}, int32(123), []byte("123")},
		{[]byte{}, int64(123), []byte("123")},
		{[]byte{}, uint(123), []byte("123")},
		{[]byte{}, uint8(123), []byte("123")},
		{[]byte{}, uint16(123), []byte("123")},
		{[]byte{}, uint32(123), []byte("123")},
		{[]byte{}, uint64(123), []byte("123")},
		{[]byte{}, float32(123), []byte("123")},
		{[]byte{}, float64(123), []byte("123")},
		{[]byte{}, "123", []byte("123")},
		{[]byte{}, []byte("123"), []byte("123")},
		{[]byte{}, []int{1, 2, 3}, []byte("[1 2 3]")},
		{[]byte{}, []string{"1", "2", "3"}, []byte("[1 2 3]")},
		{[]byte{}, true, []byte("true")},
		{[]byte{}, fs.ModePerm, []byte("-rwxrwxrwx")},
		{[]byte{}, errors.New("error msg"), []byte("error msg")},
	}

	for _, tt := range tests {
		assert.Eq(t, tt.exp, byteutil.AppendAny(tt.dst, tt.v))
	}

	tim, err := time.Parse(time.RFC3339, "2019-01-01T00:00:00Z")
	assert.NoError(t, err)
	assert.Eq(t, []byte("2019-01-01T00:00:00Z"), byteutil.AppendAny(nil, tim))
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

	// SafeCut
	b, a = byteutil.SafeCut([]byte("age=123"), '=')
	assert.Eq(t, []byte("age"), b)
	assert.Eq(t, []byte("123"), a)

	// SafeCuts
	b, a = byteutil.SafeCuts([]byte("age=123"), []byte{'='})
	assert.Eq(t, []byte("age"), b)
	assert.Eq(t, []byte("123"), a)
}
