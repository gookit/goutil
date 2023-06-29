// Package httpctype list some common http content-type
package httpctype

// Key is the header key of Content-Type
const Key = "Content-Type"

// there are some HTTP Content-Type with charset of the most common data formats.
const (
	CSS  = "text/css; charset=utf-8"
	HTML = "text/html; charset=utf-8"

	Text  = "text/plain; charset=utf-8" // equals Plain
	Plain = Text

	XML2 = "text/xml; charset=utf-8"
	XML  = "application/xml; charset=utf-8"

	YAML = "application/x-yaml; charset=utf-8"
	YML  = YAML

	JSON  = "application/json; charset=utf-8"
	JSONP = "application/javascript; charset=utf-8" // equals to JS

	JS  = "application/javascript; charset=utf-8"
	JS2 = "text/javascript; charset=utf-8"

	MSGPACK  = "application/x-msgpack; charset=utf-8"
	MSGPACK2 = "application/msgpack; charset=utf-8"

	PROTOBUF = "application/x-protobuf"

	Form = "application/x-www-form-urlencoded"
	// FormData for upload file
	FormData = "multipart/form-data"

	// Binary represents content type application/octet-stream
	Binary = "application/octet-stream"
)
