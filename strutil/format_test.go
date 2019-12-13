package strutil_test

import (
	"testing"

	"github.com/gookit/goutil/strutil"
	"github.com/stretchr/testify/assert"
)

func TestUpperOrLowerCase(t *testing.T) {
	// Uppercase, Lowercase
	assert.Equal(t, "ABC", strutil.Uppercase("abc"))
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
	}

	for sample, want := range tests {
		is.Equal(want, strutil.CamelCase(sample))
	}

	is.Equal("rangePrice", strutil.Camel("range-price", "-"))
	is.Equal("rangePrice", strutil.CamelCase("range price", " "))

	// custom sep char
	is.Equal("rangePrice", strutil.CamelCase("range+price", "+"))
	is.Equal("rangePrice", strutil.CamelCase("range*price", "*"))
}
