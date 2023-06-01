package assert

// Assertions provides assertion methods around the TestingT interface.
type Assertions struct {
	t  TestingT
	ok bool // last assert result
	// prefix message for each assert TODO
	Msg string
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT) *Assertions {
	return &Assertions{t: t}
}

// WithMsg set with prefix message.
func (as *Assertions) WithMsg(msg string) *Assertions {
	as.Msg = msg
	return as
}

// IsOk for last check
func (as *Assertions) IsOk() bool {
	return as.ok
}

// IsFail for last check
func (as *Assertions) IsFail() bool {
	return !as.ok
}
