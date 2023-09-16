package encodes_test

import (
	"testing"

	"github.com/gookit/goutil/encodes"
	"github.com/gookit/goutil/testutil/assert"
)

func TestBaseDecode(t *testing.T) {
	is := assert.New(t)

	is.Eq("GEZGCYTD", encodes.B32Encode("12abc"))
	is.Eq("12abc", encodes.B32Decode("GEZGCYTD"))

	// b23 hex
	is.Eq("64P62OJ3", encodes.B32Hex.EncodeToString([]byte("12abc")))
	// fmt.Println(time.Now().Format("20060102150405"))
	dateStr := "20230908101122"
	is.Eq("68O34CPG74O3GC9G64OJ4CG", encodes.B32Hex.EncodeToString([]byte(dateStr)))

	is.Eq("YWJj", encodes.B64Encode("abc"))
	is.Eq("abc", encodes.B64Decode("YWJj"))

	is.Eq([]byte("YWJj"), encodes.B64EncodeBytes([]byte("abc")))
	is.Eq([]byte("abc"), encodes.B64DecodeBytes([]byte("YWJj")))

	is.Eq("MTJhYmM", encodes.B64URL.EncodeToString([]byte("12abc")))
}
