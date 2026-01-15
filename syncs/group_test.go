package syncs_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gookit/goutil/netutil/httpreq"
	"github.com/gookit/goutil/syncs"
	"github.com/gookit/goutil/testutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestNewErrGroup(t *testing.T) {
	httpreq.SetTimeout(3000)

	eg := syncs.NewErrGroup()
	eg.Add(func() error {
		resp, err := httpreq.Get(testSrvAddr + "/get")
		if err != nil {
			return err
		}

		fmt.Println(testutil.ParseBodyToReply(resp.Body))
		return nil
	}, func() error {
		resp := httpreq.MustResp(httpreq.Post(testSrvAddr+"/post", "hi"))
		fmt.Println(testutil.ParseBodyToReply(resp.Body))
		return nil
	})

	err := eg.Wait()
	assert.NoErr(t, err)
}
func TestErrGroup_TryGo_Success(t *testing.T) {
	t.Run("NormalExecution", func(t *testing.T) {
		group := syncs.NewErrGroup(2) // 设置限制为2个goroutine
		executed := false

		// Define function that will be executed in goroutine
		fn := func() error {
			executed = true
			return nil
		}

		// Call TryGo
		started := group.TryGo(fn)

		// Wait for goroutine to complete
		err := group.Wait()
		assert.NoError(t, err)

		// Assert results
		assert.True(t, started, "TryGo should return true when goroutine is started")
		assert.True(t, executed, "Function should be executed in goroutine")
	})

	t.Run("WithErrorReturn", func(t *testing.T) {
		group := syncs.NewErrGroup(2)
		expectedError := errors.New("test error")
		executed := false

		fn := func() error {
			executed = true
			return expectedError
		}

		started := group.TryGo(fn)
		assert.True(t, started, "TryGo should return true when goroutine is started")

		err := group.Wait()
		assert.True(t, executed, "Function should be executed")
		assert.ErrIs(t, err, expectedError, "Error should be returned as expected")
	})
}

// TestErrGroup_TryGo_LimitReached tests when goroutine limit is reached and TryGo returns false
func TestErrGroup_TryGo_LimitReached(t *testing.T) {
	// Create a group with limit of 1 goroutine
	group := syncs.NewErrGroup(1)

	// Start one goroutine that blocks indefinitely
	blockingStarted := make(chan bool, 1)
	blockingDone := make(chan bool, 1)

	blockingFn := func() error {
		blockingStarted <- true
		<-blockingDone // Wait until we're told to finish
		return nil
	}

	// Start the blocking goroutine
	firstStarted := group.TryGo(blockingFn)
	assert.True(t, firstStarted, "First goroutine should start successfully")

	// Wait for the blocking goroutine to actually start
	<-blockingStarted

	// Try to start another goroutine - this should fail due to limit
	secondStarted := group.TryGo(func() error { return nil })
	assert.False(t, secondStarted, "Second goroutine should not start due to limit")

	// Clean up by allowing blocking goroutine to finish
	blockingDone <- true

	err := group.Wait()
	assert.NoErr(t, err)
}

// TestErrGroup_TryGo_PanicRecovery tests panic recovery mechanism
func TestErrGroup_TryGo_PanicRecovery(t *testing.T) {
	group := syncs.NewErrGroup(2)
	panicOccurred := false

	fn := func() error {
		panicOccurred = true
		panic("test panic")
	}

	started := group.TryGo(fn)
	err := group.Wait()

	assert.True(t, started, "TryGo should return true even if goroutine panics")
	assert.True(t, panicOccurred, "Function should have been executed (and panicked)")
	assert.ErrMsg(t, err, "panic recover: test panic")
}
