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
	gName   cflag.Strings
	x       string
	arg1    string
	calls   []string
}

func (o *groupOpts) addCall(name string) { o.calls = append(o.calls, name) }

func (o *groupOpts) reset() {
	o.calls = nil
	o.names = nil
	o.gName = nil
	o.x, o.arg1 = "", ""
	o.verbose = false
}

// newGroupApp build an app with a group command "session" and sub commands "list"(ls) and "send".
func newGroupApp() (*capp.App, *capp.Cmd, *capp.Cmd, *capp.Cmd, *groupOpts) {
	opts := &groupOpts{}
	app := capp.NewWith("myapp", "0.0.1", "this is a test app")

	group := capp.NewCmd("session", "manage sessions")
	group.BoolVar(&opts.verbose, "verbose", false, "verbose output;;v")
	group.Var(&opts.gName, "gname", "the group names option")

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

// Cmd created by struct literal: no CFlags and no Func.
func TestSubCmd_CmdWithoutCFlags(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app := capp.NewWith("myapp", "0.0.1", "test app")

	demo := &capp.Cmd{Name: "demo"} // no Desc, no CFlags
	other := &capp.Cmd{Name: "other", Desc: "this is other command"}
	app.Add(demo, other)

	var buf bytes.Buffer
	app.SetOutput(&buf)

	// render the help should not panic on the nil CFlags
	assert.NoErr(t, app.RunWithArgs([]string{"--help"}))
	assert.StrContains(t, buf.String(), "This is other command")

	// parse on a Cmd without CFlags should not panic
	assert.NotPanics(t, func() { demo.MustRun([]string{}) })
}

// The command help follows the app output on the command output not set.
func TestSubCmd_HelpFollowsAppOutput(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app, group, _, _, _ := newGroupApp()

	// the command output is not set: it follows the app output
	var appBuf bytes.Buffer
	app.SetOutput(&appBuf)

	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	assert.StrContains(t, appBuf.String(), "Commands:")

	// an explicitly set command output is kept
	cmdBuf := new(bytes.Buffer)
	group.SetOutput(cmdBuf)
	appBuf.Reset()

	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	assert.Eq(t, "", appBuf.String())
	assert.StrContains(t, cmdBuf.String(), "Commands:")
}

// T6 (A2): the group command parses its own options only once.
func TestSubCmd_RepeatableOptOnce(t *testing.T) {
	app, group, _, _, opts := newGroupApp()

	var ran bool
	group.Func = func(c *capp.Cmd) error {
		ran = true
		return nil
	}

	opts.reset()
	assert.NoErr(t, app.RunWithArgs([]string{"session", "--gname", "tom"}))
	assert.True(t, ran)
	assert.Eq(t, []string{"tom"}, opts.gName.Strings())
}

// T6 (A3): AfterFlagParse only fires once on a group command.
func TestSubCmd_AfterFlagParseOnce(t *testing.T) {
	app, group, _, _, _ := newGroupApp()

	count := 0
	group.AfterFlagParse = func(c *cflag.CFlags) bool {
		count++
		return true
	}

	var ran bool
	group.Func = func(c *capp.Cmd) error {
		ran = true
		return nil
	}

	// the group runs itself
	assert.NoErr(t, app.RunWithArgs([]string{"session"}))
	assert.True(t, ran)
	assert.Eq(t, 1, count)

	// descend into a sub command
	count = 0
	assert.NoErr(t, app.RunWithArgs([]string{"session", "list"}))
	assert.Eq(t, 1, count)
}

// T6 (A1/D10): a group command cannot declare positional args, so the sub command
// name can never be bound as an argument value.
func TestSubCmd_GroupWithPosArg(t *testing.T) {
	group := capp.NewCmd("session", "manage sessions")
	group.AddArg("arg1", "the arg1")

	assert.ErrSubMsg(
		t,
		group.AddOrErr(capp.NewCmd("list", "list sessions")),
		"cannot declare positional args",
	)
	assert.False(t, group.HasSubs())
}

// T6 (A4/P7): the group help works on the reverse registration order too.
func TestSubCmd_ReverseOrderGroupHelp(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app := capp.NewWith("myapp", "0.0.1", "test app")
	group := capp.NewCmd("session", "manage sessions")
	sub := capp.NewCmd("list", "list sessions")
	// add the sub command to the group at first
	group.Add(sub)
	app.Add(group)

	assert.Eq(t, "myapp session list", sub.Path())

	var buf bytes.Buffer
	group.SetOutput(&buf)

	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	help := buf.String()
	assert.StrContainsAll(t, help, []string{
		"Usage: myapp session SUBCOMMAND",
		"Commands:",
		"list",
	})
}

// T6 (D2): the group short option and the sub command short option can coexist.
func TestSubCmd_ShortsNotEaten(t *testing.T) {
	app, _, _, _, opts := newGroupApp()

	opts.reset()
	assert.NoErr(t, app.RunWithArgs([]string{"session", "-v", "send", "-x", "val", "a1"}))
	assert.True(t, opts.verbose)
	assert.Eq(t, "val", opts.x)
	assert.Eq(t, "a1", opts.arg1)
}

// T6 (S6): a missing required group option errors before the group help.
func TestSubCmd_RequiredGroupOpt(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app := capp.NewWith("myapp", "0.0.1", "test app")
	group := capp.NewCmd("session", "manage sessions")

	var req string
	group.StringVar(&req, "req", "", "the required option;true")
	group.Add(capp.NewCmd("list", "list sessions"))
	app.Add(group)

	var buf bytes.Buffer
	group.SetOutput(&buf)

	// required option not given
	err := app.RunWithArgs([]string{"session"})
	assert.ErrSubMsg(t, err, "flag option 'req' is required")

	// -h first: show the group help
	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	assert.StrContains(t, buf.String(), "Commands:")
}
func TestGroupHelp_Entries(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app, group, _, _, _ := newGroupApp()

	var buf bytes.Buffer
	app.SetOutput(&buf)
	group.SetOutput(&buf)

	// 1) by -h
	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	helpByFlag := buf.String()

	// 2) by the help keyword
	buf.Reset()
	assert.NoErr(t, app.RunWithArgs([]string{"session", "help"}))
	helpByKeyword := buf.String()

	// 3) by no positional args
	buf.Reset()
	assert.NoErr(t, app.RunWithArgs([]string{"session"}))
	helpByEmpty := buf.String()

	assert.Eq(t, helpByFlag, helpByKeyword)
	assert.Eq(t, helpByFlag, helpByEmpty)

	// with the LongHelp and Example
	group.LongHelp = "this is a long help text"
	group.Example = "myapp session list"

	buf.Reset()
	assert.NoErr(t, app.RunWithArgs([]string{"session", "-h"}))
	help := buf.String()
	assert.StrContainsAll(t, help, []string{"Help:", "this is a long help text", "Examples:", "myapp session list"})
}

// T5: the leaf command help is not affected.
func TestGroupHelp_LeafHelp(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app, _, sub1, _, _ := newGroupApp()

	var buf bytes.Buffer
	app.SetOutput(&buf)
	sub1.SetOutput(&buf)

	assert.NoErr(t, app.RunWithArgs([]string{"session", "list", "-h"}))

	help := buf.String()
	assert.StrContainsAll(t, help, []string{
		"List sessions (alias: ls)",
		"Usage: list [--Options...] [...Arguments]",
		"--names",
		"--help, -h",
	})
	assert.StrNotContains(t, help, "SUBCOMMAND")
}

// T5: the top help lists the group command as a normal command.
func TestGroupHelp_TopHelp(t *testing.T) {
	ccolor.Disable()
	defer ccolor.RevertColorSupport()

	app, _, _, _, _ := newGroupApp()

	var buf bytes.Buffer
	app.SetOutput(&buf)

	assert.NoErr(t, app.RunWithArgs([]string{"--help"}))

	help := buf.String()
	assert.StrContainsAll(t, help, []string{
		"Commands:",
		"session",
		"Manage sessions",
		`Use "myapp COMMAND --help" for about a command`,
	})
	// the sub commands are not listed on the top help
	assert.StrNotContains(t, help, "list")
	assert.StrNotContains(t, help, "SUBCOMMAND")
}
func TestSubCmd_Run(t *testing.T) {
	app, _, _, _, opts := newGroupApp()

	t.Run("run sub command by name", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "list", "--names", "tom"}))
		assert.Eq(t, []string{"list"}, opts.calls)
		assert.Eq(t, []string{"tom"}, opts.names.Strings())
	})

	t.Run("run sub command by alias", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "ls"}))
		assert.Eq(t, []string{"list"}, opts.calls)
	})

	t.Run("group option before the sub command name", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "-v", "list"}))
		assert.True(t, opts.verbose)
		assert.Eq(t, []string{"list"}, opts.calls)
	})

	t.Run("group option after the sub command name", func(t *testing.T) {
		opts.reset()
		err := app.RunWithArgs([]string{"session", "list", "-v"})
		assert.ErrSubMsg(t, err, "flag provided but not defined: -v")
	})

	t.Run("sub command option and argument", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "send", "-x", "val", "the-arg"}))
		assert.Eq(t, []string{"send"}, opts.calls)
		assert.Eq(t, "val", opts.x)
		assert.Eq(t, "the-arg", opts.arg1)
	})

	t.Run("sub command short option with value", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "send", "-x=val2", "a2"}))
		assert.Eq(t, "val2", opts.x)
		assert.Eq(t, "a2", opts.arg1)
	})

	t.Run("not exists sub command", func(t *testing.T) {
		err := app.RunWithArgs([]string{"session", "notExists"})
		assert.ErrMsg(t, err, `command "notExists" not exists in "myapp session"`)
	})

	t.Run("-- on the app args", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"--", "session", "list"}))
		assert.Eq(t, []string{"list"}, opts.calls)
	})

	t.Run("-- on the group args", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "--", "list"}))
		assert.Eq(t, []string{"list"}, opts.calls)
	})

	t.Run("-- on the leaf command args", func(t *testing.T) {
		opts.reset()
		assert.NoErr(t, app.RunWithArgs([]string{"session", "send", "the-arg", "--", "-x", "val"}))
		assert.Eq(t, "the-arg", opts.arg1)
		assert.Eq(t, "", opts.x) // -x is kept as a raw argument
	})

	t.Run("three level dispatch", func(t *testing.T) {
		l3App := capp.NewWith("myapp", "0.0.1", "test app")
		g1 := capp.NewCmd("g1", "group 1")
		g2 := capp.NewCmd("g2", "group 2")

		var ran string
		leaf := capp.NewCmd("leaf", "the leaf command", func(c *capp.Cmd) error {
			ran = "g1 g2 leaf"
			return nil
		})
		g2.Add(leaf)
		g1.Add(g2)
		l3App.Add(g1)

		assert.NoErr(t, l3App.RunWithArgs([]string{"g1", "g2", "leaf"}))
		assert.Eq(t, "g1 g2 leaf", ran)
		assert.Eq(t, "myapp g1 g2 leaf", leaf.Path())
	})
}

