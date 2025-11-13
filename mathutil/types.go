package mathutil

// ToIntFunc convert value to int
type ToIntFunc func(any) (int, error)

// ToInt64Func convert value to int64
type ToInt64Func func(any) (int64, error)

// ToUintFunc convert value to uint
type ToUintFunc func(any) (uint, error)

// ToUint64Func convert value to uint
type ToUint64Func func(any) (uint64, error)

// ToFloatFunc convert value to float
type ToFloatFunc func(any) (float64, error)

// ToTypeFunc convert value to defined type
type ToTypeFunc[T any] func(any) (T, error)
