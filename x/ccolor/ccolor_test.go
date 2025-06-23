package ccolor_test

import (
	"testing"

	"github.com/gookit/goutil/byteutil"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/ccolor"
)

func TestCommon(t *testing.T) {
	assert.Nil(t, ccolor.LastErr())

	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()
	assert.Gt(t, int(ccolor.Level()), 0)

	ccolor.Disable()
	assert.False(t, ccolor.IsSupportColor())
	assert.False(t, ccolor.IsSupport256Color())
	assert.False(t, ccolor.IsSupportTrueColor())
}

func TestPrint_func(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	buf := byteutil.NewBuffer()
	ccolor.SetOutput(buf)

	ccolor.Print("<red>text</>")
	assert.Equal(t, "\x1b[0;31mtext\x1b[0m", buf.ResetGet())

	// Printf
	ccolor.Printf("<red>text %s</>", "msg")
	assert.Equal(t, "\x1b[0;31mtext msg\x1b[0m", buf.ResetGet())

	// Sprint
	r := ccolor.Sprint("<red>text</>")
	assert.Equal(t, "\x1b[0;31mtext\x1b[0m", r)

	// Sprintf
	r = ccolor.Sprintf("<red>text %s</>", "msg")
	assert.Equal(t, "\x1b[0;31mtext msg\x1b[0m", r)
}

/*************************************************************
 * test 16 color
 *************************************************************/

func TestColor16_usage(t *testing.T) {
	ccolor.ForceEnableColor()
	defer ccolor.RevertColorSupport()

	is := assert.New(t)
	buf := byteutil.NewBuffer()
	ccolor.SetOutput(buf)

	is.True(ccolor.Bold.IsValid())
	r := ccolor.Bold.Render("text")
	is.Equal("\x1b[1mtext\x1b[0m", r)
	r = ccolor.LightYellow.Text("text")
	is.Equal("\x1b[93mtext\x1b[0m", r)
	r = ccolor.LightWhite.Sprint("text")
	is.Equal("\x1b[97mtext\x1b[0m", r)
	r = ccolor.White.Render("test", "spaces")
	is.Equal("\x1b[37mtestspaces\x1b[0m", r)
	r = ccolor.Black.Renderln("test", "spaces")
	is.Equal("\x1b[30mtest spaces\x1b[0m", r)

	str := ccolor.Red.Sprintf("A %s", "MSG")
	is.Equal("\x1b[31mA MSG\x1b[0m", str)
	// is.Equal("red", ccolor.Red.Name())
	// is.Equal("unknown", ccolor.Color(123).Name())

	// Color.Print
	ccolor.FgGray.Print("MSG")
	str = buf.ResetGet()
	is.Equal("\x1b[90mMSG\x1b[0m", str)

	// Color.Printf
	ccolor.BgGray.Printf("A %s", "MSG")
	str = buf.String()
	is.Equal("\x1b[100mA MSG\x1b[0m", str)
	buf.Reset()

	// Color.Println
	ccolor.LightMagenta.Println("MSG")
	str = buf.ResetGet()
	is.Equal("\x1b[95mMSG\x1b[0m\n", str)
	// is.Equal("lightMagenta", ccolor.LightMagenta.Name())

	ccolor.LightMagenta.Println()
	str = buf.ResetGet()
	is.Equal("\n", str)

	ccolor.LightCyan.Print("msg")
	ccolor.LightRed.Printf("m%s", "sg")
	ccolor.LightGreen.Println("msg")
	str = buf.ResetGet()
	is.Equal("\x1b[96mmsg\x1b[0m\x1b[91mmsg\x1b[0m\x1b[92mmsg\x1b[0m\n", str)
}

func TestColor_convert(t *testing.T) {
	is := assert.New(t)

	is.True(ccolor.FgBlue.IsFg())
	is.True(ccolor.FgBlue.IsValid())
	is.True(ccolor.BgBlue.IsBg())
	is.False(ccolor.FgBlue.IsBg())
	is.False(ccolor.FgBlue.IsOption())

	// to fg
	is.Equal(ccolor.FgBlue, ccolor.FgBlue.ToFg())
	is.Equal(ccolor.FgCyan, ccolor.BgCyan.ToFg())

	// to bg
	val := ccolor.FgBlue.ToBg()
	is.Equal(ccolor.BgBlue, val)
	is.Equal(ccolor.FgCyan, ccolor.BgCyan.ToFg())

	// Color.Darken
	blue := ccolor.LightBlue.Darken()
	is.Equal(94, int(ccolor.LightBlue))
	is.Equal(34, int(blue))
	c := ccolor.Color(120).Darken()
	is.Equal(120, int(c))

	// Color.Light
	lightCyan := ccolor.Cyan.Light()
	is.Equal(36, int(ccolor.Cyan))
	is.Equal(96, int(lightCyan))
	c = ccolor.Color(120).Light()
	is.Equal(120, int(c))
}

func TestColor_vars(t *testing.T) {
	is := assert.New(t)

	_, ok := ccolor.FgColors["red"]
	is.True(ok)
	_, ok = ccolor.ExFgColors["lightRed"]
	is.True(ok)
	_, ok = ccolor.BgColors["red"]
	is.True(ok)
	_, ok = ccolor.ExBgColors["lightRed"]
	is.True(ok)
}
