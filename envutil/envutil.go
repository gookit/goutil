package envutil

import (
	"os"
	"regexp"
	"strings"
)

// parse env value, allow:
// 	only key 	 - "${SHELL}"
// 	with default - "${NotExist|defValue}"
//	multi key 	 - "${GOPATH}/${APP_ENV | prod}/dir"
// Notice:
//  must add "?" - To ensure that there is no greedy match
//  var envRegex = regexp.MustCompile(`\${[\w-| ]+}`)
var envRegex = regexp.MustCompile(`\${.+?}`)

// EnvValueGetter Env value provider.
// TIPS: you can custom provide data.
var EnvValueGetter = func(name string) string {
	return os.Getenv(name)
}

// ParseEnvValue parse ENV var value from input string
func ParseEnvValue(val string) (newVal string) {
	if strings.Index(val, "${") == -1 {
		return val
	}

	var name, def string

	return envRegex.ReplaceAllStringFunc(val, func(eVar string) string {
		// eVar like "${NotExist|defValue}", first remove "${" and "}", then split it
		ss := strings.SplitN(eVar[2:len(eVar)-1], "|", 2)

		// with default value. ${NotExist|defValue}
		if len(ss) == 2 {
			name, def = strings.TrimSpace(ss[0]), strings.TrimSpace(ss[1])
		} else {
			def = eVar // use raw value
			name = strings.TrimSpace(ss[0])
		}

		// get value from ENV
		// eVal := os.Getenv(name)
		eVal := EnvValueGetter(name)
		if eVal == "" {
			eVal = def
		}
		return eVal
	})
}
