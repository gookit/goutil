package encodx_test

import (
	"testing"

	"github.com/gookit/goutil/encodx"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBaseDecode(t *testing.T) {
	is := assert.New(t)

	is.Eq("GEZGCYTD", encodx.B32Encode("12abc"))
	is.Eq("12abc", encodx.B32Decode("GEZGCYTD"))

	// b23 hex
	is.Eq("64P62OJ3", encodx.B32Hex.EncodeToString([]byte("12abc")))
	// fmt.Println(time.Now().Format("20060102150405"))
	dateStr := "20230908101122"
	is.Eq("68O34CPG74O3GC9G64OJ4CG", encodx.B32Hex.EncodeToString([]byte(dateStr)))

	is.Eq("YWJj", encodx.B64Encode("abc"))
	is.Eq("abc", encodx.B64Decode("YWJj"))

	is.Eq([]byte("YWJj"), encodx.B64EncodeBytes([]byte("abc")))
	is.Eq([]byte("abc"), encodx.B64DecodeBytes([]byte("YWJj")))

	is.Eq("MTJhYmM", encodx.B64URL.EncodeToString([]byte("12abc")))
}
