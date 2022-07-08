package main

import (
	"flag"
	"os"

	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
)

var opts1 = struct {
	age  int
	name string
	str1 string
	lOpt string
}{}

// go run ./_example/rawflag.go
// go run ./cflag/_example/rawflag.go -h
func main() {
	c := flag.NewFlagSet("mycmd", flag.ContinueOnError)

	c.IntVar(&opts1.age, "age", 0, "this is a int option")
	c.StringVar(&opts1.name, "name", "", "this is a string option and required")
	c.StringVar(&opts1.str1, "str1", "def-val", "this is a string option with default value")
	c.StringVar(&opts1.lOpt, "long-opt", "", "this is a string option with shorts")

	err := c.Parse(os.Args[1:])
	if err != nil {
		if err != flag.ErrHelp {
			cliutil.Errorln("Error:", err.Error())
		}
		return
	}

	// after parse, do something
	handleFunc1()
}

func handleFunc1() {
	cliutil.Infoln("after parse, do something")

	dump.P(opts1)
}
