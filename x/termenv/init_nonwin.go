//go:build !windows

package termenv

func init() {
	// on non-Windows: the needVTP value is always false, so ignore it.
	colorLevel, _ = detectTermColorLevel()
}
