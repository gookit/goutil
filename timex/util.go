package timex

import (
	"time"

	"github.com/gookit/goutil/basefn"
)

// NowUnix is short of time.Now().Unix()
func NowUnix() int64 { return time.Now().Unix() }

// NowDate quick get current date string. if template is empty, will use DefaultTemplate.
func NowDate(template ...string) string {
	return FormatByTpl(time.Now(), basefn.FirstOr(template, DefaultTemplate))
}

// Format convert time to string use default layout
func Format(t time.Time) string { return t.Format(DefaultLayout) }

// FormatBy given default layout
func FormatBy(t time.Time, layout string) string { return t.Format(layout) }

// Date format time by given date template. see ToLayout() for template parse.
func Date(t time.Time, template ...string) string { return Datetime(t, template...) }

// Datetime convert time to string use template. see ToLayout() for template parse.
func Datetime(t time.Time, template ...string) string {
	return FormatByTpl(t, basefn.FirstOr(template, DefaultTemplate))
}

// DateFormat format time by given date template. see ToLayout()
func DateFormat(t time.Time, template string) string {
	return FormatByTpl(t, template)
}

// FormatByTpl format time by given date template. see ToLayout()
func FormatByTpl(t time.Time, template string) string {
	return t.Format(ToLayout(template))
}

// FormatUnix time seconds use default layout
func FormatUnix(sec int64, layout ...string) string {
	return time.Unix(sec, 0).Format(basefn.FirstOr(layout, DefaultLayout))
}

// FormatUnixBy format time seconds use given layout
func FormatUnixBy(sec int64, layout string) string {
	return time.Unix(sec, 0).Format(layout)
}

// FormatUnixByTpl format time seconds use given date template.
// see ToLayout()
func FormatUnixByTpl(sec int64, template ...string) string {
	layout := ToLayout(basefn.FirstOr(template, DefaultTemplate))
	return time.Unix(sec, 0).Format(layout)
}
