package fizzbuzz

import "errors"

var (
	ErrInvalidDivisor = errors.New("int1 and int2 must be positive integers")
	ErrInvalidLimit   = errors.New("limit must be a positive integer")
)