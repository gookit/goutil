package syncs_test

import (
	"fmt"
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
	counter := 0
	expected := 5

	// 启动多个 goroutines
	for i := 0; i < expected; i++ {
		wg.Go(func() {
			counter++
		})
	}

	// 等待所有 goroutines 完成
	wg.Wait()
	assert.Eq(t, expected, counter)
}
