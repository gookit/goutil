# Dump

`goutil/dump` is a golang data printing toolkit that prints beautiful and easy to read go slice, map, struct data

- Github: https://github.com/gookit/goutil/dump
- GoDoc: https://pkg.go.dev/github.com/gookit/goutil/dump

## Install

```bash
go get github.com/gookit/goutil/dump
```

## Usage

run demo: `go run ./dump/_examples/demo1.go`

```go
package main

import "github.com/gookit/goutil/dump"

// rum demo: go run ./dump/_examples/demo1.go
func main() {
	otherFunc1()
}

func otherFunc1() {
	dump.P(
		23,
		[]string{"ab", "cd"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, // len > 10
		map[string]interface{}{
			"key": "val", "sub": map[string]string{"k": "v"},
		},
		struct {
			ab string
			Cd int
		}{
			"ab", 23,
		},
	)
}
```

You will see:

![](_examples/preview-demo1.png)

## More preview

- nested struct

![](_examples/preview-nested-struct.png)


## Functions

```go
func P(vs ...interface{})
func V(vs ...interface{})
func Print(vs ...interface{})
```

## Related

- https://github.com/kr/pretty
- https://github.com/davecgh/go-spew More detail for kr/pretty
  - https://github.com/kortschak/utter It's forks