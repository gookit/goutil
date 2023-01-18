// Package textutil provide some extra text handle util
package textutil

// ReplaceVars by regex replace given tpl vars.
//
// If format is empty, will use {const defaultVarFormat}
func ReplaceVars(text string, vars map[string]any, format string) string {
	return NewVarReplacer(format).Replace(text, vars)
}
