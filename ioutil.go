package goutil

import (
	"io"
	"strings"
)

// QuietWriteString to writer
func QuietWriteString(w io.Writer, ss ...string) {
	_, _ = io.WriteString(w, strings.Join(ss, ""))
}
