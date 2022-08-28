package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/testutil/assert"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	// Trim
	tests := map[string]string{
		"abc ":  "",
		" abc":  "",
		" abc ": "",
		"abc,,": ",",
		"abc,.": ",.",
	}
	for sample, cutSet := range tests {
		is.Eq("abc", strutil.Trim(sample, cutSet))
	}

	is.Eq("abc", strutil.Trim("abc,.", ".,"))
	is.Eq("abc", strutil.Trim(", abc ,", ",", " "))

	// TrimLeft
	is.Eq("abc ", strutil.Ltrim(" abc "))
	is.Eq("abc ", strutil.LTrim(" abc "))
	is.Eq("abc ,", strutil.TrimLeft(", abc ,", " ,"))
	is.Eq("abc ,", strutil.TrimLeft(", abc ,", ", "))
	is.Eq("abc ,", strutil.TrimLeft(", abc ,", ",", " "))
	is.Eq(" abc ,", strutil.TrimLeft(", abc ,", ","))

	// TrimRight
	is.Eq(" abc", strutil.Rtrim(" abc "))
	is.Eq(" abc", strutil.RTrim(" abc "))
	is.Eq(", abc", strutil.TrimRight(", abc ,", ", "))
	is.Eq(", abc ", strutil.TrimRight(", abc ,", ","))
	is.Eq(", abc", strutil.TrimRight(", abc ,", ",", " "))
}

func TestFilterEmail(t *testing.T) {
	is := assert.New(t)
	is.Eq("THE@inhere.com", strutil.FilterEmail("   THE@INHere.com  "))
	is.Eq("inhere.xyz", strutil.FilterEmail("   inhere.xyz  "))
}
