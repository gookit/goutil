package main

import (
	"fmt"

	"github.com/gookit/goutil/cflag/capp"
)

var groupOpts = struct {
	verbose bool
	format  string
	limit   int
	message string
}{}

// go run ./cflag/_example/group.go -h
// go run ./cflag/_example/group.go session -h
// go run ./cflag/_example/group.go session -v list -f json
// go run ./cflag/_example/group.go session send -m "hello" s3
func main() {
	app := capp.NewApp()
	app.Name = "myapp"
	app.Desc = "this is a multi-level command app"
	app.Version = "0.0.1"

	// the group command: session
	group := capp.NewCmd("session", "manage sessions")
	group.BoolVar(&groupOpts.verbose, "verbose", false, "show verbose output;;v")

	// sub command: session list. with alias: ls
	list := capp.NewCmd("list", "list all sessions")
	// NOTE: the aliases should be set before add.
	list.Aliases = []string{"ls"}
	list.StringVar(&groupOpts.format, "format", "text", "output format;;f")
	list.IntVar(&groupOpts.limit, "limit", 10, "max number of sessions;;l")
	list.Func = func(c *capp.Cmd) error {
		fmt.Println("run: session list")
		fmt.Println("verbose:", groupOpts.verbose, "format:", groupOpts.format, "limit:", groupOpts.limit)
		return nil
	}

	// sub command: session send
	send := capp.NewCmd("send", "send a message to a session")
	send.StringVar(&groupOpts.message, "message", "", "the message to send;;m")
	send.AddArg("session-id", "the target session id", true, nil)
	send.Func = func(c *capp.Cmd) error {
		fmt.Println("run: session send")
		fmt.Println("message:", groupOpts.message, "session-id:", c.Arg("session-id").String())
		return nil
	}

	// add the sub commands to the group, then add the group to the app.
	group.Add(list, send)
	app.Add(group)
	app.Run()
}
