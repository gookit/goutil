package comdef_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/testutil/assert"
)

// TestFormatter struct
type TestFormatter struct {
	comdef.BaseFormatter
}

func newFormatter(src any) *TestFormatter {
	return &TestFormatter{
		BaseFormatter: comdef.BaseFormatter{
			Src:    src,
			Prefix: "* ",
			Indent: "  ",
		},
	}
}

// Format implementation
func (tf *TestFormatter) Format() string {
	tf.doFormat()
	return tf.BsWriter().String()
}

// FormatTo implementation
func (tf *TestFormatter) FormatTo(w io.Writer) {
	tf.SetOutput(w)
	tf.doFormat()
}

func (tf *TestFormatter) doFormat() {
	// Custom format logic here
	_, _ = tf.BsWriter().WriteString(fmt.Sprint(tf.Src))
}

// TestTestFormatter_Format tests the Format method of TestFormatter
func TestTestFormatter_Format(t *testing.T) {
	f := newFormatter(map[string]string{"key": "value"})

	var buf bytes.Buffer
	f.SetOutput(&buf)

	expected := "map[key:value]"
	assert.Eq(t, expected, f.Format())
	assert.Eq(t, expected, buf.String())
}

// TestTestFormatter_FormatTo tests the FormatTo method of TestFormatter
func TestTestFormatter_FormatTo(t *testing.T) {
	var buf bytes.Buffer
	f := newFormatter([]string{"test", "value"})
	f.FormatTo(&buf)

	expected := "[test value]"
	assert.Eq(t, expected, buf.String())

	f.Reset()
	assert.Nil(t, f.Out)
	assert.Nil(t, f.Src)
}
