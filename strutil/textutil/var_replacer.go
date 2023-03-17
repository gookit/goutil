package textutil

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
)

const defaultVarFormat = "{{,}}"

// VarReplacer struct
type VarReplacer struct {
	init        bool
	Left, Right string
	lLen, rLen  int

	parseEnv bool
	varReg   *regexp.Regexp
}

// NewVarReplacer instance
func NewVarReplacer(format string) *VarReplacer {
	return (&VarReplacer{}).WithFormat(format)
}

// WithParseEnv custom var template
func (r *VarReplacer) WithParseEnv() *VarReplacer {
	r.parseEnv = true
	return r
}

// WithFormat custom var template
func (r *VarReplacer) WithFormat(format string) *VarReplacer {
	r.Left, r.Right = strutil.QuietCut(strutil.OrElse(format, defaultVarFormat), ",")
	r.Init()
	return r
}

// Init var matcher
func (r *VarReplacer) Init() *VarReplacer {
	if !r.init {
		r.lLen, r.rLen = len(r.Left), len(r.Right)
		if r.Right != "" {
			r.varReg = regexp.MustCompile(regexp.QuoteMeta(r.Left) + `([\w\s.-]+)` + regexp.QuoteMeta(r.Right))
		} else {
			// no right tag. eg: $name, $user.age
			// r.varReg = regexp.MustCompile(regexp.QuoteMeta(r.Left) + `([\w.-]+)`)
			r.varReg = regexp.MustCompile(regexp.QuoteMeta(r.Left) + `(\w[\w-]*(?:\.[\w-]+)*)`)
		}
	}

	return r
}

// Replace any-map vars in the text contents
func (r *VarReplacer) Replace(s string, tplVars map[string]any) string {
	if len(tplVars) == 0 || !strings.Contains(s, r.Left) {
		return s
	}

	varMap := make(map[string]string, len(tplVars)*2)
	maputil.FlatWithFunc(tplVars, func(path string, val reflect.Value) {
		if val.Kind() == reflect.String {
			if r.parseEnv {
				varMap[path] = comfunc.ParseEnvVar(val.String(), nil)
			} else {
				varMap[path] = val.String()
			}
		} else {
			varMap[path] = strutil.QuietString(val.Interface())
		}
	})

	return r.Init().doReplace(s, varMap)
}

// ReplaceSMap string-map vars in the text contents
func (r *VarReplacer) ReplaceSMap(s string, varMap map[string]string) string {
	return r.RenderSimple(s, varMap)
}

// RenderSimple string-map vars in the text contents. alias of ReplaceSMap()
func (r *VarReplacer) RenderSimple(s string, varMap map[string]string) string {
	if len(varMap) == 0 || !strings.Contains(s, r.Left) {
		return s
	}

	if r.parseEnv {
		for name, val := range varMap {
			varMap[name] = comfunc.ParseEnvVar(val, nil)
		}
	}

	return r.Init().doReplace(s, varMap)
}

// Replace string-map vars in the text contents
func (r *VarReplacer) doReplace(s string, varMap map[string]string) string {
	return r.varReg.ReplaceAllStringFunc(s, func(sub string) string {
		varName := strings.TrimSpace(sub[r.lLen : len(sub)-r.rLen])
		if val, ok := varMap[varName]; ok {
			return val
		}
		return sub
	})
}
