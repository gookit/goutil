package cflag_test

import (
	"testing"

	"github.com/gookit/goutil/cliutil/cflag"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/assert"
)

func TestNewCFlags(t *testing.T) {
	opts := struct {
		int int
		str string
		bol bool
	}{}

	c := cflag.NewCFlags(cflag.WithDesc("desc for command"))
	c.IntVar(&opts.int, "int", 0, "this is a int option;;")
	c.AddArg("ag1", "this is a int option", false, "")

	inArgs := []string{"--help"}
	err := c.Parse(inArgs)
	assert.NoError(t, err)

	inArgs = []string{"--int", "23"}
	err = c.Parse(inArgs)
	assert.NoError(t, err)

	dump.P(opts)
}
