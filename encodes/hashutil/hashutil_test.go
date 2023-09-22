package hashutil_test

import (
	"testing"

	"github.com/gookit/goutil/encodes/hashutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestHash(t *testing.T) {
	tests := []struct {
		src  any
		algo string
		want string
	}{
		{"abc12", "crc32", "b744b523"},
		{"abc12", "crc64", "41b31776c4200000"},
		{"abc12", "md5", "b2157e7b2ae716a747597717f1efb7a0"},
		{"abc12", "sha1", "8fe670fef2b8c74ef8987cdfccdb32e96ad4f9a2"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, hashutil.Hash(tt.algo, tt.src))
	}

	assert.Panics(t, func() {
		hashutil.Hash("unknown", nil)
	})
}

func TestHash32(t *testing.T) {
	tests := []struct {
		src  any
		algo string
		want string
	}{
		{"abc12", "crc32", "MT2BA8O"},
		{"abc12", "crc64", "86PHETM440000"},
		{"abc12", "md5", "M8ANSUPASSBAEHQPESBV3RTNK0"},
		{"abc12", "sha1", "HVJ71VNIN33KTU4OFJFSPMPIT5LD9UD2"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, hashutil.Hash32(tt.algo, tt.src))
	}
}

func TestHash64(t *testing.T) {
	tests := []struct {
		src  any
		algo string
		want string
	}{
		{"abc12", "crc32", "t0S1Iw"},
		{"abc12", "crc64", "QbMXdsQgAAA"},
		{"abc12", "md5", "shV+eyrnFqdHWXcX8e+3oA"},
		{"abc12", "sha1", "j+Zw/vK4x074mHzfzNsy6WrU+aI"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, hashutil.Hash64(tt.algo, tt.src))
	}
}

func TestMd5(t *testing.T) {
	assert.NotEmpty(t, hashutil.MD5("abc"))
	assert.NotEmpty(t, hashutil.MD5([]int{12, 34}))

	assert.Eq(t, "202cb962ac59075b964b07152d234b70", hashutil.MD5("123"))
	assert.Eq(t, "900150983cd24fb0d6963f7d28e17f72", hashutil.MD5("abc"))

	// short md5
	assert.Eq(t, "ac59075b964b0715", hashutil.ShortMD5("123"))
	assert.Eq(t, "3cd24fb0d6963f7d", hashutil.ShortMD5("abc"))
}
