package syncs_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/gookit/goutil/syncs"
	"github.com/gookit/goutil/testutil/assert"
)

func TestWaitCloseSignals(t *testing.T) {
	sgCh := make(chan os.Signal, 1)

	// delay to send signal
	go func() {
		time.Sleep(10 * time.Millisecond)
		sgCh <- os.Interrupt
	}()

	syncs.WaitCloseSignals(func(sig os.Signal) {
		assert.Eq(t, "interrupt", sig.String())
	}, sgCh)

	// sgCh has closed
	assert.Panics(t, func() {
		sgCh <- os.Interrupt
	})
}

func TestSignalHandler_ContextCancelled_ReturnsContextError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	execute, _ := syncs.SignalHandler(ctx, os.Interrupt)

	// 模拟取消上下文
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel() // 直接取消上下文而不是调用 interrupt(nil)
	}()

	err := execute()
	assert.ErrIs(t, err, context.Canceled)
}

func TestSignalHandler_NoSignalOrCancel_Blocks(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	execute, _ := syncs.SignalHandler(ctx, os.Interrupt)

	// 检查是否阻塞
	done := make(chan struct{})
	go func() {
		_ = execute()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Expected execute to block")
	case <-time.After(200 * time.Millisecond):
		// 期望的行为
	}
}
