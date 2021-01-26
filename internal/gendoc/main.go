package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
)

var (
	hidden  = []string{"netutil", "internal"}
	nameMap = map[string]string{
		"arr":  "array/Slice",
		"str":  "string",
		"sys":  "system",
		"math": "math/Number",
		"fs":   "fileSystem",
		"fmt":  "formatting",
		"test": "testing",
		"dump": "dump",
		"json": "JSON",
		"cli":  "CLI",
		"env":  "ENV",
	}

	pkgDesc = map[string]map[string]string{
		"en": {
			"arr": "Package arrutil provides some util functions for array, slice",
		},
		"zh-CN": {
			"arr": "`arrutil` 包提供一些辅助函数，用于数组、切片处理",
		},
	}
)

var genOpts = struct {
	lang     string
	output   string
	template string
}{}

func init() {
	flag.StringVar(&genOpts.lang, "l", "en", "package desc message language. allow: en, zh-CN")
	flag.StringVar(&genOpts.output,
		"o",
		"./metadata.log",
		"the result output file. if is 'stdout', will direct print it.",
	)
	flag.StringVar(&genOpts.template,
		"t",
		"",
		"use template file for generate, will inject metadata to the template. see ./internal/template/*.tpl",
	)

	flag.Usage = func() {
		color.Info.Println("Collect and dump all exported functions for goutil package")
		fmt.Println()

		color.Comment.Println("Options:")
		flag.PrintDefaults()

		color.Comment.Println("Example:")
		fmt.Println(`
  go run ./internal/gendoc -o stdout
  go run ./internal/gendoc -o stdout -t ./internal/template/README.md.tpl
  go run ./internal/gendoc -o README.md -t ./internal/template/README.md.tpl
  go run ./internal/gendoc -o README.zh-CN.md -t ./internal/template/README.zh-CN.md.tpl -l zh-CN
`)
	}
}

// go run ./internal/gendoc
func main() {
	flag.Parse()

	ms, err := filepath.Glob("./*/*.go")
	goutil.PanicIfErr(err)

	var out io.Writer
	var toFile bool

	if genOpts.output == "stdout" {
		out = os.Stdout
	} else {
		toFile = true
		out, err = os.OpenFile(genOpts.output, os.O_CREATE|os.O_WRONLY, fsutil.DefaultFilePerm)
		goutil.PanicIfErr(err)

		// close after handle
		defer out.(*os.File).Close()
	}

	// want output by template file
	// var tplFile *os.File
	var tplBody []byte
	if genOpts.template != "" {
		color.Info.Println("- read template file contents from", genOpts.template)
		tplBody = fsutil.MustReadFile(genOpts.template)
		// tplFile, err = os.OpenFile(genOpts.template, os.O_CREATE|os.O_RDONLY, fsutil.DefaultFilePerm)
		// goutil.PanicIfErr(err)
	}

	basePkg := "github.com/gookit/goutil"

	// collect functions
	buf := collectPgkFunc(ms, basePkg)

	// write to output
	if len(tplBody) > 0 {
		_, err = fmt.Fprint(out, strings.Replace(string(tplBody), "{{pgkFuncs}}", buf.String(), 1))
	} else {
		_, err = buf.WriteTo(out)
	}

	goutil.PanicIfErr(err)

	if toFile {
		color.Info.Println("OK. write result to the", genOpts.output)
	}
}

func collectPgkFunc(ms []string, basePkg string) *bytes.Buffer {
	var name string
	var pkgFuncs = make(map[string][]string)

	// match func
	reg := regexp.MustCompile(`func [A-Z]\w+\(.*\).*`)
	buf := new(bytes.Buffer)

	color.Info.Println("- find and collect exported functions...")
	for _, filename := range ms {
		// "jsonutil/jsonutil_test.go"
		if strings.HasSuffix(filename, "_test.go") {
			continue
		}

		// "sysutil/sysutil_windows.go"
		if strings.HasSuffix(filename, "_windows.go") {
			continue
		}

		idx := strings.IndexRune(filename, '/')
		dir := filename[:idx]

		if arrutil.StringsHas(hidden, dir) {
			continue
		}

		pkgPath := basePkg + "/" + dir
		if ss, ok := pkgFuncs[pkgPath]; ok {
			pkgFuncs[pkgPath] = append(ss, "xxx")
		} else {
			if len(pkgFuncs) > 0 {
				_, _ = fmt.Fprintln(buf, "```")
			}

			name = dir
			if strings.HasSuffix(dir, "util") {
				name = dir[:len(dir)-4]
			}

			if setTitle, ok := nameMap[name]; ok {
				name = setTitle
			}

			_, _ = fmt.Fprintln(buf, "\n###", strutil.UpperFirst(name))
			_, _ = fmt.Fprintf(buf, "\n> Package `%s`\n\n", pkgPath)
			pkgFuncs[pkgPath] = []string{"xx"}

			_, _ = fmt.Fprintln(buf, "```go")
		}

		// read contents
		text := fsutil.MustReadFile(filename)
		lines := reg.FindAllString(string(text), -1)

		_, _ = fmt.Fprintln(buf, "// source at", filename)
		for _, line := range lines {
			_, _ = fmt.Fprintln(buf, strings.TrimRight(line, "{ "))
		}
	}

	if len(pkgFuncs) > 0 {
		_, _ = fmt.Fprintln(buf, "```")
	}

	return buf
}
