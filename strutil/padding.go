package strutil

import (
	"strings"

	"github.com/gookit/goutil/comdef"
)

// PosFlag type
type PosFlag = comdef.Position

// Position for padding/resize string
const (
	PosLeft PosFlag = iota
	PosMiddle
	PosRight
	PosAuto
)

/*************************************************************
 * String padding operation
 *************************************************************/

// Padding Fill the string to the specified length.
//
// params:
//   - s: 原始字符串
//   - pad: 用于填充的字符或字符串
//   - length: 目标长度
//   - padPos: 填充位置标志（左填充或右填充）
func Padding(s, pad string, length int, padPos PosFlag) string {
	return padding(s, pad, len(s), length, padPos)
}

// Utf8Padding Fill the string to the specified length. use utf8 width.
func Utf8Padding(s, pad string, wantLen int, padPos PosFlag) string {
	return padding(s, pad, Utf8Width(s), wantLen, padPos)
}

func padding(s, pad string, sLen, wantLen int, padPos PosFlag) string {
	diff := sLen - wantLen
	if diff >= 0 { // do not need padding.
		return s
	}

	// pad space.
	if len(s) == sLen && (pad == "" || pad == " ") {
		// Sprintf: 是按字数来填充的，不管中英文都是一个字符 - 有问题
		if padPos == PosRight { // to right
			return s + strings.Repeat(" ", -diff)
		}
		return strings.Repeat(" ", -diff) + s
	}

	// other character.
	if padPos == PosRight { // to right
		return s + Repeat(pad, -diff)
	}
	return Repeat(pad, -diff) + s
}

// PadLeft a string.
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, PosLeft)
}

// PadRight a string.
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, PosRight)
}

// PadChars padding a rune/byte to want length and with position flag
func PadChars[T byte | rune](cs []T, pad T, length int, pos PosFlag) []T {
	ln := len(cs)
	if ln >= length {
		ns := make([]T, length)
		copy(ns, cs[:length])
		return ns
	}

	idx := length - ln
	ns := make([]T, length)
	if pos == PosRight {
		copy(ns, cs)
		for i := ln; i < length; i++ {
			ns[i] = pad
		}
		return ns
	}

	// to left
	for i := 0; i < idx; i++ {
		ns[i] = pad
	}
	copy(ns[idx:], cs)
	return ns
}

// PadBytes padding a byte to want length and with position flag
func PadBytes(bs []byte, pad byte, length int, pos PosFlag) []byte {
	return PadChars(bs, pad, length, pos)
}

// PadBytesLeft a byte to want length
func PadBytesLeft(bs []byte, pad byte, length int) []byte {
	return PadChars(bs, pad, length, PosLeft)
}

// PadBytesRight a byte to want length
func PadBytesRight(bs []byte, pad byte, length int) []byte {
	return PadChars(bs, pad, length, PosRight)
}

// PadRunes padding a rune to want length and with position flag
func PadRunes(rs []rune, pad rune, length int, pos PosFlag) []rune {
	return PadChars(rs, pad, length, pos)
}

// PadRunesLeft a rune to want length
func PadRunesLeft(rs []rune, pad rune, length int) []rune {
	return PadChars(rs, pad, length, PosLeft)
}

// PadRunesRight a rune to want length
func PadRunesRight(rs []rune, pad rune, length int) []rune {
	return PadChars(rs, pad, length, PosRight)
}

// Align a string by given length and align settings. alias of Resize
func Align(s string, length int, align comdef.Align) string {
	return resize(s, len(s), length, align, false)
}

// Utf8Align a string by given length and align settings. alias of Utf8Resize
func Utf8Align(s string, length int, align comdef.Align) string {
	return resize(s, Utf8Width(s), length, align, false)
}

// Resize a string by given length and align settings. use padding space.
// If len(s) > wantLen, will truncate it.
func Resize(s string, length int, align comdef.Align) string {
	return resize(s, len(s), length, align, true)
}

// Utf8Resize a string by given length and align settings. use padding space.
// If width(s) > wantLen, will truncate it.
func Utf8Resize(s string, length int, align comdef.Align) string {
	return resize(s, Utf8Width(s), length, align, true)
}

// resize a string by given length and align settings. use padding space.
func resize(s string, sLen, wantLen int, align comdef.Align, cutOverflow bool) string {
	diff := sLen - wantLen
	if diff >= 0 { // do not need padding.
		// cutOverflow: truncate on sLen > wantLen
		if cutOverflow && sLen > wantLen {
			if len(s) == sLen {
				return s[:wantLen]
			}
			return utf8Truncate(s, sLen, wantLen, "")
		}
		return s
	}

	if align == comdef.Center {
		padLn := (wantLen - sLen) / 2
		if diff = wantLen - padLn*2; diff > 0 {
			s += " "
		}
		padStr := string(RepeatBytes(' ', padLn))
		return padStr + s + padStr
	}

	padStr := string(RepeatBytes(' ', wantLen-sLen))
	// tip: 左对齐 - 使用空白填充右边
	if align == comdef.Left || align == comdef.PosAuto {
		return s + padStr
	}
	return padStr + s
}

/*************************************************************
 * String repeat operation
 *************************************************************/

// Repeat a string by given times.
func Repeat(s string, times int) string {
	if times < 1 {
		return ""
	}
	if times == 1 {
		return s
	}

	var sb strings.Builder
	sb.Grow(len(s) * times)

	for i := 0; i < times; i++ {
		sb.WriteString(s)
	}
	return sb.String()
}

// RepeatRune repeat a rune char.
func RepeatRune(char rune, times int) []rune { return RepeatChars(char, times) }

// RepeatBytes repeat a byte char.
func RepeatBytes(char byte, times int) []byte { return RepeatChars(char, times) }

// RepeatChars repeat a byte char.
func RepeatChars[T byte | rune](char T, times int) []T {
	if times <= 0 {
		return make([]T, 0)
	}

	chars := make([]T, times)
	for i := 0; i < times; i++ {
		chars[i] = char
	}
	return chars
}
