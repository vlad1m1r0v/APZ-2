package lab2

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixEvaluation evaluates math expression written in prefix form
type fn func(float64, float64) float64

func add(a, b float64) float64 {
	return b + a
}

func sub(a, b float64) float64 {
	return b - a
}

func mul(a, b float64) float64 {
	return b * a
}

func div(a, b float64) float64 {
	return b / a
}

func pow(a, b float64) float64 {
	return math.Pow(b, a)
}

var fnMap = map[string]fn{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"^": pow,
}

func PrefixEvaluation(input string) (string, error) {
	var (
		stack       []float64
		values      = strings.Fields(input)
		operand, _  = regexp.Compile("^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$")
		operator, _ = regexp.Compile("^[+,*,/,^,-]$")
	)

	if input == "" || input == " " {
		return "", errors.New("empty string")
	}

	for i := len(values) - 1; i >= 0; i-- {
		var value string = values[i]

		if operand.MatchString(value) {
			var number, _ = strconv.ParseFloat(value, 64)
			stack = append(stack, number)

		} else if operator.MatchString(value) {

			if len(stack) < 2 {
				return "", errors.New("invalid input")
			}

			a, b := stack[len(stack)-1], stack[len(stack)-2]
			var fn = fnMap[value]
			stack = append(stack[:len(stack)-2], fn(b, a))

		} else {
			return "", errors.New("invalid argument")
		}
	}
	return fmt.Sprintf("%.0f", stack[0]), nil
}
