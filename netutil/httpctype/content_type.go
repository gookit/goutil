// Package httpctype list some common http content-type
package httpctype

const Key = "Content-Type"

// HTTP Content-Type with charset of the most common data formats.
const (
	CSS  = "text/css; charset=utf-8"
	HTML = "text/html; charset=utf-8"

	Text  = "text/plain; charset=utf-8" // equals Plain
	Plain = "text/plain; charset=utf-8"

	XML2 = "text/xml; charset=utf-8"
	XML  = "application/xml; charset=utf-8"

	YML  = "application/x-yaml; charset=utf-8"
	YAML = "application/x-yaml; charset=utf-8"

	JSON  = "application/json; charset=utf-8"
	JSONP = "application/javascript; charset=utf-8" // equals to JS

	JS  = "application/javascript; charset=utf-8"
	JS2 = "text/javascript; charset=utf-8"

	MSGPACK  = "application/x-msgpack; charset=utf-8"
	MSGPACK2 = "application/msgpack; charset=utf-8"
	PROTOBUF = "application/x-protobuf"

	// Binary represents content type application/octet-stream
	Binary = "application/octet-stream"
)
