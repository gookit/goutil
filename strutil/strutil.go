// Package strutil provide some string,char,byte util functions
package strutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

// PosFlag type
type PosFlag uint8

// Position for padding/resize string
const (
	PosLeft PosFlag = iota
	PosRight
	PosMiddle
)

/*************************************************************
 * String padding operation
 *************************************************************/

// Padding a string.
func Padding(s, pad string, length int, pos PosFlag) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if pad == "" || pad == " " {
		mark := ""
		if pos == PosRight { // to right
			mark = "-"
		}

		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%s%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`s`, s)
	}

	if pos == PosRight { // to right
		return s + Repeat(pad, -diff)
	}
	return Repeat(pad, -diff) + s
}

// PadLeft a string.
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, PosLeft)
}

// PadRight a string.
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, PosRight)
}

// Resize a string by given length and align settings. padding space.
func Resize(s string, length int, align PosFlag) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if align == PosMiddle {
		strLn := len(s)
		padLn := (length - strLn) / 2
		padStr := string(make([]byte, padLn, padLn))

		if diff := length - padLn*2; diff > 0 {
			s += " "
		}
		return padStr + s + padStr
	}

	return Padding(s, " ", length, align)
}

/*************************************************************
 * String repeat operation
 *************************************************************/

// Repeat a string
func Repeat(s string, times int) string {
	if times <= 0 {
		return ""
	}
	if times == 1 {
		return s
	}

	ss := make([]string, 0, times)
	for i := 0; i < times; i++ {
		ss = append(ss, s)
	}

	return strings.Join(ss, "")
}

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) []rune {
	chars := make([]rune, 0, times)
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return chars
}

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) []byte {
	chars := make([]byte, 0, times)
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return chars
}

// Replaces replace multi strings
//
//	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
//
//	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	ss := make([]string, len(pairs)*2)
	for old, newVal := range pairs {
		ss = append(ss, old, newVal)
	}

	return strings.NewReplacer(ss...).Replace(str)
}

// PrettyJSON get pretty Json string
// Deprecated: please use fmtutil.PrettyJSON() or jsonutil.Pretty() instead it
func PrettyJSON(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// RenderTemplate render text template
func RenderTemplate(input string, data any, fns template.FuncMap, isFile ...bool) string {
	return RenderText(input, data, fns, isFile...)
}

// RenderText render text template
func RenderText(input string, data any, fns template.FuncMap, isFile ...bool) string {
	t := template.New("simple-text")
	t.Funcs(template.FuncMap{
		// don't escape content
		"raw": func(s string) string {
			return s
		},
		"trim": func(s string) string {
			return strings.TrimSpace(s)
		},
		// join strings
		"join": func(ss []string, sep string) string {
			return strings.Join(ss, sep)
		},
		// lower first char
		"lcFirst": func(s string) string {
			return LowerFirst(s)
		},
		// upper first char
		"upFirst": func(s string) string {
			return UpperFirst(s)
		},
	})

	// add custom template functions
	if len(fns) > 0 {
		t.Funcs(fns)
	}

	if len(isFile) > 0 && isFile[0] {
		template.Must(t.ParseFiles(input))
	} else {
		template.Must(t.Parse(input))
	}

	// use buffer receive rendered content
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.String()
}

// WrapTag for given string.
func WrapTag(s, tag string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("<%s>%s</%s>", tag, s, tag)
}
