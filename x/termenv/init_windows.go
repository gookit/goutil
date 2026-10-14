package termenv

// needVTP - need enable VTP(VirtualTerminalProcessing). only for Windows OS
var needVTP bool

func init() {
	colorLevel, needVTP = detectTermColorLevel()

	// needVTP=false OR noColor=true: Don't need to enable virtual process
	if !needVTP || noColor {
		return
	}

	// try force enable colors on Windows terminal
	TryEnableVTP(needVTP)
}
