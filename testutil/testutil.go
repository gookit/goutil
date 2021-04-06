package testutil

import (
	"io/ioutil"
	"os"
	"strings"
)

var oldStdout, oldStderr, newReader *os.File

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

// ReadOutput restore os.Stdout
// func ReadOutput() (s string) {
// }

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

// RewriteStderr rewrite os.Stderr
// Usage:
// 	RewriteStderr()
// 	fmt.Fprintln(os.Stderr, "Hello, playground")
// 	msg := RestoreStderr()
func RewriteStderr() {
	oldStderr = os.Stderr
	r, w, _ := os.Pipe()
	newReader = r
	os.Stderr = w
}

// RestoreStderr restore os.Stderr
func RestoreStderr() (s string) {
	if oldStderr == nil {
		return
	}

	// Notice: must close writer before read data
	// close now reader
	_ = os.Stderr.Close()
	// restore
	os.Stderr = oldStderr
	oldStderr = nil

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

// Env mocking

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

// MockOsEnvByText by env text string.
// clear all old ENV data, use given data map, will recover old ENV after fn run.
func MockOsEnvByText(envText string, fn func()) {
	ss := strings.Split(envText, "\n")
	mp := make(map[string]string, len(ss))
	for _, line := range ss {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		nodes := strings.SplitN(line, "=", 2)

		if len(nodes) < 2 {
			mp[nodes[0]] = ""
		} else {
			mp[nodes[0]] = nodes[1]
		}
	}

	MockOsEnv(mp, fn)
}

// MockOsEnv by env map data.
// clear all old ENV data, use given data map, will recover old ENV after fn run.
func MockOsEnv(mp map[string]string, fn func()) {
	envBak := os.Environ()

	os.Clearenv()
	for key, val := range mp {
		_ = os.Setenv(key, val)
	}

	fn()

	os.Clearenv()
	for _, str := range envBak {
		nodes := strings.SplitN(str, "=", 2)
		_ = os.Setenv(nodes[0], nodes[1])
	}
}
