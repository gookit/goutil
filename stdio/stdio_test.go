package stdio_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/stdio"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestMustReadReader(t *testing.T) {
	r := fakeobj.NewReader()
	r.WriteString("hi")
	assert.Eq(t, "hi", stdio.ReadString(r))

	r.WriteString("hi")
	r.SetErrOnRead()
	assert.Empty(t, stdio.ReadString(r))

	assert.Panics(t, func() {
		stdio.MustReadReader(r)
	})
}

func TestNewScanner(t *testing.T) {
	s := stdio.NewScanner("hi\nmy\nname\nis\ninhere")
	s = stdio.NewScanner(s)

	var ss []string
	// scan each line
	for s.Scan() {
		ss = append(ss, s.Text())
	}

	assert.Len(t, ss, 5)
	assert.Eq(t, "hi", ss[0])

	s = stdio.NewScanner(strings.NewReader("hi\nmy\nname\nis\ninhere"))
	assert.True(t, s.Scan())
	assert.Eq(t, "hi", s.Text())

	s = stdio.NewScanner([]byte("hi\nmy\nname\nis\ninhere"))
	assert.True(t, s.Scan())
	assert.Eq(t, "hi", s.Text())
}

func TestNewIOReader(t *testing.T) {
	assert.Panics(t, func() {
		stdio.NewIOReader([]int{23})
	})

	r := stdio.NewIOReader("hi")
	assert.Eq(t, "hi", stdio.ReadString(r))
	r = stdio.NewIOReader([]byte("hi"))
	assert.Eq(t, "hi", stdio.ReadString(r))
	r = stdio.NewIOReader(strings.NewReader("hi"))
	assert.Eq(t, "hi", stdio.ReadString(r))

	r = stdio.NewIOReader("hi")
	stdio.DiscardReader(r)
	assert.Eq(t, "", stdio.ReadString(r))
}

func TestWriteBytes(t *testing.T) {
	stdio.WriteByte('a')
	stdio.WritelnBytes([]byte("bc "))
	stdio.WriteBytes([]byte("hi,"))
	stdio.WriteString("inhere.")
	stdio.Writeln("welcome")
}
