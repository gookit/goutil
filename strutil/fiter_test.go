package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
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
		is.Equal("abc", strutil.Trim(sample, cutSet))
	}

	is.Equal("abc", strutil.Trim("abc,.", ".,"))
	is.Equal("abc", strutil.Trim(", abc ,", ",", " "))

	// TrimLeft
	is.Equal("abc ", strutil.Ltrim(" abc "))
	is.Equal("abc ", strutil.LTrim(" abc "))
	is.Equal("abc ,", strutil.TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", strutil.TrimLeft(", abc ,", ", "))
	is.Equal("abc ,", strutil.TrimLeft(", abc ,", ",", " "))
	is.Equal(" abc ,", strutil.TrimLeft(", abc ,", ","))

	// TrimRight
	is.Equal(" abc", strutil.Rtrim(" abc "))
	is.Equal(" abc", strutil.RTrim(" abc "))
	is.Equal(", abc", strutil.TrimRight(", abc ,", ", "))
	is.Equal(", abc ", strutil.TrimRight(", abc ,", ","))
}

func TestFilterEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", strutil.FilterEmail("   THE@INHere.com  "))
	is.Equal("inhere.xyz", strutil.FilterEmail("   inhere.xyz  "))
}
