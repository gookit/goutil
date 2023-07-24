package cliutil_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/testutil/fakeobj"
)

func TestRead_cases(t *testing.T) {
	buf := fakeobj.NewReader()
	cliutil.Input = buf
	defer func() {
		cliutil.Input = os.Stdin
	}()

	t.Run("ReadInput", func(t *testing.T) {
		buf.WriteString("inhere")
		ans, err := cliutil.ReadInput("hi, your name? ")
		fmt.Println("ans:", ans)
		assert.NoError(t, err)
		assert.Equal(t, "inhere", ans)

		// error
		buf.SetErrOnRead()
		_, err = cliutil.ReadInput("hi, your name? ")
		assert.Error(t, err)
		fmt.Println(err)
		buf.ErrOnRead = false
	})

	// test ReadLine
	t.Run("ReadLine", func(t *testing.T) {
		buf.WriteString("inhere")
		ans, err := cliutil.ReadLine("hi, your name? ")
		fmt.Println("ans:", ans)
		assert.NoError(t, err)
		assert.Equal(t, "inhere", ans)
	})

	// test ReadFirst
	t.Run("ReadFirst", func(t *testing.T) {
		buf.WriteString("Y")
		ans, err := cliutil.ReadFirst("continue?[y/n] ")
		fmt.Println("ans:", ans)
		assert.NoError(t, err)
		assert.Equal(t, "Y", ans)
	})

	// test ReadFirstRune
	t.Run("ReadFirstRune", func(t *testing.T) {
		buf.WriteString("Y")
		ans, err := cliutil.ReadFirstRune("continue?[y/n] ")
		fmt.Println("ans:", string(ans))
		assert.NoError(t, err)
		assert.Equal(t, 'Y', ans)
	})

	// test ReadAsBool
	t.Run("ReadAsBool", func(t *testing.T) {
		buf.WriteString("Y")
		ans := cliutil.ReadAsBool("continue?[y/n] ", false)
		fmt.Println("ans:", ans)
		assert.True(t, ans)
	})

	// test Confirm
	t.Run("Confirm", func(t *testing.T) {
		buf.WriteString("Y")
		ans := cliutil.Confirm("continue?", false)
		fmt.Println(ans)
		assert.True(t, ans)
		ans = cliutil.Confirm("continue?", true)
		fmt.Println(ans)
		assert.True(t, ans)
	})
}

func TestReadFirst(t *testing.T) {
	// testutil.RewriteStdout()
	// _, err := os.Stdout.WriteString("haha")
	// ans, err1 := cliutil.ReadFirst("hi?")
	// testutil.RestoreStdout()
	// assert.NoError(t, err)
	// assert.NoError(t, err1)
	// assert.Equal(t, "haha", ans)
}

func TestInputIsYes(t *testing.T) {
	tests := []struct {
		in  string
		wnt bool
	}{
		{"y", true},
		{"yes", true},
		{"yES", true},
		{"Y", true},
		{"Yes", true},
		{"YES", true},
		{"h", false},
		{"n", false},
		{"no", false},
		{"NO", false},
	}
	for _, test := range tests {
		assert.Eq(t, test.wnt, cliutil.InputIsYes(test.in))
	}
}

func TestByteIsYes(t *testing.T) {
	tests := []struct {
		in  byte
		wnt bool
	}{
		{'y', true},
		{'Y', true},
		{'h', false},
		{'n', false},
		{'N', false},
	}
	for _, test := range tests {
		assert.Eq(t, test.wnt, cliutil.ByteIsYes(test.in))
	}
}
