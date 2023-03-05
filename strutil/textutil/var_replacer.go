package textutil

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
)

const defaultVarFormat = "{{,}}"

// VarReplacer struct
type VarReplacer struct {
	init        bool
	Left, Right string
	lLen, rLen  int
	varReg      *regexp.Regexp
}

// NewVarReplacer instance
func NewVarReplacer(format string) *VarReplacer {
	return (&VarReplacer{}).WithFormat(format)
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
		varMap[path] = strutil.QuietString(val.Interface())
	})

	return r.Init().doReplace(s, varMap)
}

// RenderSimple string-map vars in the text contents
func (r *VarReplacer) RenderSimple(s string, varMap map[string]string) string {
	if len(varMap) == 0 || !strings.Contains(s, r.Left) {
		return s
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
