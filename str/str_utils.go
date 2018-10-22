package str

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
	"text/template"
)

// FindSimilar
func FindSimilar(input string, samples []string) []string {
	var ss []string
	// ins := strings.Split(input, "")

	// fmt.Print(input, ins)

	for _, str := range samples {
		if strings.Contains(str, input) {
			ss = append(ss, str)
		} else {
			// sns := strings.Split(str, "")
		}

		// max find four items
		if len(ss) == 4 {
			break
		}
	}

	// fmt.Println("found ", ss)

	return ss
}

// Substr
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length

	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// Replaces replace multi strings
// pairs - [old => new]
// can also use:
// strings.NewReplacer("old1", "new1", "old2", "new2").Replace(str)
func Replaces(str string, pairs map[string]string) string {
	for old, newVal := range pairs {
		str = strings.Replace(str, old, newVal, -1)
	}

	return str
}

// LowerFirst
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	f := s[0]
	if f >= 'A' && f <= 'Z' {
		return strings.ToLower(string(f)) + string(s[1:])
	}

	return s
}

// UpperFirst upper first char
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	f := s[0]
	if f >= 'a' && f <= 'z' {
		return strings.ToUpper(string(f)) + string(s[1:])
	}

	return s
}

// UpperWord Change the first character of each word to uppercase
func UpperWord(s string) string {
	if len(s) == 0 {
		return s
	}

	ss := strings.Split(s, " ")
	ns := make([]string, len(ss))
	for i, word := range ss {
		ns[i] = UpperFirst(word)
	}

	return strings.Join(ns, " ")
}

// PrettyJson get pretty Json string
func PrettyJson(v interface{}) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")

	return string(out), err
}

// GenMd5 生成32位md5字串
func GenMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encode
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// RenderTemplate
func RenderTemplate(input string, data interface{}, isFile ...bool) string {
	// use buffer receive rendered content
	var buf bytes.Buffer
	var isFilename bool

	if len(isFile) > 0 {
		isFilename = isFile[0]
	}

	t := template.New("cli")

	// don't escape content
	t.Funcs(template.FuncMap{"raw": func(s string) string {
		return s
	}})

	t.Funcs(template.FuncMap{"trim": func(s string) string {
		return strings.TrimSpace(string(s))
	}})

	// join strings
	t.Funcs(template.FuncMap{"join": func(ss []string, sep string) string {
		return strings.Join(ss, sep)
	}})

	// upper first char
	t.Funcs(template.FuncMap{"upFirst": func(s string) string {
		return UpperFirst(s)
	}})

	if isFilename {
		template.Must(t.ParseFiles(input))
	} else {
		template.Must(t.Parse(input))
	}

	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
