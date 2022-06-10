package stdio

import (
	"io"
	"io/ioutil"
	"strings"
)

// QuietWriteString to writer
func QuietWriteString(w io.Writer, ss ...string) {
	_, _ = io.WriteString(w, strings.Join(ss, ""))
}

// DiscardReader anything from the reader
func DiscardReader(src io.Reader) {
	_, _ = io.Copy(ioutil.Discard, src)
}

// MustReadReader read contents from io.Reader, will panic on error
func MustReadReader(r io.Reader) []byte {
	// TODO go 1.16+ bs, err := io.ReadAll(r)
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return bs
}
