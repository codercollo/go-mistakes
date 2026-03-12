// Package mathutil provides basic arithmetic helper functions.
//
// All functions operate on float64 values and return errors
// for invalid inputs such as division by zero.
package mathutil

import "errors"

// ErrDivByZero is returned when a division by zero is attempted.
var ErrDivByZero = errors.New("division by zero")

// Divide returns the result of dividing a by b.
// Returns ErrDivByZero if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}
