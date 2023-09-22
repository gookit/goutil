package strutil

import (
	mRand "math/rand"
	"time"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/encodes"
)

// some consts string chars
const (
	Numbers  = "0123456789"
	HexChars = "0123456789abcdef" // base16

	AlphaBet  = "abcdefghijklmnopqrstuvwxyz"
	AlphaBet1 = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"

	AlphaNum  = "abcdefghijklmnopqrstuvwxyz0123456789"
	AlphaNum2 = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	AlphaNum3 = "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
)

func newRand() *mRand.Rand {
	return mRand.New(mRand.NewSource(time.Now().UnixNano()))
}

// RandomChars generate give length random chars at `a-z`
func RandomChars(ln int) string {
	cs := make([]byte, ln)
	// UnixNano: 1607400451937462000
	rn := newRand()

	for i := 0; i < ln; i++ {
		// rand in 0 - 25
		cs[i] = AlphaBet[rn.Intn(25)]
	}
	return string(cs)
}

// RandomCharsV2 generate give length random chars in `0-9a-z`
func RandomCharsV2(ln int) string {
	cs := make([]byte, ln)
	// UnixNano: 1607400451937462000
	rn := newRand()

	for i := 0; i < ln; i++ {
		// rand in 0 - 35
		cs[i] = AlphaNum[rn.Intn(35)]
	}
	return string(cs)
}

// RandomCharsV3 generate give length random chars in `0-9a-zA-Z`
func RandomCharsV3(ln int) string {
	cs := make([]byte, ln)
	// UnixNano: 1607400451937462000
	rn := newRand()

	for i := 0; i < ln; i++ {
		// rand in 0 - 61
		cs[i] = AlphaNum2[rn.Intn(61)]
	}
	return string(cs)
}

// RandWithTpl generate random string with give template
func RandWithTpl(n int, letters string) string {
	if len(letters) == 0 {
		letters = AlphaNum2
	}

	ln := len(letters)
	cs := make([]byte, n)
	rn := newRand()

	for i := 0; i < n; i++ {
		// rand in 0 - ln
		cs[i] = letters[rn.Intn(ln)]
	}
	return byteutil.String(cs)
}

// RandomString generate.
//
// Example:
//
//	// this will give us a 44 byte, base64 encoded output
//	token, err := RandomString(16) // eg: "I7S4yFZddRMxQoudLZZ-eg"
func RandomString(length int) (string, error) {
	b, err := RandomBytes(length)
	return encodes.B64URL.EncodeToString(b), err
}

// RandomBytes generate
func RandomBytes(length int) ([]byte, error) {
	return byteutil.Random(length)
}
