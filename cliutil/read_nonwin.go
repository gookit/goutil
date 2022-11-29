//go:build !windows

package cliutil

import (
	"syscall"

	"golang.org/x/term"
)

// ReadPassword from console terminal
func ReadPassword(question ...string) string {
	if len(question) > 0 {
		print(question[0])
	} else {
		print("Enter Password: ")
	}

	bs, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return ""
	}

	println() // new line
	return string(bs)
}
