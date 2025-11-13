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

func TestMicroTimeBaseID(t *testing.T) {
	fmt.Println("Base 16:")
	for i := 0; i < 10; i++ {
		id := strutil.MicroTimeHexID()
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}

	fmt.Println("Base 36:")
	for i := 0; i < 10; i++ {
		id := strutil.MTimeBaseID(36)
		fmt.Println(id, "len:", len(id))
		assert.NotEmpty(t, id)
	}
	assert.NotEmpty(t, strutil.MTimeBase36())
}

func TestDateSN(t *testing.T) {
	for i := 0; i < 10; i++ {
		no := strutil.DatetimeNo("test")
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}

func TestDateSNv2(t *testing.T) {
	for i := 0; i < 10; i++ {
		no := strutil.DateSNv2("T")
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}

	for i := 0; i < 10; i++ {
		no := strutil.DateSNv2("T", 36)
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}

func TestDateSNv3(t *testing.T) {
	no := strutil.DateSNv3("", 8)
	fmt.Println(no, "len:", len(no))

	fmt.Println("------ dateLen=8, base=32:")
	for i := 0; i < 10; i++ {
		no = strutil.DateSNv3("T", 8)
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}

	fmt.Println("------ dateLen=6, base=36:")
	for i := 0; i < 6; i++ {
		no = strutil.DateSNv3("T", 6, 36)
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}

	fmt.Println("------ dateLen=6, base=48:")
	for i := 0; i < 6; i++ {
		no = strutil.DateSNv3("T", 6, 48)
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}

	fmt.Println("------ dateLen=4, base=62:")
	for i := 0; i < 6; i++ {
		no = strutil.DateSNv3("T", 4, 62)
		fmt.Println(no, "len:", len(no))
		assert.NotEmpty(t, no)
	}
}
