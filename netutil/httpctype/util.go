package httpctype

import "strings"

// ToKind name match
func ToKind(cType, defaultType string) string {
	if len(cType) == 0 {
		return defaultType
	}

	return ToKindWithFunc(cType, func(_ string) string {
		return defaultType
	})
}

// ToKindWithFunc match base kind name by content-type, with a fallback func
func ToKindWithFunc(cType string, fbFunc func(cType string) string) string {
	// JSON body request: "application/json"
	if strings.Contains(cType, "/json") {
		return KindJSON
	}

	// basic POST form data binding. content type: "application/x-www-form-urlencoded"
	if strings.Contains(cType, "/x-www-form-urlencoded") {
		return KindForm
	}

	// contains file uploaded form: "multipart/form-data" "multipart/mixed"
	// strings.HasPrefix(mediaType, "multipart/")
	if strings.Contains(cType, "/form-data") {
		return KindFormData
	}

	// XML body request: "text/xml"
	if strings.Contains(cType, "/xml") {
		return KindXML
	}

	if fbFunc != nil {
		return fbFunc(cType)
	}
	return ""
}
