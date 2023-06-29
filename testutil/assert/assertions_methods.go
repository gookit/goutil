package assert

// Nil asserts that the given is a nil value
func (as *Assertions) Nil(give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Nil(as.t, give, fmtAndArgs...)
	return as
}

// NotNil asserts that the given is a not nil value
func (as *Assertions) NotNil(val any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotNil(as.t, val, fmtAndArgs...)
	return as
}

// True check, please see True()
func (as *Assertions) True(give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = True(as.t, give, fmtAndArgs...)
	return as
}

// False check, please see False()
func (as *Assertions) False(give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = False(as.t, give, fmtAndArgs...)
	return as
}

// Empty check, please see Empty()
func (as *Assertions) Empty(give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Empty(as.t, give, fmtAndArgs...)
	return as
}

// NotEmpty check, please see NotEmpty()
func (as *Assertions) NotEmpty(give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEmpty(as.t, give, fmtAndArgs...)
	return as
}

// Panics check, please see Panics()
func (as *Assertions) Panics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Panics(as.t, fn, fmtAndArgs...)
	return as
}

// NotPanics check, please see NotPanics()
func (as *Assertions) NotPanics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotPanics(as.t, fn, fmtAndArgs...)
	return as
}

// PanicsMsg check, please see PanicsMsg()
func (as *Assertions) PanicsMsg(fn PanicRunFunc, wantVal any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsMsg(as.t, fn, wantVal, fmtAndArgs...)
	return as
}

// PanicsErrMsg check, please see PanicsErrMsg()
func (as *Assertions) PanicsErrMsg(fn PanicRunFunc, errMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsErrMsg(as.t, fn, errMsg, fmtAndArgs...)
	return as
}

// Contains asserts that the given data(string,slice,map) should contain element
func (as *Assertions) Contains(src, elem any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Contains(as.t, src, elem, fmtAndArgs...)
	return as
}

// NotContains asserts that the given data(string,slice,map) should not contain element
func (as *Assertions) NotContains(src, elem any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotContains(as.t, src, elem, fmtAndArgs...)
	return as
}

// ContainsKey asserts that the given map is contains key
func (as *Assertions) ContainsKey(mp, key any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ContainsKey(as.t, mp, key, fmtAndArgs...)
	return as
}

// StrContains asserts that the given strings is contains sub-string
func (as *Assertions) StrContains(s, sub string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = StrContains(as.t, s, sub, fmtAndArgs...)
	return as
}

// NoErr asserts that the given is a nil error
func (as *Assertions) NoErr(err error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NoErr(as.t, err, fmtAndArgs...)
	return as
}

// NoError asserts that the given is a nil error
func (as *Assertions) NoError(err error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NoErr(as.t, err, fmtAndArgs...)
	return as
}

// Err asserts that the given is a not nil error
func (as *Assertions) Err(err error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Err(as.t, err, fmtAndArgs...)
	return as
}

// Error asserts that the given is a not nil error
func (as *Assertions) Error(err error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Err(as.t, err, fmtAndArgs...)
	return as
}

// ErrIs asserts that the given error is equals wantErr
func (as *Assertions) ErrIs(err, wantErr error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrIs(as.t, err, wantErr, fmtAndArgs...)
	return as
}

// ErrMsg asserts that the given is a not nil error and error message equals wantMsg
func (as *Assertions) ErrMsg(err error, errMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrMsg(as.t, err, errMsg, fmtAndArgs...)
	return as
}

// ErrSubMsg asserts that the given is a not nil error and the error message contains subMsg
func (as *Assertions) ErrSubMsg(err error, subMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrSubMsg(as.t, err, subMsg, fmtAndArgs...)
	return as
}

// Len assert given length is equals to wantLn
func (as *Assertions) Len(give any, wantLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Len(as.t, give, wantLn, fmtAndArgs...)
	return as
}

// LenGt assert given length is greater than to minLn
func (as *Assertions) LenGt(give any, minLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = LenGt(as.t, give, minLn, fmtAndArgs...)
	return as
}

// Eq asserts that the want should equal to the given
func (as *Assertions) Eq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Eq(as.t, want, give, fmtAndArgs...)
	return as
}

// Equal asserts that the want should equal to the given
//
// Alias of Eq()
func (as *Assertions) Equal(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Eq(as.t, want, give, fmtAndArgs...)
	return as
}

// Neq asserts that the want should not be equal to the given.
// alias of NotEq()
func (as *Assertions) Neq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Neq(as.t, want, give, fmtAndArgs...)
	return as
}

// NotEq asserts that the want should not be equal to the given
func (as *Assertions) NotEq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEq(as.t, want, give, fmtAndArgs...)
	return as
}

// NotEqual asserts that the want should not be equal to the given
//
// Alias of NotEq()
func (as *Assertions) NotEqual(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEq(as.t, want, give, fmtAndArgs...)
	return as
}

// Lt asserts that the give(intX) should not be less than max
func (as *Assertions) Lt(give, max any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Lt(as.t, give, max, fmtAndArgs...)
	return as
}

// Lte asserts that the give(intX) should not be less than or equal to max
func (as *Assertions) Lte(give, max any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Lte(as.t, give, max, fmtAndArgs...)
	return as
}

// Gt asserts that the give(intX) should not be greater than min
func (as *Assertions) Gt(give, min any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Gt(as.t, give, min, fmtAndArgs...)
	return as
}

// Gte asserts that the give(intX) should not be greater than or equal to min
func (as *Assertions) Gte(give, min any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Gte(as.t, give, min, fmtAndArgs...)
	return as
}

// IsType type equals assert
func (as *Assertions) IsType(wantType, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = IsType(as.t, wantType, give, fmtAndArgs...)
	return as
}

// Fail reports a failure through
func (as *Assertions) Fail(failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Fail(as.t, failMsg, fmtAndArgs...)
	return as
}

// FailNow fails test
func (as *Assertions) FailNow(failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = FailNow(as.t, failMsg, fmtAndArgs...)
	return as
}
