package maputil

import (
	"strings"

	"github.com/gookit/goutil/strutil"
)

// KeyToLower convert keys to lower case.
func KeyToLower(src map[string]string) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		k = strings.ToLower(k)
		newMp[k] = v
	}

	return newMp
}

// ToStringMap convert map[string]interface{} to map[string]string
func ToStringMap(src map[string]interface{}) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		newMp[k] = strutil.MustString(v)
	}

	return newMp
}

// HttpQueryString convert map[string]interface{} data to http query string.
func HttpQueryString(data map[string]interface{}) string {
	ss := make([]string, len(data))
	for k, v := range data {
		ss = append(ss, k+"="+strutil.MustString(v))
	}

	return strings.Join(ss, "&")
}
