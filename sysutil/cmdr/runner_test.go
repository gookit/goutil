package cmdr_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gookit/goutil/sysutil/cmdr"
	"github.com/gookit/goutil/testutil/assert"
)

func TestRunner_Run(t *testing.T) {
	buf := new(bytes.Buffer)
	rr := cmdr.NewRunner()

	rr.Add(&cmdr.Task{
		ID:  "task1",
		Cmd: cmdr.NewCmd("id", "-F").WithOutput(buf, buf),
	})
	rr.AddCmd(cmdr.NewCmd("ls").AddArgs([]string{"-l", "-h"}).WithOutput(buf, buf))

	task, err := rr.Task("task1")
	assert.NoErr(t, err)

	ids := rr.TaskIDs()
	// dump.P(rr.TaskIDs())
	assert.Len(t, ids, 2)
	assert.NotEmpty(t, ids)
	assert.Contains(t, ids, task.ID)

	err = rr.Run()
	assert.NoErr(t, err)

	fmt.Println(buf.String())
}
