package httphelper

import (
	"net/http"
	"strings"
)

// HeaderToStringMap convert
func HeaderToStringMap(rh http.Header) map[string]string {
	if len(rh) == 0 {
		return nil
	}

	mp := make(map[string]string, len(rh))
	for name, values := range rh {
		mp[name] = strings.Join(values, "; ")
	}
	return mp
}
