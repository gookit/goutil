package dump

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBufDumper(buf *bytes.Buffer) *Dumper {
	return NewDumper(buf, 2)
}

var (
	ints1 = []int{1, 2, 3, 4}
	ints2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	user  = &struct {
		id   string
		Name string
		Age  int
	}{"ab1234", "inhere", 22}
)

func newStd() *Dumper {
	return NewDumper(os.Stdout, 2)
}

func TestDumper_Fprint(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := newStd()

	dumper.Fprint(buffer, user)
	str := buffer.String()
	assert.Contains(t, str, "{ id string; Name string; Age int }")
	assert.Contains(t, str, `id: string("ab1234"),`)
	assert.Contains(t, str, `Name: string("inhere"),`)

	dumper.Print(user)
}

func TestDump_Basic(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := NewDumper(buffer, 2)

	dumper.Dump(
		nil,
		// bool
		false,
		true,
		// int(X)
		12,
		int8(12),
		int16(12),
		int32(12),
		int64(12),
		// uint(X)
		uint(12),
		uint8(12),
		uint16(12),
		uint32(12),
		uint64(12),
		// float
		float32(3.1415926),
		3.1415926, // float64
		// string
		"abc123",
		'a', // rune
		byte('a'),
	)

	str := buffer.String()
	assert.Contains(t, str, "github.com/gookit/goutil/dump.TestDumper_Dump_Basic(dumper_test.go")
	assert.Contains(t, str, "float64(3.1415926)")
	assert.Contains(t, str, `string("abc123")`)

	fmt.Println(str)
}

func TestDump_Ints(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := NewDumper(buffer, 2)

	assert.Equal(t, 8, dumper.MoreLenNL)

	dumper.Println(ints1)
	str := buffer.String()
	buffer.Reset()
	assert.Contains(t, str, "1, 2, 3, 4")
	assert.Contains(t, str, "[]int{1, 2, 3, 4}")
	fmt.Println(str)

	// elements > 5
	dumper.Print(ints2)
	str = buffer.String()
	buffer.Reset()
	assert.NotContains(t, str, "1, 2, 3, 4")
	assert.NotContains(t, str, "[]int{1, 2, 3, 4}")
	fmt.Println(str)
}

func TestDump_Ptr(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := NewDumper(buffer, 2)

	var s string

	// refer string
	dumper.Print(&s)
	dumper.Print(s)

	s = "abc"
	dumper.Print(&s)
	dumper.Print(s)

	// refer struct
	dumper.Println(user)

	str := buffer.String()
	assert.Contains(t, str, "*struct")
	assert.Contains(t, str, "Age: int(22),")
	assert.Contains(t, str, "id: string(\"ab1234\"),")
	assert.Contains(t, str, "Name: string(\"inhere\"),")

	fmt.Println(str)
	// Output:
	// *struct { id string; Name string; Age int } {
	//  id: string("ab1234"),
	//  Name: string("inhere"),
	//  Age: int(22),
	// }
}

// ------------------------- map -------------------------

func TestDump_Map(t *testing.T) {

}

func TestMap_Simpled(t *testing.T) {
	m1 := map[int]int{
		23: 12,
		24: 13,
	}

	m2 := map[string]int{
		"key1": 12,
		"key2": 13,
	}

	m3 := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	P(m1, m2, m3)
	/*
		Output:
		PRINT AT github.com/gookit/goutil/dump.TestMap_Simpled(dump_test.go:309)
		map[int]int {
		  24: int(13),
		  23: int(12),
		}
		map[string]int {
		  key1: int(12),
		  key2: int(13),
		}
		map[string]string {
		  key1: string("val1"),
		  key2: string("val2"),
		}

	*/

	m4 := map[string]interface{}{
		"key1": 12,
		"key2": "val1",
		"key3": 34,
		"key4": 3.14,
		"key5": -34,
		"key6": nil,
	}
	Print(m4)
	/*
		PRINT AT github.com/gookit/goutil/dump.TestMap_Simpled(dump_test.go:335)
		map[string]interface {} {
		  key4: float64(3.14),
		  key5: int(-34),
		  key6: <nil>,
		  key1: int(12),
		  key2: string("val1"),
		  key3: int(34),
		}
	*/
}

func TestMap_InterfaceNested(t *testing.T) {
	s1 := st1{st0{2}, 23, "inhere"}
	m1 := map[string]interface{}{
		"key1": 112,
		"key2": uint(112),
		"key3": int64(112),
		"key4": 112.23,
		"key5": nil,
		"key6": 'b', // rune
		"key7": byte('a'),
		"st1":  s1,
		"user": user,
		"submap1": map[string]int{
			"key1": 12,
			"key2": 13,
		},
		"submap2": map[string]interface{}{
			"key1": 12,
			"key2": "abc",
			"submap21": map[string]string{
				"key1": "val1",
				"key2": "val2",
			},
		},
		"submap3": map[string]interface{}{
			"key1": 12,
			"key2": "abc",
			"submap31": map[string]interface{}{
				"key31": 12,
				"key32": 13,
				"user":  user,
				"submap311": map[string]int{
					"key1": 12,
					"key2": 13,
				},
			},
		},
	}

	newStd().Dump(m1)
}

var (
	myOpts = struct {
		opt0 *int
		opt1 bool
		opt2 int
		opt3 float64
		opt4 string
	}{nil, true, 22, 34.45, "abc"}
)

// ------------------------- map -------------------------

type st0 struct {
	Sex int
}

type st1 struct {
	st0
	Age  int
	Name string
}

func TestDump_Struct(t *testing.T) {

}

func TestStruct_WithNested(t *testing.T) {
	// buffer := new(bytes.Buffer)
	dumper := newStd()
	dumper.IndentChar = '.'

	s1 := st1{st0{2}, 23, "inhere"}

	dumper.Println(s1)
	// OUT:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:223)
	// struct { dump.st0; Age int; Name string } {
	//  st0: dump.st0 {
	//    Sex: 2,
	//  },
	//  Age: 23,
	//  Name: "inhere",
	// }

	type st2 struct {
		st1
		Github string
	}

	s2 := st2{st1: s1, Github: "https://github.com/inhere"}
	dumper.IndentChar = ' '
	dumper.Print(s2)

	// Output:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:257)
	// dump.st2 {
	//  st1: dump.st1 {
	//    st0: dump.st0 {
	//      Sex: int(2),
	//    },
	//    Age: int(23),
	//    Name: string("inhere"),
	//  },
	//  Github: string("https://github.com/inhere"),
	// }

	s3 := struct {
		st1
		Github string
	}{st1: s1, Github: "https://github.com/inhere"}

	dumper.IndentChar = '.'
	dumper.Print(s3)

	// Output:
	// PRINT AT github.com/gookit/goutil/dump.TestStruct_WithNested(dump_test.go:278)
	// struct { dump.st1; Github string } {
	//  st1: dump.st1 {
	//    st0: dump.st0 {
	//      Sex: int(2),
	//    },
	//    Age: int(23),
	//    Name: string("inhere"),
	//  },
	//  Github: string("https://github.com/inhere"),
	// }
}
