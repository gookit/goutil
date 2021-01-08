package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestUpperOrLowerCase(t *testing.T) {
	// Uppercase, Lowercase
	assert.Equal(t, "ABC", strutil.Upper("abc"))
	assert.Equal(t, "ABC", strutil.Uppercase("abc"))
	assert.Equal(t, "abc", strutil.Lower("ABC"))
	assert.Equal(t, "abc", strutil.Lowercase("ABC"))
}

func TestUpperFirst(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"a", "A"},
		{"", ""},
		{"ab", "Ab"},
		{"Ab", "Ab"},
		{"中文 support", "中文 support"},
		{"support 中文", "Support 中文"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.UpperFirst(tt.give))
	}
}

func TestLowerFirst(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"A", "a"},
		{"", ""},
		{"Ab", "ab"},
		{"ab", "ab"},
		{"中文 support", "中文 support"},
		{"Support 中文", "support 中文"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.LowerFirst(tt.give))
	}
}

func TestUpperWord(t *testing.T) {
	tests := []struct {
		give string
		want string
	}{
		{"a", "A"},
		{"", ""},
		{"ab", "Ab"},
		{"Ab", "Ab"},
		{"hi lo", "Hi Lo"},
		{"hi lo wr", "Hi Lo Wr"},
		{"!Test it!", "!Test It!"},
		{"This is a Test.", "This Is A Test."},
		{"TTtest TThis...good at This", "TTtest TThis...Good At This"},
		{"test...test...this...is..WOrk", "Test...Test...This...Is..WOrk"},
		{"试一试中文", "试一试中文"},
		{"中文也可以upper word", "中文也可以Upper Word"},
		{"文", "文"},
		{"wo...shi...中文", "Wo...Shi...中文"},
		{"...", "..."},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, strutil.UpperWord(tt.give))
	}
}

func TestSnakeCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"RangePrice":  "range_price",
		"rangePrice":  "range_price",
		"range_price": "range_price",
		"中文Snake":     "中文_snake",
	}

	for sample, want := range tests {
		is.Equal(want, strutil.SnakeCase(sample))
	}

	is.Equal("range-price", strutil.Snake("rangePrice", "-"))
	is.Equal("range price", strutil.SnakeCase("rangePrice", " "))
}

func TestCamelCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"rangePrice":   "rangePrice",
		"range_price":  "rangePrice",
		"_range_price": "RangePrice",
		"try中文":        "try中文",
		"_try中文":       "Try中文",
		"中文try":        "中文try",
		"中文_try":       "中文Try",
	}

	for sample, want := range tests {
		is.Equal(want, strutil.CamelCase(sample))
	}

	is.Equal("rangePrice", strutil.Camel("range-price", "-"))
	is.Equal("rangePrice", strutil.CamelCase("range price", " "))
	is.Equal("中文Try", strutil.CamelCase("中文 try", " "))

	// custom sep char
	is.Equal("rangePrice", strutil.CamelCase("range+price", "+"))
	is.Equal("rangePrice", strutil.CamelCase("range*price", "*"))
}
