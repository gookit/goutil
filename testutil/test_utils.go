package testutil

import (
	"io/ioutil"
	"os"
)

var oldStdout, newReader *os.File

// DiscardStdout Discard os.Stdout output
// Usage:
// 	DiscardStdout()
// 	fmt.Println("Hello, playground")
// 	RestoreStdout()
func DiscardStdout() error {
	// save old os.Stdout
	oldStdout = os.Stdout

	stdout, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err == nil {
		os.Stdout = stdout
	}

	return err
}

// RewriteStdout rewrite os.Stdout
// Usage:
// 	RewriteStdout()
// 	fmt.Println("Hello, playground")
// 	msg := RestoreStdout()
func RewriteStdout() {
	oldStdout = os.Stdout
	r, w, _ := os.Pipe()
	newReader = r
	os.Stdout = w
}

// ReadOutput restore os.Stdout
// func ReadOutput() (s string) {
// }

// RestoreStdout restore os.Stdout
func RestoreStdout() (s string) {
	if oldStdout == nil {
		return
	}

	// Notice: must close writer before read data
	// close now reader
	os.Stdout.Close()
	// restore
	os.Stdout = oldStdout
	oldStdout = nil

	// read output data
	if newReader != nil {
		out, _ := ioutil.ReadAll(newReader)
		s = string(out)

		// close reader
		newReader.Close()
		newReader = nil
	}

	return
}
