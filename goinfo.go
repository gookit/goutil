package goutil

import "github.com/gookit/goutil/x/goinfo"

// FuncName get func name
func FuncName(f any) string {
	return goinfo.FuncName(f)
}

// PkgName get the current package name. alias of goinfo.PkgName()
//
// Usage:
//
//	funcName := goutil.FuncName(fn)
//	pgkName := goutil.PkgName(funcName)
func PkgName(funcName string) string {
	return goinfo.PkgName(funcName)
}
