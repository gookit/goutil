// Package assert Provides commonly asserts functions for help write Go testing.
//
// inspired the package: github.com/stretchr/testify/assert
package assert

import (
	"strings"

	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/x/ccolor"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Helper()
	Name() string
	Error(args ...any)
}

//
// -------------------- render error --------------------
//

var (
	// ShowFullPath on show error trace
	ShowFullPath = true
	// EnableColor on show error trace
	EnableColor = true
	// FailFast fail fast, stop test when first error(will call testing.T.FailNow())
	FailFast = false
)

// DisableColor render
func DisableColor() { EnableColor = false }

// HideFullPath render
func HideFullPath() { ShowFullPath = false }

// SetFailFast set fail fast
func SetFailFast(enable bool) { FailFast = enable }

// fail reports a failure through
func fail(t TestingT, failMsg string, fmtAndArgs []any) bool {
	t.Helper()

	tName := t.Name()
	if EnableColor {
		tName = ccolor.Red.Sprint(tName)
	}

	labeledTexts := []labeledText{
		{"Test Name", tName},
		{"Error Pos", strings.Join(callerInfos(), "\n")},
		{"Error Msg", failMsg},
	}

	// user custom message
	if userMsg := comfunc.FormatWithArgs(fmtAndArgs); len(userMsg) > 0 {
		labeledTexts = append(labeledTexts, labeledText{"User Msg", userMsg})
	}

	t.Error("\n" + formatLabeledTexts(labeledTexts))

	// fail fast handle
	if FailFast {
		if fnr, ok := t.(failNower); ok {
			fnr.FailNow()
		}
	}
	return false
}

//
// -------------------- required --------------------
//

// Must assert that the given condition is true. alias of Require()
//
// If it's not, it calls t.FailNow() to terminate the test.
//
// Usage:
//	assert.Must(t, assert.True(false))
func Must(t TestingT, condition bool, fmtAndArgs ...any) {
	t.Helper()
	if !condition {
		FailNow(t, "Required", fmtAndArgs...)
	}
}

// Require asserts that the given condition is true.
//
// If it's not, it calls t.FailNow() to terminate the test.
//
// Usage:
//	assert.Require(t, assert.True(false))
func Require(t TestingT, condition bool, fmtAndArgs ...any) {
	t.Helper()
	if !condition {
		FailNow(t, "Required", fmtAndArgs...)
	}
}

//
// -------------------- fail --------------------
//

// Fail reports a failure through
func Fail(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	return fail(t, failMsg, fmtAndArgs)
}

type failNower interface {
	FailNow()
}

// FailNow fails test
func FailNow(t TestingT, failMsg string, fmtAndArgs ...any) bool {
	t.Helper()
	fail(t, failMsg, fmtAndArgs)

	if fnr, ok := t.(failNower); ok {
		fnr.FailNow()
	}
	return false
}

