package strutil

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	errConvertFail  = errors.New("convert data type is failure")

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
