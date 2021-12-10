package ops

import (
	"errors"
	"strconv"
	"strings"
)

// Handles the result making process
func Handle(rawInput string) (string, error) {
	var output string = "EMPTY"
	var a, b, c int = 0, 0, 0
	var err error = nil
	input := strings.TrimSpace(rawInput)
	// Mux
	if strings.Contains(input, "+") {
		a, b, err = splitConvString(input, "+")
		if err == nil {
			c = add(a, b)
		}
	} else if strings.Contains(input, "-") {
		a, b, err = splitConvString(input, "-")
		if err == nil {
			c = sub(a, b)
		}
	} else if strings.Contains(input, "*") {
		a, b, err = splitConvString(input, "*")
		if err == nil {
			c = mp(a, b)
		}
	} else if strings.Contains(input, "/") {
		a, b, err = splitConvString(input, "/")
		if err == nil {
			c, err = div(a, b)
		}
	} else {
		err = errors.New("NO VALIED OPS")
	}
	if err == nil {
		strVal := strconv.Itoa(c)
		output = "Result: " + strVal
	}
	return output, err
}

// splits and converts values from str to int
func splitConvString(input, operand string) (int, int, error) {
	var a, b int = 0, 0
	splitedInput := strings.Split(input, operand)
	if len(splitedInput) != 2 {
		err := errors.New("MULTI OPS")
		return a, b, err
	}
	a, err1 := strconv.Atoi(splitedInput[0])
	if err1 != nil {
		return 0, 0, err1
	}
	b, err2 := strconv.Atoi(splitedInput[1])
	if err2 != nil {
		return 0, 0, err2
	}
	return a, b, nil
}
