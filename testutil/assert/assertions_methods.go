package assert

func (as *Assertions) NotNil(val any) *Assertions {
	as.t.Helper()
	NotNil(as.t, val)
	return as
}

func (as *Assertions) NoErr(err error) *Assertions {
	as.t.Helper()
	NoErr(as.t, err)
	return as
}

func (as *Assertions) Err(err error) *Assertions {
	as.t.Helper()
	Err(as.t, err)
	return as
}

func (as *Assertions) ErrMsg(err error, errMsg string) *Assertions {
	as.t.Helper()
	ErrMsg(as.t, err, errMsg)
	return as
}
