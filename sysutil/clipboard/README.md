# Clipboard

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

## Related

- https://github.com/zyedidia/clipper
