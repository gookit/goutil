package goinfo_test

import (
	"testing"

	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/goinfo"
	"github.com/gookit/goutil/testutil/assert"
)

func TestFuncName(t *testing.T) {
	name := goinfo.FuncName(goinfo.PkgName)
	assert.Eq(t, "github.com/gookit/goutil/goinfo.PkgName", name)

	assert.True(t, goinfo.GoodFuncName("MyFunc"))
	assert.False(t, goinfo.GoodFuncName(""))
	assert.False(t, goinfo.GoodFuncName("+MyFunc"))
	assert.False(t, goinfo.GoodFuncName("My+Func"))
}

func TestPkgName(t *testing.T) {
	name := goinfo.PkgName(goinfo.FuncName(goinfo.GetCallerInfo))
	assert.Eq(t, "github.com/gookit/goutil/goinfo", name)
}

func TestFullFcName_Parse(t *testing.T) {
	fullName := goinfo.FuncName(goinfo.GetCallerInfo)

	ffn := goinfo.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, fullName, ffn.String())
	assert.Eq(t, "goinfo", ffn.PkgName())
	assert.Eq(t, "GetCallerInfo", ffn.FuncName())
	assert.Eq(t, "github.com/gookit/goutil/goinfo", ffn.PkgPath())
	dump.P(ffn)

	st := goinfo.FullFcName{}
	fullName = goinfo.FuncName(st.FuncName)

	ffn = goinfo.FullFcName{FullName: fullName}
	ffn.Parse()
	assert.Eq(t, "(*FullFcName).FuncName-fm", ffn.FuncName())
	dump.P(ffn)
}

func TestCutFuncName(t *testing.T) {
	fullName := goinfo.FuncName(goinfo.GetCallerInfo)

	pkgPath, funcName := goinfo.CutFuncName(fullName)
	assert.Eq(t, "GetCallerInfo", funcName)
	assert.Eq(t, "github.com/gookit/goutil/goinfo", pkgPath)
}
