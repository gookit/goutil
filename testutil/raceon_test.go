//go:build race

package testutil_test

// raceEnabled is true when the tests are built with the race detector.
//
// The tests which change the process global time.Local must be skipped then: any
// goroutine that calls time.Now() (eg: an http server goroutine that is just shutting
// down) reads time.Local at the same time, and the race detector reports it.
const raceEnabled = true
