package assert

// Assertions provides assertion methods around the TestingT interface.
type Assertions struct {
	t  TestingT
	ok bool // last result
}

// New makes a new Assertions object for the specified TestingT.
func New(t TestingT) *Assertions {
	return &Assertions{t: t}
}

// IsOk for last check
func (as Assertions) IsOk() bool {
	return as.ok
}

// IsFail for last check
func (as Assertions) IsFail() bool {
	return !as.ok
}
