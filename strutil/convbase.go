package strutil

import (
	"strconv"
	"strings"

	"github.com/gookit/goutil/x/basefn"
)

//
// -------------------- convert base --------------------
//

const (
	Base10Chars = "0123456789"
	Base16Chars = "0123456789abcdef"
	Base32Chars = "0123456789abcdefghjkmnpqrstvwxyz"
	Base36Chars = "0123456789abcdefghijklmnopqrstuvwxyz"
	Base48Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKL"
	Base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base64Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
)

// Base10Conv convert base10 string to new base string.
func Base10Conv(src string, to int) string { return BaseConv(src, 10, to) }

// BaseConv convert base string by from and to base.
//
// Note: from and to base must be in [2, 64]
//
// Usage:
//
//	BaseConv("123", 10, 16) // Output: "7b"
//	BaseConv("7b", 16, 10) // Output: "123"
func BaseConv(src string, from, to int) string {
	if from > 64 || from < 2 {
		from = 10
	}
	if to > 64 || to < 2 {
		to = 16
	}
	return BaseConvByTpl(src, Base64Chars[:from], Base64Chars[:to])
}

// BaseConvInt convert base int to new base string.
//
// Usage:
//
//	BaseConv(123, 16) // Output: "7b"
func BaseConvInt(src uint64, to int) string {
	if to > 64 || to < 2 {
		to = 16
	}
	return BaseConvIntByTpl(src, Base64Chars[:to])
}

// BaseConvByTpl convert base string by template.
//
// Usage:
//
//	BaseConvert("123", Base62Chars, Base16Chars) // Output: "1e"
//	BaseConvert("1e", Base16Chars, Base62Chars) // Output: "123"
func BaseConvByTpl(src string, fromBase, toBase string) string {
	if fromBase == toBase {
		return src
	}

	// convert to base 10
	var dec uint64
	if fromBase == Base10Chars {
		var err error
		dec, err = strconv.ParseUint(src, 10, 0)
		if err != nil {
			basefn.Panicf("input is not a valid decimal number: %s(%v)", src, err)
		}
	} else {
		fLen := uint64(len(fromBase))
		for _, c := range src {
			dec = dec*fLen + uint64(strings.IndexRune(fromBase, c))
		}
	}

	// convert to new base
	return BaseConvIntByTpl(dec, toBase)
}

// BaseConvIntByTpl convert base int to new base string.
func BaseConvIntByTpl(dec uint64, toBase string) string {
	// convert to new base
	var res string
	toLen := uint64(len(toBase))
	for dec > 0 {
		res = string(toBase[dec%toLen]) + res
		dec /= toLen
	}
	return res
}
