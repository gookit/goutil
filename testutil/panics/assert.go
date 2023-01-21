package panics

import (
	"github.com/gookit/goutil/internal/comfunc"
	"github.com/gookit/goutil/stdutil"
)

// IsTrue assert result is true, otherwise will panic
func IsTrue(result bool, fmtAndArgs ...any) {
	if !result {
		panicWithMsg("result should be True", fmtAndArgs)
	}
}

// IsFalse assert result is false, otherwise will panic
func IsFalse(result bool, fmtAndArgs ...any) {
	if result {
		panicWithMsg("result should be False", fmtAndArgs)
	}
}

// IsNil assert result is nil, otherwise will panic
func IsNil(result any, fmtAndArgs ...any) {
	if result != nil {
		panicWithMsg("result should be nil", fmtAndArgs)
	}
}

// NotNil assert result is non-nil, otherwise will panic
func NotNil(result any, fmtAndArgs ...any) {
	if result == nil {
		panicWithMsg("result should not be nil", fmtAndArgs)
	}
}

// IsEmpty assert result is empty, otherwise will panic
func IsEmpty(result any, fmtAndArgs ...any) {
	if !stdutil.IsEmpty(result) {
		panicWithMsg("result should be empty", fmtAndArgs)
	}
}

// NotEmpty assert result is empty, otherwise will panic
func NotEmpty(result any, fmtAndArgs ...any) {
	if stdutil.IsEmpty(result) {
		panicWithMsg("result should not be empty", fmtAndArgs)
	}
}

func panicWithMsg(errMsg string, fmtAndArgs []any) {
	if len(fmtAndArgs) > 0 {
		errMsg = comfunc.FormatTplAndArgs(fmtAndArgs)
	}
	panic(errMsg)
}
