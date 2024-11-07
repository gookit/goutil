package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/cflag"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/stdio"
	"github.com/gookit/goutil/strutil"
)

var (
	// hidden pacakges.
	hidden = []string{
		"basefn",
		"netutil",
		"comdef",
		"internal",
	}
	nameMap = map[string]string{
		"arr":     "array and Slice",
		"str":     "strings",
		"byte":    "Bytes Utils",
		"sys":     "system Utils",
		"math":    "math/Number",
		"fs":      "file System",
		"fmt":     "format Utils",
		"test":    "testing Utils",
		"dump":    "dumper",
		"structs": "structs",
		"json":    "JSON Utils",
		"cli":     "CLI/Console",
		"env":     "ENV/Environment",
		"goinfo":  "Go Info",
	}

	// allowLang = map[string]int{
	// 	"en":    1,
	// 	"zh-CN": 1,
	// }
	exFileNames = []string{
		"color_print.go",
	}
	exSuffixes = []string{
		"_test.go",
		"_windows.go",
		"_darwin.go",
		// "_linux",
	}
)

type genOptsSt struct {
	lang     string
	baseDir  string
	output   string
	template string
	tplDir   string
}

//lint:ignore U1000 for test
func (o genOptsSt) filePattern() string {
	baseDir := genOpts.baseDir
	if baseDir == "/" || baseDir == "./" {
		return baseDir + "*/*.go"
	}
	return strings.TrimRight(baseDir, "/") + "/*/*.go"
}

func (o genOptsSt) tplFilename() string {
	if o.lang == "en" {
		return "README.md.tpl"
	}

	return fmt.Sprintf("README.%s.md.tpl", o.lang)
}

func (o genOptsSt) tplFilepath(givePath string) string {
	if givePath != "" {
		return filepath.Join(o.tplDir, givePath)
	}
	return filepath.Join(o.tplDir, o.tplFilename())
}

var (
	genOpts = genOptsSt{}
	// collected sub package names.
	// short name => full name.
	pkgNames = make(map[string]string, 16)

	partDocTplS = "part-%s-s%s.md"
	partDocTplE = "part-%s%s.md"
)

// go run ./internal/gendoc -h
// go run ./internal/gendoc
func main() {
	cmd := cflag.New(func(c *cflag.CFlags) {
		c.Version = "0.1.2"
		c.Desc = "Collect and dump all exported functions for goutil"
	})

	cmd.StringVar(&genOpts.lang, "lang", "en", "package desc message language. allow: en, zh-CN;;l")
	cmd.StringVar(&genOpts.baseDir, "dir", "./", "the base dir path for collect;;d")
	cmd.StringVar(&genOpts.output,
		"output",
		"./metadata.log",
		"the result output file. if is 'stdout', will direct print it;;o",
	)
	cmd.StringVar(&genOpts.tplDir,
		"tpl",
		"./internal/template",
		"template file dir, use for generate, will inject metadata to the template.\nsee ./internal/template/*.tpl;;t",
	)
	cmd.StringVar(&genOpts.template, "template", "", "the template file")

	cmd.Func = handle
	cmd.Example = `
  go run ./internal/gendoc -o stdout
  go run ./internal/gendoc -o stdout -l zh-CN
  go run ./internal/gendoc -o README.md
  go run ./internal/gendoc -o README.zh-CN.md -l zh-CN
`
	cmd.MustParse(nil)
}

