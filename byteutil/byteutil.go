package byteutil

import (
	"bytes"
)

// FirstLine from command output
func FirstLine(bs []byte) []byte {
	if i := bytes.IndexByte(bs, '\n'); i >= 0 {
		return bs[0:i]
	}
	return bs
}
