//go:build !race

package testutil_test

// raceEnabled is false when the tests are built without the race detector.
const raceEnabled = false
