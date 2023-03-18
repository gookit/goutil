package textutil

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
)

const defaultVarFormat = "{{,}}"

// VarReplacer struct
type VarReplacer struct {
	init bool

	Left, Right string
	lLen, rLen  int

	varReg *regexp.Regexp
	// flatten sub map in vars
	flatSubs bool
	parseEnv bool
	missVars []string
}

// NewVarReplacer instance
func NewVarReplacer(format string) *VarReplacer {
	return (&VarReplacer{flatSubs: true}).WithFormat(format)
}

// DisableFlatten on the input vars map
func (r *VarReplacer) DisableFlatten() *VarReplacer {
	r.flatSubs = false
	return r
}

// WithParseEnv on the input vars value
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

// ParseVars the text contents and collect vars
func (r *VarReplacer) ParseVars(s string) []string {
	ss := arrutil.StringsMap(r.varReg.FindAllString(s, -1), func(val string) string {
		return strings.TrimSpace(val[r.lLen : len(val)-r.rLen])
	})

	return arrutil.Unique(ss)
}

// Render any-map vars in the text contents
func (r *VarReplacer) Render(s string, tplVars map[string]any) string {
	return r.Replace(s, tplVars)
}

// Replace any-map vars in the text contents
func (r *VarReplacer) Replace(s string, tplVars map[string]any) string {
	if len(tplVars) == 0 || !strings.Contains(s, r.Left) {
		return s
	}

	var varMap map[string]string

	if r.flatSubs {
		varMap = make(map[string]string, len(tplVars)*2)
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
	} else {
		varMap = maputil.ToStringMap(tplVars)
	}

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

// MissVars list
func (r *VarReplacer) MissVars() []string {
	return r.missVars
}

// Replace string-map vars in the text contents
func (r *VarReplacer) doReplace(s string, varMap map[string]string) string {
	r.missVars = make([]string, 0) // clear each replace

	return r.varReg.ReplaceAllStringFunc(s, func(sub string) string {
		varName := strings.TrimSpace(sub[r.lLen : len(sub)-r.rLen])
		if val, ok := varMap[varName]; ok {
			return val
		}

		r.missVars = append(r.missVars, varName)
		return sub
	})
}
