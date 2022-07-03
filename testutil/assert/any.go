package assert

// alias of interface{}
//
// TIP: cannot add `go:build !go1.18` in file head, that require the go.mod set `go 1.18`
type any = interface{}
