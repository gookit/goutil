// Package mathutil provide math(int, number) util functions. eg: convert, math calc, random
package mathutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gookit/goutil/comdef"
	"github.com/gookit/goutil/internal/checkfn"
)

// Mul computes the `a*b` value, rounding the result.
func Mul[T1, T2 comdef.Number](a T1, b T2) float64 {
	return math.Round(SafeFloat(a) * SafeFloat(b))
}

// MulF2i computes the float64 type a * b value, rounding the result to an integer.
func MulF2i(a, b float64) int {
	return int(math.Round(a * b))
}

// Div computes the `a/b` value, result uses a round handle.
func Div[T1, T2 comdef.Number](a T1, b T2) float64 {
	return math.Round(SafeFloat(a) / SafeFloat(b))
}

// DivInt computes the int type a / b value, rounding the result to an integer.
func DivInt[T comdef.Integer](a, b T) int {
	fv := math.Round(float64(a) / float64(b))
	return int(fv)
}

// DivF2i computes the float64 type a / b value, rounding the result to an integer.
func DivF2i(a, b float64) int {
	return int(math.Round(a / b))
}

// Percent returns a value percentage of the total. eg: 1/100 = 1.0%
func Percent(val, total int) float64 {
	if total == 0 {
		return float64(0)
	}
	return (float64(val) / float64(total)) * 100
}

// Range a number range expression, and handle each value. eg: "1-100,123,124"
func Range(expr string, handle func(val int)) error {
	for _, item := range strings.Split(expr, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		// is range, eg: "1-100", "-20-2"
		if idx := checkfn.IndexByteAfter(item, '-', 1); idx > 0 {
			start, end, err := parseIntRange(item, idx)
			if err != nil {
				return err
			}

			// range number
			for i := start; i <= end; i++ {
				handle(i)
			}
		} else {
			iVal, err := strconv.Atoi(item)
			if err != nil {
				return fmt.Errorf("invalid integer value: %q", item)
			}
			handle(iVal)
		}
	}
	return nil
}

// Expand a number range expression to int[]. eg: "1-100,123,124"
func Expand(expr string) ([]int, error) {
	var nums []int
	for _, item := range strings.Split(expr, ",") {
		if item == "" {
			continue
		}

		// is range, eg: "1-100", "-20-2"
		if idx := checkfn.IndexByteAfter(item, '-', 1); idx > 0 {
			ints, err := expandIntRange(item, idx)
			if err != nil {
				return nil, err
			}
			nums = append(nums, ints...)
		} else {
			iVal, err := strconv.Atoi(item)
			if err != nil {
				return nil, fmt.Errorf("invalid integer value: %q", item)
			}
			nums = append(nums, iVal)
		}
	}
	return nums, nil
}

// 处理范围格式
// eg: "1-30" -> [1, 30], "-20-2" -> [-20, 2]
func parseIntRange(value string, sepIdx int) (min int, max int, err error) {
	start, end := value[:sepIdx], value[sepIdx+1:]

	min, err = strconv.Atoi(start)
	if err != nil {
		err = fmt.Errorf("invalid range start value: %q", start)
		return
	}

	max, err = strconv.Atoi(end)
	if err != nil {
		err = fmt.Errorf("invalid range end value: %q", end)
		return
	}

	// swap min and max
	if min > max {
		min, max = max, min
	}
	return
}

// 将 "1-30" 转换为 int 列表
func expandIntRange(value string, sepIdx int) ([]int, error) {
	var ints []int
	start, end, err := parseIntRange(value, sepIdx)
	if err != nil {
		return nil, err
	}

	for i := start; i <= end; i++ {
		ints = append(ints, i)
	}
	return ints, nil
}
