package fsutil

import (
	"io"
	"mime"
	"net/http"
	"os"
	"strings"
)

// refer https://www.freeformatter.com/mime-types-list.html
var builtinMimeTypes = map[string]string{
	".xml": "application/xml",
}

func init() {
	// register builtin mime types
	for ext, mimeTyp := range builtinMimeTypes {
		_ = mime.AddExtensionType(ext, mimeTyp)
	}
}

// ExtsByMimeType returns the extensions known to be associated with the MIME type typ.
//
// returns like: [".html"] [".jpg", ".jpeg]
func ExtsByMimeType(mimeTyp string) ([]string, error) {
	return mime.ExtensionsByType(mimeTyp)
}

// ExtByMimeType returns an extension known to be associated with the MIME type typ.
//
// allow with a default ext on not found.
func ExtByMimeType(mimeTyp, defExt string) (string, error) {
	ss, err := mime.ExtensionsByType(mimeTyp)
	if err != nil {
		if defExt != "" {
			return defExt, nil
		}
		return "", err
	}

	if len(ss) == 0 {
		return defExt, nil
	}

	// always return the best match
	for _, ext := range ss {
		if strings.Index(mimeTyp, ext[1:]) > 0 {
			return ext, nil
		}
	}
	return ss[0], nil
}

// DetectMime detect file mime type. alias of MimeType()
func DetectMime(path string) string {
	return MimeType(path)
}

// MimeType get File Mime Type name. eg "image/png"
func MimeType(path string) (mime string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	return ReaderMimeType(file)
}

// ReaderMimeType get the io.Reader mimeType
//
// Usage:
//
//	file, err := os.Open(filepath)
//	if err != nil {
//		return
//	}
//	mime := ReaderMimeType(file)
func ReaderMimeType(r io.Reader) (mime string) {
	var buf [MimeSniffLen]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}

	return http.DetectContentType(buf[:n])
}
