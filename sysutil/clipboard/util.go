package clipboard

import "strings"

// clipboard writer, reader program names
const (
	// WriterOnMac driver
	//
	// Example:
	//   echo hello | pbcopy
	//   pbcopy < tempfile.txt
	WriterOnMac = "pbcopy"

	// WriterOnWin driver on Windows
	//
	// TIP: clip only support write contents to clipboard.
	WriterOnWin = "clip"

	// WriterOnLin driver name
	//
	// linux:
	//   echo "hello-c" | xclip -selection c
	WriterOnLin = "xclip -selection clipboard"

	// ReaderOnMac driver
	//
	// Example:
	// 	Mac: pbpaste >> tasklist.txt
	ReaderOnMac = "pbpaste"

	// ReaderOnWin driver on Windows
	//
	// read clipboard should use: powershell get-clipboard
	ReaderOnWin = "powershell get-clipboard"

	// ReaderOnLin driver name
	//
	// Usage:
	// 	xclip -o -selection clipboard
	// 	xclip -o -selection c // can use shorts
	ReaderOnLin = "xclip -o -selection clipboard"
)

var (
	writerOnLin = []string{"xclip", "xsel"}
)

// std instance
var std = New()

// Std get
func Std() *Clipboard {
	return std
}

// Reset clipboard data
func Reset() error {
	return std.Reset()
}

// Available clipboard available check
func Available() bool {
	return std.Available()
}

// ReadString contents from clipboard
func ReadString() (string, error) {
	return std.ReadString()
}

// WriteString contents to clipboard and flush
func WriteString(s string) error {
	if _, err := std.WriteString(s); err != nil {
		return err
	}
	return std.Flush()
}

// special handle on with args
func parseLine(line string) (bin string, args []string) {
	bin = line
	if strings.ContainsRune(line, ' ') {
		list := strings.Split(line, " ")
		bin, args = list[0], list[1:]
	}
	return
}
