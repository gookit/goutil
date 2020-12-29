package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
)

var (
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

	hidden = []string{"netutil", "internal"}
)

var genOpts = struct {
	output string
}{}

func init() {
	flag.StringVar(&genOpts.output, "o", "./metadata.log", "the result output file. if is 'stdout', will direct print it.")
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
	}

	var name string
	var pkgFuncs = make(map[string][]string)

	basePkg := "github.com/gookit/goutil"
	reg := regexp.MustCompile(`func [A-Z]\w+\(.*\).*`)

	fmt.Println("find and collect exported functions...")
	for _, filename := range ms {
		// "jsonutil/jsonutil.go"
		if strings.HasSuffix(filename, "_test.go") {
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
				_, _ = fmt.Fprintln(out, "```")
			}

			name = dir
			if strings.HasSuffix(dir, "util") {
				name = dir[:len(dir)-4]
			}

			if setTitle, ok := nameMap[name]; ok {
				name = setTitle
			}

			_, _ = fmt.Fprintln(out, "\n###", strutil.UpperFirst(name))
			_, _ = fmt.Fprintf(out, "\n> Package `%s`\n\n", pkgPath)
			pkgFuncs[pkgPath] = []string{"xx"}

			_, _ = fmt.Fprintln(out, "```go")
		}

		// read contents
		text := fsutil.MustReadFile(filename)
		lines := reg.FindAllString(string(text), -1)

		_, _ = fmt.Fprintln(out, "// source at", filename)
		for _, line := range lines {
			_, _ = fmt.Fprintln(out, strings.TrimRight(line, "{ "))
		}
	}

	if len(pkgFuncs) > 0 {
		_, _ = fmt.Fprintln(out, "```")
	}

	if toFile {
		fmt.Println("OK. write result to the", genOpts.output)
	}
}
