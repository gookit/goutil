package cliutil

import (
	"bufio"
	"os"
	"strings"

	"github.com/gookit/color"
)

// ReadInput read user input form Stdin
func ReadInput(question string) (string, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() { // reading
		return "", scanner.Err()
	}

	answer := scanner.Text()
	return strings.TrimSpace(answer), nil
}

// ReadLine read one line from user input.
//
// Usage:
// 	in := cliutil.ReadLine("")
// 	ans, _ := cliutil.ReadLine("your name?")
func ReadLine(question string) (string, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	reader := bufio.NewReader(os.Stdin)
	answer, _, err := reader.ReadLine()
	return strings.TrimSpace(string(answer)), err
}

// ReadFirst read first char
func ReadFirst(question string) (string, error) {
	answer, err := ReadFirstByte(question)

	return string(answer), err
}

// ReadFirstByte read first byte char
//
// Usage:
// 	ans, _ := cliutil.ReadFirstByte("continue?[y/n] ")
func ReadFirstByte(question string) (byte, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	reader := bufio.NewReader(os.Stdin)
	return reader.ReadByte()
}

// ReadFirstRune read first rune char
func ReadFirstRune(question string) (rune, error) {
	if len(question) > 0 {
		color.Print(question)
	}

	reader := bufio.NewReader(os.Stdin)
	answer, _, err := reader.ReadRune()
	return answer, err
}
