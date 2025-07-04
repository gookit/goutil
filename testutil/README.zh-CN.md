# Test Utils


## Install

```bash
go get github.com/gookit/goutil/testutil
```

## [`assert`](./assert) tests

```go
package assert_test

import (
	"testing"

	"github.com/gookit/goutil/errorx"
	"github.com/gookit/goutil/testutil/assert"
)

func TestErr(t *testing.T) {
	err := errorx.Raw("this is a error")

	assert.NoErr(t, err, "user custom message")
	assert.ErrMsg(t, err, "this is a error")
}
```

Run tests for special method:

```shell
go test -v -run ^TestErr$
go test -v -run ^TestErr$ ./testutil/assert/...
```

**Error on fail**:

![test-err](_example/test-err.png)

## 单元测试MOCK

### 使用 echo server 测试HTTP

使用 `testutil.NewEchoServer()` 可以快速的创建一个HTTP echo server. 方便测试HTTP请求，响应等。

```go
var testSrvAddr string

func TestMain(m *testing.M) {
    s := testutil.NewEchoServer()
    defer s.Close()

    testSrvAddr = s.PrintHttpHost()
    m.Run()
}

func TestNewEchoServer(t *testing.T) {
    // 可直接请求测试server
    r, err := http.Post(testSrvAddr, "text/plain", strings.NewReader("hello!"))
    assert.NoErr(t, err)

    // 将响应信息绑定到 testutil.EchoReply
    rr := testutil.ParseRespToReply(r)
    dump.P(rr)
    assert.Eq(t, "POST", rr.Method)
    assert.Eq(t, "text/plain", rr.ContentType())
    assert.Eq(t, "hello!", rr.Body)
}
```