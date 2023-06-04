package termctrl

// some op chars
// \x0D - Move the cursor to the beginning of the line
// \x1B[2K - Erase(Delete) the line
const (
	GotoLineStart = "\x0D"
	EraseLine     = "\x1B[2K"
)
