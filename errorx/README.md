# ErrorX

`errorx` provide an enhanced error implements for go, allow with stacktraces and wrap another error.

## Usage


### Create error with call stack info

- use the `errorx.New` instead `errors.New`

```go
func doSomething() error {
    if false {
	    // return errors.New("a error happen")
	    return errorx.New("a error happen")
	}
}
```

- use the `errorx.Newf` or `errorx.Errorf` instead `fmt.Errorf`

```go
func doSomething() error {
    if false {
	    // return fmt.Errorf("a error %s", "happen")
	    return errorx.Newf("a error %s", "happen")
	}
}
```

### Wrap the previous error

used like this before:

```go
    if err := SomeFunc(); err != nil {
	    return err
	}
```

can be replaced with:

```go
    if err := SomeFunc(); err != nil {
	    return errors.Stacked(err)
	}
```

## Refers

- golang errors
- https://github.com/joomcode/errorx
- https://github.com/pkg/errors
- https://github.com/juju/errors
- https://github.com/go-errors/errors
