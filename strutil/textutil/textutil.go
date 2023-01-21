// Package textutil provide some extra text handle util
package textutil

import "strings"

// ReplaceVars by regex replace given tpl vars.
//
// If format is empty, will use {const defaultVarFormat}
func ReplaceVars(text string, vars map[string]any, format string) string {
	return NewVarReplacer(format).Replace(text, vars)
}

// IsMatchAll keywords in the give text string.
//
// TIP: can use ^ for exclude match.
func IsMatchAll(s string, keywords []string) bool {
	for _, keyword := range keywords {
		if keyword == "" {
			continue
		}

		// exclude
		if keyword[0] == '^' && len(keyword) > 1 {
			if strings.Contains(s, keyword[1:]) {
				return false
			}
			continue
		}

		// include
		if !strings.Contains(s, keyword) {
			return false
		}
	}
	return true
}
