// Package fmtutil provide some format util functions.
package fmtutil

import (
	"encoding/json"

	"github.com/gookit/goutil/byteutil"
)

// StringOrJSON to string or encode pretty JSON data to json bytes.
func StringOrJSON(v any) ([]byte, error) {
	return byteutil.ToBytesWithFunc(v, func(v any) ([]byte, error) {
		return json.MarshalIndent(v, "", "    ")
	})
}
