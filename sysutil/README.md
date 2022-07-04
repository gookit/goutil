# System Utils

- quick exec a command line string.

## Install

```bash
go get github.com/gookit/goutil/sysutil
```

## Usage

```go
sysutil.ExecCmd("ls", []string{"-al"})
```

## Clipboard

Package `clipboard` provide a simple clipboard read and write operations.

```bash
go get github.com/gookit/goutil/sysutil/clipboard
```

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


