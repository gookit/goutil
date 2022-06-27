package cliutil_test

import (
	"strings"
	"testing"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestCurrentShell(t *testing.T) {
	path := cliutil.CurrentShell(true)

	if path != "" {
		assert.NotEmpty(t, path)
		assert.True(t, cliutil.HasShellEnv(path))

		path = cliutil.CurrentShell(false)
		assert.NotEmpty(t, path)
	}
}

func TestExecCmd(t *testing.T) {
	ret, err := cliutil.ExecCmd("echo", []string{"OK"})
	assert.NoError(t, err)
	// *nix: "OK\n" win: "OK\r\n"
	assert.Equal(t, "OK", strings.TrimSpace(ret))

	ret, err = cliutil.ExecCommand("echo", []string{"OK1"})
	assert.NoError(t, err)
	assert.Equal(t, "OK1", strings.TrimSpace(ret))

	ret, err = cliutil.QuickExec("echo OK2")
	assert.NoError(t, err)
	assert.Equal(t, "OK2", strings.TrimSpace(ret))

	ret, err = cliutil.ExecLine("echo OK3")
	assert.NoError(t, err)
	assert.Equal(t, "OK3", strings.TrimSpace(ret))
}

func TestShellExec(t *testing.T) {
	ret, err := cliutil.ShellExec("echo OK")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)

	ret, err = cliutil.ShellExec("echo OK", "bash")
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", ret)
}

func TestLineBuild(t *testing.T) {
	s := cliutil.LineBuild("myapp", []string{"-a", "val0", "arg0"})
	assert.Equal(t, "myapp -a val0 arg0", s)

	s = cliutil.BuildLine("./myapp", []string{
		"-a", "val0",
		"-m", "this is message",
		"arg0",
	})
	assert.Equal(t, `./myapp -a val0 -m "this is message" arg0`, s)
}

func TestParseLine(t *testing.T) {
	args := cliutil.ParseLine(`./app top sub -a ddd --xx "msg"`)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cliutil.String2OSArgs(`./app top sub --msg "has inner 'quote'"`)
	dump.P(args)
	assert.Len(t, args, 5)
	assert.Equal(t, "has inner 'quote'", args[4])

	// exception line string.
	args = cliutil.ParseLine(`./app top sub -a ddd --xx msg"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg", args[6])

	args = cliutil.StringToOSArgs(`./app top sub -a ddd --xx "msg "text"`)
	dump.P(args)
	assert.Len(t, args, 7)
	assert.Equal(t, "msg text", args[6])
}

func TestWorkdir(t *testing.T) {
	assert.NotEmpty(t, cliutil.Workdir())
	assert.NotEmpty(t, cliutil.BinDir())
	assert.NotEmpty(t, cliutil.BinFile())
	assert.NotEmpty(t, cliutil.BinName())
}

func TestColorPrint(t *testing.T) {
	// code gen by: kite gen parse cliutil/_demo/gen-code.tpl
	cliutil.Redp("p:red color message, ")
	cliutil.Redf("f:%s color message, ", "red")
	cliutil.Redln("ln:red color message print in cli.")
	cliutil.Bluep("p:blue color message, ")
	cliutil.Bluef("f:%s color message, ", "blue")
	cliutil.Blueln("ln:blue color message print in cli.")
	cliutil.Cyanp("p:cyan color message, ")
	cliutil.Cyanf("f:%s color message, ", "cyan")
	cliutil.Cyanln("ln:cyan color message print in cli.")
	cliutil.Grayp("p:gray color message, ")
	cliutil.Grayf("f:%s color message, ", "gray")
	cliutil.Grayln("ln:gray color message print in cli.")
	cliutil.Greenp("p:green color message, ")
	cliutil.Greenf("f:%s color message, ", "green")
	cliutil.Greenln("ln:green color message print in cli.")
	cliutil.Yellowp("p:yellow color message, ")
	cliutil.Yellowf("f:%s color message, ", "yellow")
	cliutil.Yellowln("ln:yellow color message print in cli.")
	cliutil.Magentap("p:magenta color message, ")
	cliutil.Magentaf("f:%s color message, ", "magenta")
	cliutil.Magentaln("ln:magenta color message print in cli.")

	cliutil.Infop("p:info color message, ")
	cliutil.Infof("f:%s color message, ", "info")
	cliutil.Infoln("ln:info color message print in cli.")
	cliutil.Successp("p:success color message, ")
	cliutil.Successf("f:%s color message, ", "success")
	cliutil.Successln("ln:success color message print in cli.")
	cliutil.Warnp("p:warn color message, ")
	cliutil.Warnf("f:%s color message, ", "warn")
	cliutil.Warnln("ln:warn color message print in cli.")
	cliutil.Errorp("p:error color message, ")
	cliutil.Errorf("f:%s color message, ", "error")
	cliutil.Errorln("ln:error color message print in cli.")
}
