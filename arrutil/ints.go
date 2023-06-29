package arrutil

import (
	"strconv"
	"strings"

	"github.com/gookit/goutil/comdef"
)

// IntsToString convert []T to string
func IntsToString[T comdef.Integer](ints []T) string {
	if len(ints) == 0 {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteByte('[')
	for i, v := range ints {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.FormatInt(int64(v), 10))
	}
	sb.WriteByte(']')
	return sb.String()
}
