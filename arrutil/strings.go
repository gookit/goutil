package arrutil

import "strconv"

// StringsToSlice convert []string to []any
func StringsToSlice(ss []string) []any {
	args := make([]any, len(ss))
	for i, s := range ss {
		args[i] = s
	}
	return args
}

// StringsAsInts convert and ignore error
func StringsAsInts(ss []string) []int {
	ints, _ := StringsTryInts(ss)
	return ints
}

// StringsToInts string slice to int slice
func StringsToInts(ss []string) (ints []int, err error) {
	return StringsTryInts(ss)
}

// StringsTryInts string slice to int slice
func StringsTryInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		ints = append(ints, iVal)
	}
	return
}
