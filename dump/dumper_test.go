package dump

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"

	"github.com/gookit/color"
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
	}{"ab12345", "inhere", 22}
)

func newStd() *Dumper {
	return NewDumper(os.Stdout, 2)
}

func TestDumper_Fprint(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := newStd()
	dumper.WithoutColor()

	dumper.Fprint(buffer, user)
	str := buffer.String()
	assert.Contains(t, str, "{ id string; Name string; Age int }")
	assert.Contains(t, str, `id: string("ab12345"),`)
	assert.Contains(t, str, `Name: string("inhere"),`)

	dumper.ResetOptions()
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
		"abc1234",
		'a', // rune
		byte('a'),
	)

	str := buffer.String()
	str = color.ClearCode(str) // clear color codes.
	assert.Contains(t, str, "github.com/gookit/goutil/dump.TestDump_Basic(dumper_test.go")
	assert.Contains(t, str, "float64(3.1415926)")
	assert.Contains(t, str, `string("abc1234")`)

	// fmt.Println(str)
	fmt.Println(buffer.String())
}

func TestDump_Ints(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := NewDumper(buffer, 2)
	dumper.WithoutColor()

	// assert.Equal(t, 8, dumper.MoreLenNL)

	dumper.Println(ints1)
	str := buffer.String()
	buffer.Reset()
	assert.Contains(t, str, "[]int [ #len=4")
	assert.Contains(t, str, "int(1),\n")

	dumper.Print(ints2)
	str = buffer.String()
	buffer.Reset()
	assert.Contains(t, str, "[]int [ #len=11")
	assert.Contains(t, str, "int(1),\n")
	assert.NotContains(t, str, "1, 2, 3, 4")
	assert.NotContains(t, str, "[]int{1, 2, 3, 4}")

	dumper.ResetOptions()
	dumper.Dump(ints1)
	dumper.Println(ints2)
}

func TestDump_Ptr(t *testing.T) {
	buffer := new(bytes.Buffer)
	dumper := NewDumper(buffer, 2)
	// dumper.WithoutColor()

	var s string

	// refer string
	dumper.Print(&s)
	dumper.Print(s)

	s = "abc23dddd"
	dumper.Print(&s)
	dumper.Print(s)

	// refer struct
	dumper.Println(user)
	str := buffer.String()
	str = color.ClearCode(str)
	assert.Contains(t, str, "&struct")
	assert.Contains(t, str, "Age: int(22),")
	assert.Contains(t, str, `id: string("ab12345"),`)
	assert.Contains(t, str, `Name: string("inhere"),`)

	fmt.Println(buffer.String())
	// Output:
	// *struct { id string; Name string; Age int } {
	//  id: string("ab12345"),
	//  Name: string("inhere"),
	//  Age: int(22),
	// }
}

// code from https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields-in-golang
func TestDumper_AccessCantExportedField(t *testing.T) {
	type MyStruct struct {
		// id string
		id interface{}
	}

	myStruct := MyStruct{
		id: "abc111222",
	}

	// - 下面的方式适用于： 结构体指针
	rs := reflect.ValueOf(&myStruct).Elem()
	rf := rs.Field(0)

	fmt.Println(rf.CanInterface(), rf.String())
	P(myStruct)

	// rf can't be read or set.
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	// Now rf can be read and set.

	fmt.Println(rf.CanInterface(), rf.Interface())

	// - 下面的方式适用于： 结构体值
	rs = reflect.ValueOf(myStruct)
	rs2 := reflect.New(rs.Type()).Elem()
	rs2.Set(rs)
	rf = rs2.Field(0)

	fmt.Println(rf.CanInterface(), rf.String())

	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	// Now rf can be read.  Setting will succeed but only affects the temporary copy
	fmt.Println(rf.CanInterface(), rf.String())
}

// code from https://stackoverflow.com/questions/42664837/how-to-access-unexported-struct-fields-in-golang
func TestDumper_AccessCantExportedField1(t *testing.T) {
	// init an nested struct
	s1 := st1{st0{2}, 23, "inhere"}
	myS1 := struct {
		// cannotExport interface{} // ok
		cannotExport st1 // ok
		// CanExport interface{} ok
		CanExport st1 // ok
	}{
		cannotExport: s1,
		CanExport:    s1,
	}

	Println(myS1)
}

// ------------------------- map -------------------------

func TestDump_Map(t *testing.T) {
	m4 := map[string]interface{}{
		"key1": 12,
		"key2": "val1",
		"key3": [][]int{
			{23, 34},
			{230, 340},
		},
		"key4": 3.14,
		"key5": -34,
		"key6": nil,
		"key7": []int{23, 34},
		"key8": map[string]interface{} {
			"key8sub1": []int{23, 34},
			"key8sub2": []string{"a", "b"},
		},
	}
	Print(m4)
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

	newStd().WithOptions(func(opts *Options) {
		opts.IndentChar = '-'
	}).Dump(m1)
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

var (
	s1 = st1{st0{2}, 23, "inhere"}
)

func TestDump_Struct(t *testing.T) {

}

func TestStruct_WithNested(t *testing.T) {
	// buffer := new(bytes.Buffer)
	dumper := newStd()
	dumper.IndentChar = '.'
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
