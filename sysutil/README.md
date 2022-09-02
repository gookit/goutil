# System Utils

Provide some system util functions. eg: sysenv, exec, user, process

- quick exec a command line string.
- quick build a command for run

## Install

```bash
go get github.com/gookit/goutil/sysutil
```

## Usage

```go
out, err := sysutil.ExecCmd("ls", []string{"-al"})
```

## Clipboard

Package `clipboard` provide a simple clipboard read and write operations.

### Install

```bash
go get github.com/gookit/goutil/sysutil/clipboard
```

### Usage

Examples:

```go
src := "hello, this is clipboard"
err = clipboard.WriteString(src)
assert.NoErr(t, err)

// str: "hello, this is clipboard"
str, err = clipboard.ReadString()
assert.NoErr(t, err)
assert.NotEmpty(t, str)
assert.Eq(t, src, str)
```

## Cmd run

### Install

```bash
go get github.com/gookit/goutil/sysutil/cmdr
```
### Usage

**Cmd Builder**:

```go
c := cmdr.NewCmd("ls").
    WithArg("-l").
    WithArgs([]string{"-h"}).
    AddArg("-a").
    AddArgf("%s", "./")

c.OnBefore(func(c *cmdr.Cmd) {
    assert.Eq(t, "ls -l -h -a ./", c.Cmdline())
})

out := c.SafeOutput()
fmt.Println(out)
```

**Batch Cmd Tasks**:

Can use `cmdr.Runner` run multi cmd tasks at once.

```go
buf := new(bytes.Buffer)
rr := cmdr.NewRunner()

rr.Add(&cmdr.Task{
    ID:  "task1",
    Cmd: cmdr.NewCmd("id", "-F").WithOutput(buf, buf),
})
rr.AddCmd(cmdr.NewCmd("ls").AddArgs([]string{"-l", "-h"}).WithOutput(buf, buf))

err = rr.Run()
```
