
#### FsUtil Usage

**files finder:**

```go
package main

import (
	"fmt"
	"os"

	"github.com/gookit/goutil/fsutil/finder"
)

func main() {
	f := finder.EmptyFinder()

	f.
		AddDir("./testdata").
		AddFile("finder.go").
		NoDotFile().
		// NoDotDir().
		Find().
		Each(func(filePath string) {
			fmt.Println(filePath)
		})

	finder.NewFinder([]string{"./testdata"}).
		AddFile("finder.go").
		NoDotDir().
		EachStat(func(fi os.FileInfo, filePath string) {
			fmt.Println(filePath, "=>", fi.ModTime())
		})
}
```
