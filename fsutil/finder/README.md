# finder

[![GoDoc](https://godoc.org/github.com/goutil/fsutil/finder?status.svg)](https://godoc.org/github.com/goutil/fsutil/finder)

`finder` provide a finder tool for find files, dirs.

## Usage

```go
package main

import (
	"github.com/gookit/goutil/dump"
	"github.com/goutil/fsutil/finder"
)

func main() {
	ff := finder.NewFinder()
	ff.AddPath("/tmp")
	ff.AddPath("/usr/local")
	ff.AddPath("/usr/local/bin")
	ff.AddPath("/usr/local/lib")
	ff.AddPath("/usr/local/libexec")
	ff.AddPath("/usr/local/sbin")
	ff.AddPath("/usr/local/share")

	ss := ff.FindPaths()
	dump.P(ss)
}
```

