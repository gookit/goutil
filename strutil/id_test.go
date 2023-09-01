package strutil_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestMicroTimeID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}

	id := strutil.MicroTimeID()
	fmt.Println("mtID :", id, "len:", len(id))
	b16id := strutil.Base10Conv(id, 16)
	fmt.Println("b16id:", b16id, "len:", len(b16id))
	b32id := strutil.Base10Conv(id, 32)
	fmt.Println("b32id:", b32id, "len:", len(b32id))
	b36id := strutil.Base10Conv(id, 36)
	fmt.Println("b36id:", b36id, "len:", len(b36id))
	b62id := strutil.Base10Conv(id, 62)
	fmt.Println("b62id:", b62id, "len:", len(b62id))
}

func TestMicroTimeHexID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeHexID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}
}

func TestDatetimeNo(t *testing.T) {
	for i := 0; i < 10; i++ {
		no := strutil.DatetimeNo("test")
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}
