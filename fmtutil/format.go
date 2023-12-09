package fmtutil

import (
	"encoding/json"
	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/mathutil"
	"github.com/gookit/goutil/strutil"
)

// data size
const (
	OneKByte = 1024
	OneMByte = 1024 * 1024
	OneGByte = 1024 * 1024 * 1024
)

// DataSize format bytes number friendly.
//
// Usage:
//
//	file, err := os.Open(path)
//	fl, err := file.Stat()
//	fmtSize := DataSize(fl.Size())
func DataSize(size uint64) string {
	return mathutil.DataSize(size)
}

// SizeToString alias of the DataSize
func SizeToString(size uint64) string { return DataSize(size) }

// StringToByte alias of the ParseByte
func StringToByte(sizeStr string) uint64 { return ParseByte(sizeStr) }

// ParseByte converts size string like 1GB/1g or 12mb/12M into an unsigned integer number of bytes
func ParseByte(sizeStr string) uint64 {
	val, _ := strutil.ToByteSize(sizeStr)
	return val
}

// PrettyJSON get pretty Json string
func PrettyJSON(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// ArgsWithSpaces it like Println, will add spaces for each argument
func ArgsWithSpaces(vs []any) (message string) {
	ln := len(vs)
	if ln == 0 {
		return ""
	}
	if ln == 1 {
		return strutil.SafeString(vs[0])
	}

	bs := make([]byte, 0, ln*8)
	for i := range vs {
		if i > 0 { // add space
			bs = append(bs, ' ')
		}
		bs = byteutil.AppendAny(bs, vs[i])
	}
	return string(bs)
}
