package textutil

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/basefn"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/reflects"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
)

// STemplateOptFn template option func
type STemplateOptFn func(opt *StrTemplateOpt)

// StrTemplateOpt template options for StrTemplate
type StrTemplateOpt struct {
	// func name alias map. eg: {"up_first": "upFirst"}
	nameMp structs.Aliases
	Funcs  template.FuncMap

	Left, Right string

	ParseDef bool
	ParseEnv bool
}

// StrTemplate implement a simple string template
//
//   - support parse template vars
//   - support access multi-level map field. eg: {{ user.name }}
//   - support parse default value
//   - support parse env vars
//   - support custom pipeline func handle. eg: {{ name | upper }} {{ name | def:guest }}
//
// NOTE: not support control flow, eg: if/else/for/with
type StrTemplate struct {
	StrTemplateOpt
	vr VarReplacer
	// template func map. refer the text/template
	//
	// Func allow return 1 or 2 values, if return 2 values, the second value is error.
	fxs map[string]*reflects.FuncX
}

// NewStrTemplate instance
func NewStrTemplate(opFns ...STemplateOptFn) *StrTemplate {
	st := &StrTemplate{
		fxs: make(map[string]*reflects.FuncX),
		vr: VarReplacer{
			Left:  "{{",
			Right: "}}",
		},
	}

	st.ParseDef = true
	st.ParseEnv = true
	st.vr.RenderFn = st.renderVars
	for _, fn := range opFns {
		fn(&st.StrTemplateOpt)
	}

	st.Init()
	return st
}

// Init StrTemplate
func (t *StrTemplate) Init() {
	if t.vr.init {
		return
	}

	basefn.PanicIf(t.vr.Right == "", "var format Right chars is required")

	t.vr.init = true
	t.vr.parseDef = t.ParseDef
	t.vr.parseEnv = t.ParseEnv

	t.vr.lLen, t.vr.rLen = len(t.vr.Left), len(t.vr.Right)
	// (?s:...) - 让 "." 匹配换行
	// (?s:(.+?)) - 第二个 "?" 非贪婪匹配
	t.vr.varReg = regexp.MustCompile(regexp.QuoteMeta(t.vr.Left) + `(?s:(.+?))` + regexp.QuoteMeta(t.vr.Right))

	// add built-in funcs
	t.AddFuncs(builtInFuncs)
	t.nameMp.AddAliasMap(map[string]string{
		"up_first": "upFirst",
		"lc_first": "lcFirst",
		"def":      "default",
	})
}

// AddFuncs add custom template functions
func (t *StrTemplate) AddFuncs(fns map[string]any) {
	for name, fn := range fns {
		t.fxs[name] = reflects.NewFunc(fn)
	}
}

// RenderString render template string with vars
func (t *StrTemplate) RenderString(s string, vars map[string]any) string {
	return t.vr.Replace(s, vars)
}

// RenderFile render template file with vars
func (t *StrTemplate) RenderFile(filePath string, vars map[string]any) (string, error) {
	if !fsutil.FileExists(filePath) {
		return "", fmt.Errorf("template file not exists: %s", filePath)
	}

	// read file contents
	s, err := fsutil.ReadStringOrErr(filePath)
	if err != nil {
		return "", err
	}

	return t.vr.Replace(s, vars), nil
}

// RenderWrite render template string with vars, and write to writer
func (t *StrTemplate) RenderWrite(wr io.Writer, s string, vars map[string]any) error {
	s = t.vr.Replace(s, vars)
	_, err := io.WriteString(wr, s)
	return err
}

func (t *StrTemplate) renderVars(s string, varMap map[string]string) string {
	return t.vr.varReg.ReplaceAllStringFunc(s, func(sub string) string {
		// var name or pipe expression.
		name := strings.TrimSpace(sub[t.vr.lLen : len(sub)-t.vr.rLen])
		name = strings.TrimLeft(name, ".")

		var defVal string
		var pipes []string
		if strings.ContainsRune(name, '|') {
			pipes = strutil.Split(name, "|")
			// compatible default value. eg: {{ name | inhere }}
			if len(pipes) == 2 && !strings.ContainsRune(pipes[1], ':') && !t.isFunc(pipes[1]) {
				name, defVal = pipes[0], pipes[1]
				pipes = nil // clear pipes
			} else { // collect pipe functions
				name, pipes = pipes[0], pipes[1:]
			}
		}

		if val, ok := varMap[name]; ok {
			if len(pipes) > 0 {
				var err error
				val, err = t.applyPipes(val, pipes)
				if err != nil {
					return fmt.Sprintf("Render var %q error: %v", name, err)
				}
			}
			return val
		}

		// var not found
		if len(defVal) > 0 {
			return defVal
		}

		if t.vr.NotFound != nil {
			if val, ok := t.vr.NotFound(name); ok {
				return val
			}
		}

		// check is default func. eg: {{ name | def:guest }}
		if len(pipes) == 1 && strings.ContainsRune(pipes[0], ':') {
			fName, argVal := strutil.TrimCut(pipes[0], ":")
			if t.isDefaultFunc(fName) {
				return argVal
			}
		}

		t.vr.missVars = append(t.vr.missVars, name)
		return sub
	})
}

func (t *StrTemplate) applyPipes(val any, pipes []string) (string, error) {
	var err error

	// pipe expr: "trim|upper|substr:1,2"
	// =>
	// pipes: ["trim", "upper", "substr:1,2"]
	for _, name := range pipes {
		args := []any{val}

		// has custom args. eg: "substr:1,2"
		if strings.ContainsRune(name, ':') {
			var argStr string
			name, argStr = strutil.TrimCut(name, ":")

			if otherArgs := parseArgStr(argStr); len(otherArgs) > 0 {
				args = append(args, otherArgs...)
			}
		}

		name = t.nameMp.ResolveAlias(name)

		// call pipe func
		if fx, ok := t.fxs[name]; ok {
			val, err = fx.Call2(args...)
			if err != nil {
				return "", err
			}
		} else {
			return "", fmt.Errorf("template func %q not found", name)
		}
	}

	return strutil.ToString(val)
}

func (t *StrTemplate) isFunc(name string) bool {
	_, ok := t.fxs[name]
	if !ok {
		// check name alias
		return t.nameMp.HasAlias(name)
	}
	return ok
}

func (t *StrTemplate) isDefaultFunc(name string) bool {
	return name == "default" || name == "def"
}

var stdTpl = NewStrTemplate()

// RenderString render str template string or file.
func RenderString(input string, data map[string]any, optFns ...RenderOptFn) string {
	return stdTpl.RenderString(input, data)
}

func parseArgStr(argStr string) (ss []any) {
	if argStr == "" { // no arg
		return
	}

	if len(argStr) == 1 { // one char
		return []any{argStr}
	}
	return arrutil.StringsToAnys(strutil.Split(argStr, ","))
}