func handle(c *cflag.CFlags) error {
	ms, err := filepath.Glob(genOpts.baseDir + "*/*.go")
	goutil.PanicIfErr(err)

	var out io.Writer
	var toFile bool

	if genOpts.output == "stdout" {
		out = os.Stdout
	} else {
		toFile = true
		out, err = os.OpenFile(genOpts.output, fsutil.FsCWTFlags, fsutil.DefaultFilePerm)
		goutil.PanicIfErr(err)

		// close after handle
		defer out.(*os.File).Close()
	}

	// want output by template file
	// var tplFile *os.File
	var tplBody []byte
	if genOpts.tplDir != "" {
		tplFile := genOpts.tplFilepath("")
		color.Info.Println("- read template file contents from", tplFile)
		tplBody = fsutil.MustReadFile(tplFile)
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

	color.Cyanln("Collected packages:")
	dump.Clear(pkgNames)

	if toFile {
		color.Info.Println("OK. write result to the", genOpts.output)
	}
	return nil
}

func collectPgkFunc(ms []string, basePkg string) *bytes.Buffer {
	var name, dirname string
	var pkgFuncs = make(map[string][]string)

	// match func
	reg := regexp.MustCompile(`func [A-Z]\w+.*`)
	buf := new(bytes.Buffer)

	color.Info.Println("- find and collect exported functions...")
	for _, filename := range ms { // for each go file
		// "jsonutil/jsonutil_test.go"
		// "sysutil/sysutil_windows.go"
		if strutil.HasOneSuffix(filename, exSuffixes) {
			continue
		}
		if strutil.ContainsOne(filename, exFileNames) {
			continue
		}

		idx := strings.IndexRune(filename, '/')
		dir := filename[:idx] // sub pkg name.

		if arrutil.StringsHas(hidden, dir) {
			continue
		}

		pkgPath := basePkg + "/" + dir
		pkgNames[dir] = pkgPath

		if ss, ok := pkgFuncs[pkgPath]; ok {
			pkgFuncs[pkgPath] = append(ss, "added")
		} else {
			if len(pkgFuncs) > 0 { // end of prev package.
				bufWriteln(buf, "```")

				// load prev sub-pkg doc file.
				bufWriteDoc(buf, partDocTplE, dirname)
			}

			dirname = dir
			name = dir
			if strings.HasSuffix(dir, "util") {
				name = dir[:len(dir)-4]
			}

			if setTitle, ok := nameMap[name]; ok {
				name = setTitle
			}

			// now: name is package name.
			bufWriteln(buf, "\n###", strutil.UpperFirst(name))
			bufWritef(buf, "\n> Package `%s`\n\n", pkgPath)
			pkgFuncs[pkgPath] = []string{"xx"}

			// load sub-pkg start doc file.
			bufWriteDoc(buf, partDocTplS, name)

			bufWriteln(buf, "```go")
		}

		// read contents
		text := fsutil.MustReadFile(filename)
		lines := reg.FindAllString(string(text), -1)

		if len(lines) > 0 {
			bufWriteln(buf, "// source at", filename)
			for _, line := range lines {
				idx := strings.IndexByte(line, '{')
				if idx > 0 {
					bufWriteln(buf, strings.TrimSpace(line[:idx]))
				} else {
					bufWriteln(buf, strings.TrimSpace(line))
				}
			}
		}
	}

	if len(pkgFuncs) > 0 {
		bufWriteln(buf, "```")
		// load last sub-pkg doc file.
		bufWriteDoc(buf, partDocTplE, dirname)
	}

	return buf
}

func bufWritef(buf *bytes.Buffer, f string, a ...any) {
	_, _ = fmt.Fprintf(buf, f, a...)
}

func bufWriteln(buf *bytes.Buffer, a ...any) {
	_, _ = fmt.Fprintln(buf, a...)
}

func bufWriteDoc(buf *bytes.Buffer, partNameTpl, pkgName string) {
	var lang string
	if genOpts.lang != "en" {
		lang = "." + genOpts.lang
	}

	filename := fmt.Sprintf(partNameTpl, pkgName, lang)

	if !doWriteDoc2buf(buf, filename) {
		// fallback use en docs
		filename = fmt.Sprintf(partNameTpl, pkgName, "")
		doWriteDoc2buf(buf, filename)
	}
}

func doWriteDoc2buf(buf *bytes.Buffer, filename string) bool {
	partFile := genOpts.tplDir + "/" + filename
	// color.Infoln("- try read part readme from", partFile)
	partBody := fsutil.ReadExistFile(partFile)

	if len(partBody) > 0 {
		color.Infoln("- find and inject sub-package doc:", filename)
		stdio.QuietFprintln(buf, string(partBody))
		return true
	}

	return false
}
