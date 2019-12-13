package strutil

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gookit/goutil/mathutil"
)

var (
	errConvertFail  = errors.New("convert data type is failure")
	errInvalidParam = errors.New("invalid input parameter")

	// some regex for convert string.
	toSnakeReg  = regexp.MustCompile("[A-Z][a-z]")
	toCamelRegs = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
)

/*************************************************************
 * convert value to string
 *************************************************************/

// String convert val to string
func String(val interface{}) (string, error) {
	return ToString(val)
}

// MustString convert value to string
func MustString(in interface{}) string {
	val, _ := ToString(in)
	return val
}

// ToString convert value to string
func ToString(val interface{}) (str string, err error) {
	switch tVal := val.(type) {
	case int:
		str = strconv.Itoa(tVal)
	case int8:
		str = strconv.Itoa(int(tVal))
	case int16:
		str = strconv.Itoa(int(tVal))
	case int32:
		str = strconv.Itoa(int(tVal))
	case int64:
		str = strconv.Itoa(int(tVal))
	case uint:
		str = strconv.Itoa(int(tVal))
	case uint8:
		str = strconv.Itoa(int(tVal))
	case uint16:
		str = strconv.Itoa(int(tVal))
	case uint32:
		str = strconv.Itoa(int(tVal))
	case uint64:
		str = strconv.Itoa(int(tVal))
	case float32:
		str = fmt.Sprint(tVal)
	case float64:
		str = fmt.Sprint(tVal)
	case string:
		str = tVal
	case nil:
		str = ""
	default:
		err = errConvertFail
	}
	return
}

/*************************************************************
 * convert string value to bool
 *************************************************************/

// ToBool convert string to bool
func ToBool(s string) (bool, error) {
	return Bool(s)
}

// MustBool convert.
func MustBool(s string) bool {
	val, _ := Bool(strings.TrimSpace(s))
	return val
}

// Bool parse string to bool
func Bool(s string) (bool, error) {
	// return strconv.ParseBool(Trim(s))
	lower := strings.ToLower(s)
	switch lower {
	case "1", "on", "yes", "true":
		return true, nil
	case "0", "off", "no", "false":
		return false, nil
	}

	return false, fmt.Errorf("'%s' cannot convert to bool", s)
}

/*************************************************************
 * convert string value to int/string slice, time.Time
 *************************************************************/

// ToInts alias of the ToIntSlice()
func ToInts(s string, sep ...string) ([]int, error) {
	return ToIntSlice(s, sep...)
}

// ToIntSlice split string to slice and convert item to int.
func ToIntSlice(s string, sep ...string) (ints []int, err error) {
	ss := ToSlice(s, sep...)
	for _, item := range ss {
		iVal, err := mathutil.ToInt(item)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}
	return
}

// ToArray alias of the ToSlice()
func ToArray(s string, sep ...string) []string {
	return ToSlice(s, sep...)
}

// ToSlice split string to array.
func ToSlice(s string, sep ...string) []string {
	if len(sep) > 0 {
		return Split(s, sep[0])
	}

	return Split(s, ",")
}

// ToTime convert date string to time.Time
func ToTime(s string, layouts ...string) (t time.Time, err error) {
	var layout string
	if len(layouts) > 0 { // custom layout
		layout = layouts[0]
	} else { // auto match layout.
		switch len(s) {
		case 8:
			layout = "20060102"
		case 10:
			layout = "2006-01-02"
		case 13:
			layout = "2006-01-02 15"
		case 16:
			layout = "2006-01-02 15:04"
		case 19:
			layout = "2006-01-02 15:04:05"
		case 20: // time.RFC3339
			layout = "2006-01-02T15:04:05Z07:00"
		}
	}

	if layout == "" {
		err = errInvalidParam
		return
	}

	// has 'T' eg.2006-01-02T15:04:05
	if strings.ContainsRune(s, 'T') {
		layout = strings.Replace(layout, " ", "T", -1)
	}

	// eg: 2006/01/02 15:04:05
	if strings.ContainsRune(s, '/') {
		layout = strings.Replace(layout, "-", "/", -1)
	}

	t, err = time.Parse(layout, s)
	// t, err = time.ParseInLocation(layout, s, time.Local)
	return
}
