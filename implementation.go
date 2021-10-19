package lab2

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	var (
		stack       []string
		firstArg    int
		secondArg   int
		result      string
		values      = strings.Fields(input)
		value       string
		digit, _    = regexp.Compile("^[0-9]+$")
		operator, _ = regexp.Compile("^[+,*,/,^,-]$")
	)

	for i := len(values) - 1; i >= 0; i-- {
		value = values[i]

		if digit.MatchString(value) {
			stack = append(stack, value)

		} else if operator.MatchString(value) {

			if len(stack) < 2 {
				return "", errors.New("invalid input")
			}

			firstArg, _ = strconv.Atoi(stack[len(stack)-1])
			secondArg, _ = strconv.Atoi(stack[len(stack)-2])

			if value == "+" {
				result = strconv.Itoa(firstArg + secondArg)
			} else if value == "-" {
				result = strconv.Itoa(firstArg - secondArg)
			} else if value == "*" {
				result = strconv.Itoa(firstArg * secondArg)
			} else if value == "/" {
				result = strconv.Itoa(firstArg / secondArg)
			} else if value == "^" {
				result = strconv.Itoa(int(math.Pow(float64(firstArg), float64(secondArg))))
			}

			stack = stack[:len(stack)-2]
			stack = append(stack, result)

		} else {
			return "", errors.New("invalid input")
		}

	}

	return stack[0], nil
}
