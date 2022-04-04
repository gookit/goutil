// Package strutil provide some string,char,byte util functions
package strutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"
)

// Position for padding string
const (
	PosLeft uint8 = iota
	PosRight
)

/*************************************************************
 * String operation
 *************************************************************/

// SplitValid string to slice. will filter empty string node.
func SplitValid(s, sep string) (ss []string) { return Split(s, sep) }

// Split string to slice. will filter empty string node.
func Split(s, sep string) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.Split(s, sep) {
		if val = strings.TrimSpace(val); val != "" {
			ss = append(ss, val)
		}
	}
	return
}

// SplitNValid string to slice. will filter empty string node.
func SplitNValid(s, sep string, n int) (ss []string) { return SplitN(s, sep, n) }

// SplitN string to slice. will filter empty string node.
func SplitN(s, sep string, n int) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	rawList := strings.Split(s, sep)
	for i, val := range rawList {
		if val = strings.TrimSpace(val); val != "" {
			if len(ss) == n-1 {
				ss = append(ss, strings.TrimSpace(strings.Join(rawList[i:], sep)))
				break
			}

			ss = append(ss, val)
		}
	}
	return
}

// SplitTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitTrimmed(s, sep string) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.Split(s, sep) {
		ss = append(ss, strings.TrimSpace(val))
	}
	return
}

// SplitNTrimmed split string to slice.
// will trim space for each node, but not filter empty
func SplitNTrimmed(s, sep string, n int) (ss []string) {
	if s = strings.TrimSpace(s); s == "" {
		return
	}

	for _, val := range strings.SplitN(s, sep, n) {
		ss = append(ss, strings.TrimSpace(val))
	}
	return
}

// Substr for a string.
// if length <= 0, return pos to end.
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	strLn := len(runes)

	// pos is too large
	if pos >= strLn {
		return ""
	}

	stopIdx := pos + length
	if length == 0 || stopIdx > strLn {
		stopIdx = strLn
	} else if length < 0 {
		stopIdx = strLn + length
	}

	return string(runes[pos:stopIdx])
}

// Padding a string.
func Padding(s, pad string, length int, pos uint8) string {
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

// Repeat a string
func Repeat(s string, times int) string {
	if times < 2 {
		return s
	}

	var ss []string
	for i := 0; i < times; i++ {
		ss = append(ss, s)
	}

	return strings.Join(ss, "")
}

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) (chars []rune) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) (chars []byte) {
	for i := 0; i < times; i++ {
		chars = append(chars, char)
	}
	return
}

// Replaces replace multi strings
//
// 	pairs: {old1: new1, old2: new2, ...}
//
// Can also use:
// 	strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	ss := make([]string, len(pairs)*2)
	for old, newVal := range pairs {
		ss = append(ss, old, newVal)
	}

	return strings.NewReplacer(ss...).Replace(str)
}

// PrettyJSON get pretty Json string
// Deprecated: please use fmtutil.PrettyJSON() or jsonutil.Pretty() instead it
func PrettyJSON(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// RenderTemplate render text template
func RenderTemplate(input string, data interface{}, fns template.FuncMap, isFile ...bool) string {
	return RenderText(input, data, fns, isFile...)
}

// RenderText render text template
func RenderText(input string, data interface{}, fns template.FuncMap, isFile ...bool) string {
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

	// custom add template functions
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
