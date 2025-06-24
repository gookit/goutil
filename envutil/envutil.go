// Package envutil provide some commonly ENV util functions.
package envutil

import (
	"os"
	"strings"

	"github.com/gookit/goutil/internal/varexpr"
)

// ValueGetter Env value provider func.
//
// TIPS: you can custom provide data.
var ValueGetter = os.Getenv

// VarReplace replaces ${var} or $var in the string according to the values.
//
// is alias of the os.ExpandEnv()
func VarReplace(s string) string { return os.ExpandEnv(s) }

// ParseOrErr parse ENV var value from input string, support default value.
//
// Diff with the ParseValue, this support return error.
//
// With error format: ${VAR_NAME | ?error}
func ParseOrErr(val string) (string, error) {
	return varexpr.Parse(val)
}

// ParseValue parse ENV var value from input string, support default value.
//
// Format:
//
//	${var_name}            Only var name
//	${var_name | default}  With default value
//
// Usage:
//
//	envutil.ParseValue("${ APP_NAME }")
//	envutil.ParseValue("${ APP_ENV | dev }")
func ParseValue(val string) string {
	return varexpr.SafeParse(val)
}

// VarParse alias of the ParseValue
func VarParse(val string) string {
	return varexpr.SafeParse(val)
}

// ParseEnvValue alias of the ParseValue
func ParseEnvValue(val string) string {
	return varexpr.SafeParse(val)
}

// SplitText2map parse ENV text to map. Can use to parse .env file contents.
func SplitText2map(text string) map[string]string {
	lines := strings.Split(text, "\n")
	envMp := make(map[string]string)

	for _, line := range lines {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}

		// skip comments line
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		k, v := splitLineToKv(line)
		if len(k) > 0 {
			envMp[k] = v
		}
	}

	return envMp
}

// SplitLineToKv parse ENV line to k-v. eg: 'DEBUG=true' => ['DEBUG', 'true']
func SplitLineToKv(line string) (string, string) {
	if line = strings.TrimSpace(line); line == "" {
		return "", ""
	}
	return splitLineToKv(line)
}

// splitLineToKv parse ENV line to k-v. eg:
// 	'DEBUG=true' => ['DEBUG', 'true']
//
// NOTE: line must contain '=', allow: 'ENV_KEY='
func splitLineToKv(line string) (string, string) {
	nodes := strings.SplitN(line, "=", 2)
	envKey := strings.TrimSpace(nodes[0])

	// key cannot be empty
	if envKey == "" {
		return "", ""
	}

	if len(nodes) < 2 {
		if strings.ContainsRune(line, '=') {
			return envKey, ""
		}
		return "", ""
	}
	return envKey, strings.TrimSpace(nodes[1])
}