// T4: the app hooks only fire on the command that really runs.
func TestSubCmd_Hooks(t *testing.T) {
	app, group, _, _, opts := newGroupApp()

	var beforeCalls, afterCalls []string
	var beforeArgs []string
	app.BeforeRun = func(c *capp.Cmd, cmdArgs []string) bool {
		beforeCalls = append(beforeCalls, c.Name)
		beforeArgs = cmdArgs
		return true
	}
	app.AfterRun = func(c *capp.Cmd, err error) {
		afterCalls = append(afterCalls, c.Name)
	}

	t.Run("leaf command", func(t *testing.T) {
		beforeCalls, afterCalls = nil, nil
		opts.reset()

		assert.NoErr(t, app.RunWithArgs([]string{"session", "list", "--names", "tom"}))
		assert.Eq(t, []string{"list"}, beforeCalls)
		assert.Eq(t, []string{"list"}, afterCalls)
		assert.Eq(t, []string{"--names", "tom"}, beforeArgs)
	})

	t.Run("descend into a sub command", func(t *testing.T) {
		beforeCalls, afterCalls = nil, nil
		opts.reset()

		assert.NoErr(t, app.RunWithArgs([]string{"session", "ls"}))
		assert.Eq(t, []string{"list"}, beforeCalls)
		assert.Eq(t, []string{"list"}, afterCalls)
	})

	t.Run("group command runs itself", func(t *testing.T) {
		var ran bool
		group.Func = func(c *capp.Cmd) error {
			ran = true
			return nil
		}
		defer func() { group.Func = nil }()

		beforeCalls, afterCalls = nil, nil
		assert.NoErr(t, app.RunWithArgs([]string{"session"}))
		assert.True(t, ran)
		assert.Eq(t, []string{"session"}, beforeCalls)
		assert.Eq(t, []string{"session"}, afterCalls)
	})
}
