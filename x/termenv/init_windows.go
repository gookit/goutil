package termenv

func init() {
	// needVTP=false OR noColor=true: Don't need to enable virtual process
	if !needVTP || noColor {
		return
	}

	// try force enable colors on Windows terminal
	TryEnableVTP(needVTP)
}
