package httpctype

// there are some Content-Type MIME of the most common data formats.
const (
	MIMEHTML  = "text/html"
	MIMEHtml  = MIMEHTML
	MIMEText  = "text/plain" // equals MIMEPlain
	MIMEPlain = MIMEText

	MIMEJSON = "application/json"

	MIMEYAML = "application/x-yaml"
	MIMEYaml = MIMEYAML
	MIMEXML  = "application/xml"
	MIMEXML2 = "text/xml"

	MIMEForm          = "application/x-www-form-urlencoded"
	MIMEPOSTForm      = MIMEForm
	MIMEDataForm      = "multipart/form-data"
	MIMEMultiDataForm = MIMEDataForm

	MIMEPROTOBUF = "application/x-protobuf"
	MIMEMSGPACK  = "application/x-msgpack"
	MIMEMSGPACK2 = "application/msgpack"
)
