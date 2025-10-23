package syncs_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/gookit/goutil/syncs"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

var testSrvAddr string

func TestMain(m *testing.M) {
	s := testutil.NewEchoServer()
	defer s.Close()
	testSrvAddr = "http://" + s.Listener.Addr().String()
	fmt.Println("Test server listen on:", testSrvAddr)

	m.Run()
}

func TestWaitGroupGo(t *testing.T) {
	var wg syncs.WaitGroup
	var counter int32
	expected := int32(5)

	// 启动多个 goroutines
	for i := int32(0); i < expected; i++ {
		wg.Go(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	// 等待所有 goroutines 完成
	wg.Wait()
	assert.Eq(t, expected, counter)
}

func TestContextValue(t *testing.T) {
	assert.NotEmpty(t, syncs.ContextValue("key", "value"))
}
