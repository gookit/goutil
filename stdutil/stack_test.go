package stdutil_test

import (
	"testing"

	"github.com/gookit/goutil/stdutil"
	"github.com/stretchr/testify/assert"
)

func TestGetCallersInfo(t *testing.T) {
	cs := someFunc1()
	assert.Len(t, cs, 1)
	assert.Contains(t, cs[0], "goutil/stdutil_test.someFunc1(),stack_test.go")

	cs = someFunc2()
	assert.Len(t, cs, 1)
	assert.Contains(t, cs[0], "goutil/stdutil_test.someFunc2(),stack_test.go")

	loc := someFunc3()
	assert.NotEmpty(t, loc)
	assert.Contains(t, loc, "goutil/stdutil_test.someFunc3(),stack_test.go")
	// dump.P(cs)
}

func someFunc1() []string {
	return stdutil.GetCallersInfo(1, 2)
}

func someFunc2() []string {
	return stdutil.SimpleCallersInfo(1, 1)
}

func someFunc3() string {
	return stdutil.GetCallerInfo(1)
}
