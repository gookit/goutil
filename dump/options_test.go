package dump

import (
	"fmt"
	"testing"

	"github.com/gookit/color"
	"github.com/gookit/goutil/testutil/assert"
)

func TestSkipPrivate(t *testing.T) {
	buf := newBuffer()
	dumper := newStd().WithOptions(WithoutOutput(buf), WithoutPosition(), WithoutColor())

	dumper.Println(user)
	str := buf.String()
	fmt.Print("Default: \n", str)
	assert.StrContains(t, str, "id: string(\"ab12345\")")

	buf.Reset()
	dumper.WithOptions(SkipNilField(), SkipPrivate())
	dumper.Println(user)

	str = buf.String()
	fmt.Print("SkipPrivate: \n", str)
	assert.NotContains(t, str, "id: string(\"ab12345\")")
}

// see https://github.com/gookit/goutil/issues/41
func TestSkipNilField(t *testing.T) {
	buf := newBuffer()
	dumper := newStd().WithOptions(WithoutOutput(buf), WithoutPosition(), WithoutColor())
	assert.False(t, dumper.SkipNilField)

	mp := map[string]any{
		"name": "inhere",
		"age":  nil,
	}

	dumper.Println(mp)
	str := buf.String()
	fmt.Print("Default: \n", str)
	assert.StrContains(t, str, `"age": nil`)
	buf.Reset()

	dumper.WithOptions(SkipNilField())
	assert.True(t, dumper.SkipNilField)
	dumper.Println(mp)

	str = buf.String()
	fmt.Print("SkipNilField: \n", str)
	assert.NotContains(t, str, `"age": nil`)
}

func TestWithoutColor(t *testing.T) {
	ol := color.ForceColor()
	defer func() {
		color.ForceSetColorLevel(ol)
	}()

	buf := newBuffer()
	dumper := newStd().WithOptions(WithoutOutput(buf), WithoutPosition(), WithCallerSkip(2))

	dumper.Println("a string")
	assert.Equal(t, "string(\"\x1b[0;32ma string\x1b[0m\"), \x1b[0;90m#len=8\x1b[0m\n", buf.String())

	buf.Reset()
	dumper.WithOptions(SkipNilField(), WithoutColor())
	dumper.Println("a string")
	assert.Equal(t, "string(\"a string\"), #len=8\n", buf.String())
}

// see https://github.com/gookit/goutil/issues/74
func TestBytesAsString(t *testing.T) {
	buf := newBuffer()
	dumper := newStd().WithOptions(WithoutOutput(buf), WithoutColor())

	bts := []byte("hello")
	dumper.Println(bts)
	str := buf.String()
	fmt.Print(str)
	assert.StrContains(t, str, "[]uint8 [ #len=5,cap=5")
	assert.StrContains(t, str, "uint8(104),")
	/* Output:
	PRINT AT github.com/gookit/goutil/dump.TestBytesAsString(options_test.go:28)
	[]uint8 [ #len=5,cap=5
	  uint8(104),
	  uint8(101),
	  uint8(108),
	  uint8(108),
	  uint8(111),
	],
	*/

	buf.Reset()
	dumper.WithOptions(WithoutType(), BytesAsString())
	dumper.Print(bts)
	str = buf.String()
	fmt.Print("BytesAsString: \n", str)
	assert.StrContains(t, str, "[]byte(\"hello\"), #len=5,cap=5")
	/*Output:
	PRINT AT github.com/gookit/goutil/dump.TestBytesAsString(options_test.go:49)
	[]byte("hello"), #len=5,cap=5
	*/
}
