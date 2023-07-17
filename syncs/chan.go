package syncs

import (
	"os"
	"os/signal"
	"syscall"
)

// WaitCloseSignals for some huang program.
//
// Usage:
//
//	// do something. eg: start a http server
//
//	syncs.WaitCloseSignals(func(sig os.Signal) {
//		// do something
//	})
func WaitCloseSignals(onClose func(sig os.Signal)) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	// block until a signal is received.
	onClose(<-signals)
}

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
func Go(f func() error) error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()

	return <-ch
}
