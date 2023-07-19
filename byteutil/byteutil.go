package byteutil

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Md5 Generate a 32-bit md5 bytes
func Md5(src any) []byte {
	h := md5.New()

	if s, ok := src.(string); ok {
		h.Write([]byte(s))
	} else {
		h.Write([]byte(fmt.Sprint(src)))
	}
	return h.Sum(nil)
}

// Random bytes generate
func Random(length int) ([]byte, error) {
	b := make([]byte, length)
	// Note that err == nil only if we read len(b) bytes.
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// FirstLine from command output
func FirstLine(bs []byte) []byte {
	if i := bytes.IndexByte(bs, '\n'); i >= 0 {
		return bs[0:i]
	}
	return bs
}

// AppendAny append any value to byte slice
func AppendAny(dst []byte, v any) []byte {
	if v == nil {
		return append(dst, "<nil>"...)
	}

	switch val := v.(type) {
	case []byte:
		dst = append(dst, val...)
	case string:
		dst = append(dst, val...)
	case int:
		dst = strconv.AppendInt(dst, int64(val), 10)
	case int8:
		dst = strconv.AppendInt(dst, int64(val), 10)
	case int16:
		dst = strconv.AppendInt(dst, int64(val), 10)
	case int32:
		dst = strconv.AppendInt(dst, int64(val), 10)
	case int64:
		dst = strconv.AppendInt(dst, val, 10)
	case uint:
		dst = strconv.AppendUint(dst, uint64(val), 10)
	case uint8:
		dst = strconv.AppendUint(dst, uint64(val), 10)
	case uint16:
		dst = strconv.AppendUint(dst, uint64(val), 10)
	case uint32:
		dst = strconv.AppendUint(dst, uint64(val), 10)
	case uint64:
		dst = strconv.AppendUint(dst, val, 10)
	case float32:
		dst = strconv.AppendFloat(dst, float64(val), 'f', -1, 32)
	case float64:
		dst = strconv.AppendFloat(dst, val, 'f', -1, 64)
	case bool:
		dst = strconv.AppendBool(dst, val)
	case time.Time:
		dst = val.AppendFormat(dst, time.RFC3339)
	case time.Duration:
		dst = strconv.AppendInt(dst, int64(val), 10)
	case error:
		dst = append(dst, val.Error()...)
	case fmt.Stringer:
		dst = append(dst, val.String()...)
	default:
		dst = append(dst, fmt.Sprint(v)...)
	}
	return dst
}

// Cut bytes by one byte char. like bytes.Cut(), but sep is byte.
func Cut(bs []byte, sep byte) (before, after []byte, found bool) {
	return bytes.Cut(bs, []byte{sep})
}

// SafeCut bytes by one byte char. always return before and after
func SafeCut(bs []byte, sep byte) (before, after []byte) {
	before, after, _ = bytes.Cut(bs, []byte{sep})
	return
}

// SafeCuts bytes by sub bytes. like the bytes.Cut(), but always return before and after
func SafeCuts(bs []byte, sep []byte) (before, after []byte) {
	before, after, _ = bytes.Cut(bs, sep)
	return
}
