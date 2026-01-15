package syncs

import "fmt"

// Go is a basic promise implementation: it wraps calls a function in a goroutine
// and returns a channel which will later return the function's return value.
//
// if panic happen, it will be recovered and return as error
func Go(f func() error) error {
	ch := make(chan error)
	go func() {
		// add recovery handle
		defer func() {
			if r := recover(); r != nil {
				ch <- fmt.Errorf("panic recover: %v", r)
			}
		}()

		ch <- f()
	}()
	return <-ch
}
