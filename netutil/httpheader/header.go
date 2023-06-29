// Package httpheader provides some common http header names.
package httpheader

// some common http header names
const (
	UserAgent = "User-Agent"
	UserAuth  = "Authorization"

	Accept = "Accept"
	Cookie = "Cookie"
	// Upgrade header. check websocket:
	// 	header["Connection"] == "Upgrade" and header["Upgrade"] == "websocket"
	Upgrade = "Upgrade"

	AcceptEnc   = "Accept-Encoding"
	ContentType = "Content-Type"
	Connection  = "Connection"

	XRealIP = "X-Real-IP"

	XForwardedFor   = "X-Forwarded-For"
	XForwardedHost  = "X-Forwarded-Host"
	XForwardedProto = "X-Forwarded-Proto"

	// XRequestedWith header. check ajax: header["X-Requested-With"] == XMLHttpRequest
	XRequestedWith = "X-Requested-With"
)
