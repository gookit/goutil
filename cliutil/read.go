package cliutil

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/term"
)

// the global input output stream
var (
	// Input global input stream
	Input io.Reader = os.Stdin
	// Output global output stream
	Output io.Writer = os.Stdout
)

// ReadInput read user input form Stdin
func ReadInput(question string) (string, error) {
	if len(question) > 0 {
		color.Fprint(Output, question)
	}

	scanner := bufio.NewScanner(Input)
	if !scanner.Scan() { // reading
		return "", scanner.Err()
	}

	return strings.TrimSpace(scanner.Text()), nil
}

// ReadLine read one line from user input.
//
// Usage:
//
//	in := cliutil.ReadLine("")
//	ans, _ := cliutil.ReadLine("your name?")
func ReadLine(question string) (string, error) {
	if len(question) > 0 {
		color.Fprint(Output, question)
	}

	reader := bufio.NewReader(Input)
	answer, _, err := reader.ReadLine()
	return strings.TrimSpace(string(answer)), err
}

// ReadFirst read first char
//
// Usage:
//
//	ans, _ := cliutil.ReadFirst("continue?[y/n] ")
func ReadFirst(question string) (string, error) {
	answer, err := ReadFirstByte(question)
	return string(answer), err
}

// ReadFirstByte read first byte char
//
// Usage:
//
//	ans, _ := cliutil.ReadFirstByte("continue?[y/n] ")
func ReadFirstByte(question string) (byte, error) {
	if len(question) > 0 {
		color.Fprint(Output, question)
	}

	reader := bufio.NewReader(Input)
	return reader.ReadByte()
}

// ReadFirstRune read first rune char
func ReadFirstRune(question string) (rune, error) {
	if len(question) > 0 {
		color.Fprint(Output, question)
	}

	reader := bufio.NewReader(Input)
	answer, _, err := reader.ReadRune()
	return answer, err
}

// ReadAsBool check user inputted answer is right
//
// Usage:
//
//	ok := ReadAsBool("are you OK? [y/N]", false)
func ReadAsBool(tip string, defVal bool) bool {
	fChar, err := ReadFirstByte(tip)
	if err == nil && fChar != 0 {
		return ByteIsYes(fChar)
	}
	return defVal
}

// Confirm with user input
func Confirm(tip string, defVal ...bool) bool {
	var defV bool
	mark := " [y/N]: "

	if len(defVal) > 0 && defVal[0] {
		defV = true
		mark = " [Y/n]: "
	}

	return ReadAsBool(tip+mark, defV)
}

// InputIsYes answer: yes, y, Yes, Y
func InputIsYes(ans string) bool {
	return len(ans) > 0 && (ans[0] == 'y' || ans[0] == 'Y')
}

// ByteIsYes answer: yes, y, Yes, Y
func ByteIsYes(ans byte) bool {
	return ans == 'y' || ans == 'Y'
}

// ReadPassword from console terminal
func ReadPassword(question ...string) string {
	if len(question) > 0 {
		print(question[0])
	} else {
		print("Enter Password: ")
	}

	bs, err := term.ReadPassword(syscallStdinFd())
	if err != nil {
		return ""
	}

	println() // new line
	return string(bs)
}
