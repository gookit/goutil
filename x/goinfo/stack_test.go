package goinfo_test

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
	"github.com/gookit/goutil/x/goinfo"
)

func TestGetCallStacks(t *testing.T) {
	msg := goinfo.GetCallStacks(false)
	fmt.Println(string(msg))

	fmt.Println("-------------full stacks-------------")
	msg = goinfo.GetCallStacks(true)
	fmt.Println(string(msg))
}

func TestGetCallersInfo(t *testing.T) {
	cs := someFunc1()
	assert.Len(t, cs, 1)
	assert.Contains(t, cs[0], "goutil/x/goinfo_test.someFunc1(),stack_test.go")

	cs = someFunc2()
	assert.Len(t, cs, 1)
	assert.Contains(t, cs[0], "goutil/x/goinfo_test.someFunc2(),stack_test.go")

	loc := someFunc3()
	assert.NotEmpty(t, loc)
	assert.Contains(t, loc, "goutil/x/goinfo_test.someFunc3(),stack_test.go")
	// dump.P(cs)
}

func someFunc1() []string {
	return goinfo.GetCallersInfo(1, 2)
}

func someFunc2() []string {
	return goinfo.SimpleCallersInfo(1, 1)
}

func someFunc3() string {
	return goinfo.GetCallerInfo(1)
}
