
#### Usage

**创建错误带有调用栈信息**

- 使用 `errorx.New` 替代 `errors.New`

```go
func doSomething() error {
    if false {
	    // return errors.New("a error happen")
	    return errorx.New("a error happen")
	}
}
```

- 使用 `errorx.Newf` 或者 `errorx.Errorf` 替代 `fmt.Errorf`

```go
func doSomething() error {
    if false {
	    // return fmt.Errorf("a error %s", "happen")
	    return errorx.Newf("a error %s", "happen")
	}
}
```

**包装上一级错误**

之前这样使用:

```go
    if err := SomeFunc(); err != nil {
	    return err
	}
```

可以替换成:

```go
    if err := SomeFunc(); err != nil {
	    return errors.Stacked(err)
	}
```
