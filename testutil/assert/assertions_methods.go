package assert

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

func (as *Assertions) True(give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = True(as.t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) False(give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = False(as.t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Empty(give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Empty(as.t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) NotEmpty(give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEmpty(as.t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Panics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Panics(as.t, fn, fmtAndArgs...)
	return as
}

func (as *Assertions) NotPanics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotPanics(as.t, fn, fmtAndArgs...)
	return as
}

func (as *Assertions) PanicsMsg(fn PanicRunFunc, wantVal interface{}, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsMsg(as.t, fn, wantVal, fmtAndArgs...)
	return as
}

func (as *Assertions) PanicsErrMsg(fn PanicRunFunc, errMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsErrMsg(as.t, fn, errMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) Contains(src, elem any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Contains(as.t, src, elem, fmtAndArgs...)
	return as
}

func (as *Assertions) NotContains(src, elem any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotContains(as.t, src, elem, fmtAndArgs...)
	return as
}

func (as *Assertions) ContainsKey(mp, key any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ContainsKey(as.t, mp, key, fmtAndArgs...)
	return as
}

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

// Err asserts that the given is a not nil error
func (as *Assertions) Err(err error, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Err(as.t, err, fmtAndArgs...)
	return as
}

func (as *Assertions) ErrMsg(err error, errMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrMsg(as.t, err, errMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) ErrSubMsg(err error, subMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrSubMsg(as.t, err, subMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) Len(give any, wantLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Len(as.t, give, wantLn, fmtAndArgs...)
	return as
}

func (as *Assertions) LenGt(give any, minLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = LenGt(as.t, give, minLn, fmtAndArgs...)
	return as
}

func (as *Assertions) Eq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Eq(as.t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Neq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Neq(as.t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) NotEq(want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEq(as.t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Lt(give, max int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Lt(as.t, give, max, fmtAndArgs...)
	return as
}

func (as *Assertions) Gt(give, min int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Gt(as.t, give, min, fmtAndArgs...)
	return as
}

func (as *Assertions) Fail(failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Fail(as.t, failMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) FailNow(failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = FailNow(as.t, failMsg, fmtAndArgs...)
	return as
}
