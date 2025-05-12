//go:build !windows

package termenv

// detect special term color support
func detectSpecialTermColor(termVal string) (Level, bool) {
	if termVal == "" {
		// on Windows WSL:
		// - runtime.GOOS == "Linux"
		// - support true-color
		// ENV:
		// 	WSL_DISTRO_NAME=Debian
		if val := os.Getenv("WSL_DISTRO_NAME"); val != "" {
			// detect WSL as it has True Color support
			if detectWSL() {
				debugf("True Color support on WSL environment")
				return TermColorTrue, false
			}
		}
		return TermColorNone, false
	}

	debugf("terminfo check - fallback detect color by check TERM value")

	// on TERM=screen:
	// - support 256, not support true-color. test on macOS
	if termVal == noTrueColorTerm {
		return TermColor256, false
	}

	if strings.Contains(termVal, "256color") {
		return TermColor256, false
	}

	if strings.Contains(termVal, "xterm") {
		return TermColor256, false
	}
	return TermColor16, false
}
