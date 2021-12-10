package ops

import "errors"

// addition operation
func add(a, b int) int {
	c := a + b
	return c
}

// subtraction operation
func sub(a, b int) int {
	c := a - b
	return c
}

// multiplier operation
func mp(a, b int) int {
	c := a * b
	return c
}

// division operation
func div(a, b int) (int, error) {
	var err error = nil
	var c int = 0
	if b == 0 {
		err = errors.New("ZERO DEVITION")
	} else {
		c = a / b
	}
	return c, err
}
