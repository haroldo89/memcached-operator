package adder

import "errors"

var ErrInvalidSummand = errors.New("invalid summand")

func Add(x, y int) (int, error) {
	if x <= 0 || y <= 0 {
		return 0, ErrInvalidSummand
	}
	return x + y, nil
}
