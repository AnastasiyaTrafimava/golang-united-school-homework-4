package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errorEmptyInput     = errors.New("input is empty")
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	if count(input) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	var first, second int

	a := strings.LastIndex(input, "+")
	if a == -1 {
		a = strings.LastIndex(input, "-")
	}

	if first, err = strconv.Atoi(strings.TrimLeft(input[0:a], "+")); err != nil {
		return "", fmt.Errorf("failed to convert %q: %w", input[0:a], err)
	}
	if second, err = strconv.Atoi(strings.TrimLeft(input[a:], "+")); err != nil {
		return "", fmt.Errorf("failed to convert %q: %w", input[a:], err)
	}

	return strconv.Itoa(first + second), nil

}

func count(input string) (count int) {
	input = strings.TrimLeft(input, "+")
	input = strings.TrimLeft(input, "-")

	for _, str := range strings.Split(input, "-") {
		count += len(strings.Split(str, "+"))
	}
	return
}
