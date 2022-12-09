package httpctype

// there are some Content-Type MIME of the most common data formats.
const (
	MIMEHTML  = "text/html"
	MIMEText  = "text/plain" // equals MIMEPlain
	MIMEPlain = "text/plain"
	MIMEJSON  = "application/json"
	MIMEXML   = "application/xml"
	MIMEXML2  = "text/xml"
	MIMEYAML  = "application/x-yaml"

	MIMEPOSTForm      = "application/x-www-form-urlencoded"
	MIMEMultiDataForm = "multipart/form-data"

	MIMEPROTOBUF = "application/x-protobuf"
	MIMEMSGPACK  = "application/x-msgpack"
	MIMEMSGPACK2 = "application/msgpack"
)
