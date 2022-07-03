package main

import (
	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/cliutil"
)

var opts = struct {
	tplFile string
}{}

// go run ./internal/gencode -h
// go run ./internal/gencode
func main() {
	cmd := cflag.New(func(c *cflag.CFlags) {
		c.Version = "0.1.2"
		c.Desc = "auto generate some codes for goutil"
	})

	cmd.StringVar(&opts.tplFile, "tpl-file", "", "the template file path;true;tpl")

	cmd.Func = handle
	cmd.QuickRun()
}

func handle(c *cflag.CFlags) error {
	cliutil.Infoln("TODO")
	return nil
}
