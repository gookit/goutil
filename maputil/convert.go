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

// ToStringMap convert map[string]any to map[string]string
func ToStringMap(src map[string]any) map[string]string {
	newMp := make(map[string]string, len(src))
	for k, v := range src {
		newMp[k] = strutil.MustString(v)
	}

	return newMp
}

// HttpQueryString convert map[string]any data to http query string.
func HttpQueryString(data map[string]any) string {
	ss := make([]string, 0, len(data))
	for k, v := range data {
		ss = append(ss, k+"="+strutil.QuietString(v))
	}

	return strings.Join(ss, "&")
}

// ToString simple and quickly convert map[string]any to string.
func ToString(mp map[string]any) string {
	if mp == nil {
		return ""
	}
	if len(mp) == 0 {
		return "{}"
	}

	buf := make([]byte, 0, len(mp)*16)
	buf = append(buf, '{')

	for k, val := range mp {
		buf = append(buf, k...)
		buf = append(buf, ':')

		str := strutil.QuietString(val)
		buf = append(buf, str...)
		buf = append(buf, ',', ' ')
	}

	// remove last ', '
	buf = append(buf[:len(buf)-2], '}')
	return strutil.Byte2str(buf)
}

// ToString2 simple and quickly convert a map to string.
func ToString2(mp any) string {
	return NewFormatter(mp).Format()
}

// FormatIndent format map data to string with newline and indent.
func FormatIndent(mp any, indent string) string {
	return NewFormatter(mp).WithIndent(indent).Format()
}
