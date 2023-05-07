
#### FsUtil Usage

**files finder:**

```go
package main

import (
	"fmt"
	"io/fs"

	"github.com/gookit/goutil/fsutil"
)

func main() {
	// find all files in dir
	fsutil.FindInDir("./", func(filePath string, de fs.DirEntry) error {
		fmt.Println(filePath)
		return nil
	})

	// find files with filters
	fsutil.FindInDir("./", func(filePath string, de fs.DirEntry) error {
		fmt.Println(filePath)
		return nil
	}, fsutil.ExcludeDotFile)
}
```
