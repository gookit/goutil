package sysutil

import (
	"runtime"

	"github.com/gookit/goutil/goinfo"
)

// GoVersion get go runtime version. eg: "1.18.2"
func GoVersion() string {
	return runtime.Version()[2:]
}

// GoInfo define. alias of goinfo.GoInfo
type GoInfo = goinfo.GoInfo

// ParseGoVersion get info by parse `go version` results. alias of goinfo.ParseGoVersion()
//
// Examples:
//
//		line, err := sysutil.ExecLine("go version")
//		if err != nil {
//			return err
//		}
//
//		info, err := sysutil.ParseGoVersion()
//	 	dump.P(info)
func ParseGoVersion(line string) (*GoInfo, error) {
	return goinfo.ParseGoVersion(line)
}

// OsGoInfo fetch and parse. alias of goinfo.OsGoInfo()
func OsGoInfo() (*GoInfo, error) {
	return goinfo.OsGoInfo()
}
