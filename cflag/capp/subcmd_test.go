package capp_test

import (
	"bytes"
	"testing"

	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cflag/capp"
	"github.com/gookit/goutil/x/assert"
	"github.com/gookit/goutil/x/ccolor"
)

// groupOpts record the parsed options and the run calls.
type groupOpts struct {
	verbose bool
	names   cflag.Strings
	x       string
	arg1    string
	calls   []string
}

func (o *groupOpts) addCall(name string) { o.calls = append(o.calls, name) }
func (o *groupOpts) reset()              { o.calls = nil }

// newGroupApp build an app with a group command "session" and sub commands "list"(ls) and "send".
func newGroupApp() (*capp.App, *capp.Cmd, *capp.Cmd, *capp.Cmd, *groupOpts) {
	opts := &groupOpts{}
	app := capp.NewWith("myapp", "0.0.1", "this is a test app")

	group := capp.NewCmd("session", "manage sessions")
	group.BoolVar(&opts.verbose, "verbose", false, "verbose output;;v")

	sub1 := capp.NewCmd("list", "list sessions")
	sub1.Aliases = []string{"ls"}
	sub1.Var(&opts.names, "names", "the names option")
	sub1.Func = func(c *capp.Cmd) error {
		opts.addCall("list")
		return nil
	}

	sub2 := capp.NewCmd("send", "send a message")
	sub2.StringVar(&opts.x, "x", "", "the x option;;x")
	sub2.AddArg("arg1", "the arg1", false, nil)
	sub2.Func = func(c *capp.Cmd) error {
		opts.addCall("send")
		opts.arg1 = c.Arg("arg1").String()
		return nil
	}

	group.Add(sub1, sub2)
	app.Add(group)
	return app, group, sub1, sub2, opts
}

// T2: the top help still uses the exported App.NameWidth.
func TestGroupHelp_NameWidth(t *testing.T) {
	app := capp.NewWith("myapp", "0.0.1", "this is a test app")

	app.Add(capp.NewCmd("demo", "this is a demo command"))
	app.Add(capp.NewCmd("list", "this is a list command"))

	var buf bytes.Buffer
	app.SetOutput(&buf)

	// default width
	assert.NoErr(t, app.RunWithArgs([]string{"--help"}))
	helpDefault := buf.String()
	assert.StrContainsAll(t, helpDefault, []string{"Commands:", "demo", "list"})

	// custom width by the exported field
	buf.Reset()
	app.NameWidth = 24
	assert.NoErr(t, app.RunWithArgs([]string{"--help"}))
	helpWide := buf.String()

	assert.StrContainsAll(t, helpWide, []string{"Commands:", "demo", "list"})
	assert.NotEq(t, helpDefault, helpWide)

	// the render is stable on repeat calls
	buf.Reset()
	assert.NoErr(t, app.RunWithArgs([]string{"--help"}))
	assert.Eq(t, helpWide, buf.String())
}

// T3: add/find sub commands on both registration orders.
func TestSubCmd_AddAndFind(t *testing.T) {
	t.Run("add group to app at first", func(t *testing.T) {
		app := capp.NewWith("myapp", "0.0.1", "test app")
		group := capp.NewCmd("session", "manage sessions")
		app.Add(group)

		sub := capp.NewCmd("list", "list sessions")
		sub.Aliases = []string{"ls"}
		assert.NoErr(t, group.AddOrErr(sub))

		assert.True(t, group.HasSubs())
		assert.False(t, sub.HasSubs())
		assert.Eq(t, "myapp session list", sub.Path())
		assert.NotNil(t, group.CFlags.HelpFunc)

		got, ok := group.Sub("ls") // by alias
		assert.True(t, ok)
		assert.Eq(t, sub, got)

		// duplicate sub command name
		dup := capp.NewCmd("list", "another list command")
		assert.ErrSubMsg(t, group.AddOrErr(dup), "has been exists")
		assert.Panics(t, func() { group.Add(dup) })
	})

	t.Run("add sub commands to group at first", func(t *testing.T) {
		app := capp.NewWith("myapp", "0.0.1", "test app")
		group := capp.NewCmd("session", "manage sessions")
		sub := capp.NewCmd("list", "list sessions")
		group.Add(sub)
		app.Add(group)

		assert.Eq(t, "myapp session list", sub.Path())
		assert.NotNil(t, group.CFlags.HelpFunc)
	})

	t.Run("empty name and not exists", func(t *testing.T) {
		app := capp.NewWith("myapp", "0.0.1", "test app")
		group := capp.NewCmd("session", "manage sessions")
		app.Add(group)

		assert.ErrSubMsg(t, group.AddOrErr(capp.NewCmd("", "empty name")), "cannot be empty")

		_, ok := group.Sub("notExists")
		assert.False(t, ok)
	})
}

// T3: standalone Cmd use. no app, no app prefix on Path().
func TestSubCmd_StandaloneCmd(t *testing.T) {
	root := capp.NewCmd("root", "the root command")
	sub := capp.NewCmd("sub", "the sub command")
	root.Add(sub)

	assert.True(t, root.HasSubs())
	assert.False(t, sub.HasSubs())
	assert.Eq(t, root, sub.Parent())
	assert.Eq(t, "root sub", sub.Path())
	assert.Eq(t, "root", root.Path())

	got, ok := root.Sub("sub")
	assert.True(t, ok)
	assert.Eq(t, sub, got)
}

// T3: a group command cannot declare positional args.
func TestSubCmd_PosArgConflict(t *testing.T) {
	t.Run("add arg before sub commands", func(t *testing.T) {
		group := capp.NewCmd("session", "manage sessions")
		group.AddArg("arg1", "the arg1")

		err := group.AddOrErr(capp.NewCmd("list", "list sessions"))
		assert.ErrSubMsg(t, err, "cannot declare positional args")
	})

	t.Run("add sub commands before arg", func(t *testing.T) {
		group := capp.NewCmd("session", "manage sessions")
		group.Add(capp.NewCmd("list", "list sessions"))

		assert.Panics(t, func() { group.AddArg("arg1", "the arg1") })
		assert.Panics(t, func() { group.BindArg(&cflag.FlagArg{Name: "arg1"}) })
	})
}

// T3: show the group help by -h.
func TestGroupHelp_GroupHelp(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app, group, _, _, _ := newGroupApp()

	// NOTE: the command help is written to the command's own FlagSet output.
	var buf bytes.Buffer
	app.SetOutput(&buf)
	group.SetOutput(&buf)

	err := app.RunWithArgs([]string{"session", "-h"})
	assert.NoErr(t, err)

	help := buf.String()
	assert.StrContainsAll(t, help, []string{
		"Manage sessions",
		"Usage: myapp session SUBCOMMAND [--Options...] [...Arguments]",
		"Options:",
		"--verbose",
		"Commands:",
		"list, ls",
		"send",
		`Use "myapp session SUBCOMMAND --help" for about a command`,
	})
	assert.Nil(t, group.Func)
}
