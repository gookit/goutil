// Package assert Provides commonly asserts functions for help write Go testing.
//
// inspired the package: github.com/stretchr/testify/assert
package assert

import (
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/goutil/internal/comfunc"
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
)

// DisableColor render
func DisableColor() {
	EnableColor = false
}

// HideFullPath render
func HideFullPath() {
	ShowFullPath = false
}

// fail reports a failure through
func fail(t TestingT, failMsg string, fmtAndArgs []any) bool {
	t.Helper()

	tName := t.Name()
	if EnableColor {
		tName = color.Red.Sprint(tName)
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
	return false
}
