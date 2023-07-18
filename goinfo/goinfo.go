// Package goinfo provide some standard util functions for go.
package goinfo

import "runtime"

// GoVersion get go runtime version. eg: "1.18.2"
func GoVersion() string {
	return runtime.Version()[2:]
}
