package assert

func (as *Assertions) Nil(t TestingT, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Nil(t, give, fmtAndArgs...)
	return as
}

// NotNil asserts that the given is a not nil value
func (as *Assertions) NotNil(val any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotNil(as.t, val, fmtAndArgs...)
	return as
}

func (as *Assertions) True(t TestingT, give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = True(t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) False(t TestingT, give bool, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = False(t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Empty(t TestingT, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Empty(t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) NotEmpty(t TestingT, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEmpty(t, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Panics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Panics(t, fn, fmtAndArgs...)
	return as
}

func (as *Assertions) NotPanics(t TestingT, fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotPanics(t, fn, fmtAndArgs...)
	return as
}

func (as *Assertions) PanicsMsg(t TestingT, fn PanicRunFunc, wantVal interface{}, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsMsg(t, fn, wantVal, fmtAndArgs...)
	return as
}

func (as *Assertions) PanicsErrMsg(t TestingT, fn PanicRunFunc, errMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = PanicsErrMsg(t, fn, errMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) Contains(t TestingT, src, sub any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Contains(t, src, sub, fmtAndArgs...)
	return as
}

func (as *Assertions) ContainsKey(t TestingT, mp, key any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ContainsKey(t, mp, key, fmtAndArgs...)
	return as
}

func (as *Assertions) StrContains(t TestingT, s, sub string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = StrContains(t, s, sub, fmtAndArgs...)
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

func (as *Assertions) ErrSubMsg(t TestingT, err error, subMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = ErrSubMsg(t, err, subMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) Len(t TestingT, give any, wantLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Len(t, give, wantLn, fmtAndArgs...)
	return as
}

func (as *Assertions) LenGt(t TestingT, give any, minLn int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = LenGt(t, give, minLn, fmtAndArgs...)
	return as
}

func (as *Assertions) Eq(t TestingT, want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Eq(t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Neq(t TestingT, want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Neq(t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) NotEq(t TestingT, want, give any, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = NotEq(t, want, give, fmtAndArgs...)
	return as
}

func (as *Assertions) Lt(t TestingT, give, max int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Lt(t, give, max, fmtAndArgs...)
	return as
}

func (as *Assertions) Gt(t TestingT, give, min int, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Gt(t, give, min, fmtAndArgs...)
	return as
}

func (as *Assertions) Fail(t TestingT, failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = Fail(t, failMsg, fmtAndArgs...)
	return as
}

func (as *Assertions) FailNow(t TestingT, failMsg string, fmtAndArgs ...any) *Assertions {
	as.t.Helper()
	as.ok = FailNow(t, failMsg, fmtAndArgs...)
	return as
}
