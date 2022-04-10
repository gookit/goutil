package errorx

import "errors"

/*************************************************************
 * helper func for error
 *************************************************************/

// Cause returns the first cause error by call err.Cause().
// Otherwise, will returns current error.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if err, ok := err.(Causer); ok {
		return err.Cause()
	}
	return err
}

// Unwrap returns previous error by call err.Unwrap().
// Otherwise, will returns nil.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if err, ok := err.(Unwrapper); ok {
		return err.Unwrap()
	}
	return nil
}

// Previous alias of Unwrap()
func Previous(err error) error { return Unwrap(err) }

// Has check err has contains target, or err is eq target.
// alias of errors.Is()
func Has(err, target error) bool {
	return errors.Is(err, target)
}

// Is alias of errors.Is()
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// To try convert err to target, returns is result.
//
// NOTICE: target must be ptr and not nil
//
// alias of errors.As()
func To(err error, target interface{}) bool {
	return errors.As(err, target)
}

// As alias of errors.As()
//
// NOTICE: target must be ptr and not nil
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
