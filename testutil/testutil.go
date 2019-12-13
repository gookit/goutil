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
	_ = os.Stdout.Close()
	// restore
	os.Stdout = oldStdout
	oldStdout = nil

	// read output data
	if newReader != nil {
		out, _ := ioutil.ReadAll(newReader)
		s = string(out)

		// close reader
		_ = newReader.Close()
		newReader = nil
	}
	return
}

// MockEnvValue will store old env value, set new val. will restore old value on end.
func MockEnvValue(key, val string, fn func(nv string)) {
	old := os.Getenv(key)
	err := os.Setenv(key, val)
	if err != nil {
		panic(err)
	}

	fn(os.Getenv(key))

	// if old is empty, unset key.
	if old == "" {
		err = os.Unsetenv(key)
	} else {
		err = os.Setenv(key, old)
	}
	if err != nil {
		panic(err)
	}
}

// MockEnvValues will store old env value, set new val. will restore old value on end.
func MockEnvValues(kvMap map[string]string, fn func()) {
	backups := make(map[string]string, len(kvMap))

	for key, val := range kvMap {
		backups[key] = os.Getenv(key)
		_ = os.Setenv(key, val)
	}

	fn()

	for key := range kvMap {
		if old := backups[key]; old == "" {
			_ = os.Unsetenv(key)
		} else {
			_ = os.Setenv(key, old)
		}
	}
}
