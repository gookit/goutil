package ccolor

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

// ColorsToCode convert colors to code. return like "32;45;3"
func ColorsToCode(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, color := range colors {
		codes = append(codes, color.String())
	}

	return strings.Join(codes, ";")
}

/*************************************************************
 * print methods(will auto parse color tags)
 *************************************************************/

// Print parse color tag and print messages
func Print(v ...any) { Fprint(output, v...) }

// Printf format and print messages
func Printf(format string, v ...any) { Fprintf(output, format, v...) }

// Println messages with new line
func Println(v ...any) { Fprintln(output, v...) }

// Sprint parse color tags, return rendered string
func Sprint(v ...any) string {
	return ReplaceTag(fmt.Sprint(v...))
}

// Sprintf format and return rendered string
func Sprintf(format string, a ...any) string {
	return ReplaceTag(fmt.Sprintf(format, a...))
}

// Fprint auto parse color-tag, print rendered messages to the writer
func Fprint(w io.Writer, v ...any) {
	_, lastErr = fmt.Fprint(w, ReplaceTag(fmt.Sprint(v...)))
}

// Fprintf auto parse color-tag, print rendered messages to the writer.
func Fprintf(w io.Writer, format string, v ...any) {
	_, lastErr = fmt.Fprint(w, ReplaceTag(fmt.Sprintf(format, v...)))
}

// Fprintln auto parse color-tag, print rendered messages to the writer
func Fprintln(w io.Writer, v ...any) {
	_, lastErr = fmt.Fprintln(w, ReplaceTag(formatLikePrintln(v)))
}

// Lprint passes colored messages to a log.Logger for printing.
func Lprint(l *log.Logger, v ...any) {
	l.Print(ReplaceTag(fmt.Sprint(v...)))
}

/*************************************************************
 * render color code
 *************************************************************/

// RenderCode render message by color code.
//
// Usage:
//
//	msg := RenderCode("3;32;45", "some", "message")
func RenderCode(code string, args ...any) string {
	var message string
	if ln := len(args); ln == 0 {
		return ""
	}

	message = fmt.Sprint(args...)
	if len(code) == 0 {
		return message
	}

	// disabled OR not support color
	if !supportColor {
		return ClearCode(message)
	}

	// return fmt.Sprintf(FullColorTpl, code, message)
	return StartSet + code + "m" + message + ResetSet
}

// RenderString render a string with color code.
//
// Usage:
//
//	msg := RenderString("3;32;45", "a message")
func RenderString(code string, str string) string {
	if len(code) == 0 || str == "" {
		return str
	}

	// disabled OR not support color
	if !supportColor {
		return ClearCode(str)
	}

	// return fmt.Sprintf(FullColorTpl, code, str)
	return StartSet + code + "m" + str + ResetSet
}

// RenderWithSpaces Render code with spaces.
// If the number of args is > 1, a space will be added between the args
func RenderWithSpaces(code string, args ...any) string {
	msg := formatLikePrintln(args)
	if len(code) == 0 {
		return msg
	}

	// disabled OR not support color
	if !supportColor {
		return ClearCode(msg)
	}
	return StartSet + code + "m" + msg + ResetSet
}

// ClearCode clear color codes.
//
// eg:
//
//	"\033[36;1mText\x1b[0m" -> "Text"
func ClearCode(str string) string {
	if !strings.Contains(str, CodeSuffix) {
		return str
	}
	return codeRegex.ReplaceAllString(str, "")
}

/*************************************************************
 * helper methods for print
 *************************************************************/

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrint(code, str string) {
	_, lastErr = fmt.Fprint(output, RenderString(code, str))
}

func doPrintTo(w io.Writer, code, str string) {
	_, lastErr = fmt.Fprint(w, RenderString(code, str))
}

// new implementation, support render full color code on pwsh.exe, cmd.exe
func doPrintln(code string, args []any) {
	_, lastErr = fmt.Fprintln(output, RenderString(code, formatLikePrintln(args)))
}

// use Println, will add spaces for each arg
func formatLikePrintln(args []any) (message string) {
	if ln := len(args); ln == 0 {
		message = ""
	} else if ln == 1 {
		message = fmt.Sprint(args[0])
	} else {
		message = fmt.Sprintln(args...)
		// clear last "\n"
		message = message[:len(message)-1]
	}
	return
}

// 检查是否为终端设备中
func isTerminal() bool {
	fd := int(os.Stdout.Fd())
	return term.IsTerminal(fd)
}
