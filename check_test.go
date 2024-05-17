package goutil_test

import (
	"testing"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestIsEmpty(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.IsEmpty(nil))
	is.False(goutil.IsZero("abc"))
	is.False(goutil.IsEmpty("abc"))

	is.True(goutil.IsEmptyReal(nil))
	is.False(goutil.IsZeroReal("abc"))
	is.False(goutil.IsEmptyReal("abc"))
}

func TestIsFunc(t *testing.T) {
	is := assert.New(t)

	is.False(goutil.IsFunc(nil))
	is.True(goutil.IsFunc(goutil.IsEqual))
}

func TestIsEqual(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.IsNil(nil))
	is.False(goutil.IsNil("abc"))

	is.True(goutil.IsEqual("a", "a"))
	is.True(goutil.IsEqual([]string{"a"}, []string{"a"}))
	is.True(goutil.IsEqual(23, 23))
	is.True(goutil.IsEqual(nil, nil))
	is.True(goutil.IsEqual([]byte("abc"), []byte("abc")))

	is.False(goutil.IsEqual([]byte("abc"), "abc"))
	is.False(goutil.IsEqual(nil, 23))
	is.False(goutil.IsEqual(goutil.IsEmpty, 23))
}

func TestIsContains(t *testing.T) {
	is := assert.New(t)

	is.True(goutil.Contains("abc", "a"))
	is.True(goutil.Contains([]string{"abc", "def"}, "abc"))
	is.True(goutil.Contains(map[int]string{2: "abc", 4: "def"}, 4))

	is.False(goutil.Contains("abc", "def"))
	is.False(goutil.IsContains("abc", "def"))
}
