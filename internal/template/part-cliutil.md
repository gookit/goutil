
### Examples

**cmdline parse:**

```go
package main

import "github.com/gookit/goutil/cliutil"
import "github.com/gookit/goutil/dump"

func main() {
	args := cliutil.ParseLine(`./app top sub --msg "has multi words"`)
	dump.P(args)
}
```
