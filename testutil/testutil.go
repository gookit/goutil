// Package testutil provide some test help util functions. eg: http test, mock ENV value
package testutil

import (
	"io"
	"os"
)

var oldStdout, oldStderr, newReader *os.File

// DiscardStdout Discard os.Stdout output
//
// Usage:
//
//	DiscardStdout()
//	fmt.Println("Hello, playground")
//	RestoreStdout()
func DiscardStdout() error {
	// save old os.Stdout
	oldStdout = os.Stdout

	stdout, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err == nil {
		os.Stdout = stdout
	}
	return err
}

// ReadOutput restore os.Stdout
// func ReadOutput() (s string) {
// }

// RewriteStdout rewrite os.Stdout
//
// Usage:
//
//	RewriteStdout()
//	fmt.Println("Hello, playground")
//	msg := RestoreStdout()
func RewriteStdout() {
	if oldStdout != nil {
		return
	}

	oldStdout = os.Stdout
	r, w, _ := os.Pipe()
	newReader = r
	os.Stdout = w
}

// RestoreStdout restore os.Stdout
func RestoreStdout(printData ...bool) (s string) {
	if oldStdout == nil {
		return
	}

	// Notice: must close writer before read data
	// close now reader
	_ = os.Stdout.Close()
	// restore
	os.Stdout = oldStdout
	oldStdout = nil
	if newReader == nil {
		return
	}

	// read output data
	out, _ := io.ReadAll(newReader)
	s = string(out)

	// print the read data to stdout
	if len(printData) > 0 && printData[0] {
		_, _ = os.Stdout.WriteString(s)
	}

	// close reader
	_ = newReader.Close()
	newReader = nil
	return
}

// RewriteStderr rewrite os.Stderr
//
// Usage:
//
//	RewriteStderr()
//	fmt.Fprintln(os.Stderr, "Hello, playground")
//	msg := RestoreStderr()
func RewriteStderr() {
	if oldStderr != nil {
		return
	}

	oldStderr = os.Stderr
	r, w, _ := os.Pipe()
	newReader = r
	os.Stderr = w
}

// RestoreStderr restore os.Stderr
func RestoreStderr(printData ...bool) (s string) {
	if oldStderr == nil {
		return
	}

	// Notice: must close writer before read data
	// close now reader
	_ = os.Stderr.Close()
	// restore
	os.Stderr = oldStderr
	oldStderr = nil
	if newReader == nil {
		return
	}

	// read output data
	bts, _ := io.ReadAll(newReader)
	s = string(bts)

	// print the read data to stderr
	if len(printData) > 0 && printData[0] {
		_, _ = os.Stderr.WriteString(s)
	}

	// close reader
	_ = newReader.Close()
	newReader = nil
	return
}
