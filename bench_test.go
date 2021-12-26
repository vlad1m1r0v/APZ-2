package lab2

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"unsafe"
)

type test struct {
	elements   int
	expression string
}

var tests []test

func pickOperation() string {
	var chance = rand.Float64()
	if chance < 0.25 {
		return "+"
	} else if chance < 0.5 {
		return "-"
	} else if chance < 0.75 {
		return "*"
	} else {
		return "/"
	}
}

func pickNumber() string {
	var value = rand.Intn(9) + 1
	return fmt.Sprintf("%v", value)
}

func Clone(s string) string {
	c := make([]byte, len(s))
	copy(c, s)
	return *(*string)(unsafe.Pointer(&c))
}

var UpperBound = int(math.Pow(2.0, 18.0))

func generateTests() {
	var prefixExpression string = "+ 2 2"
	for i := 1; i <= UpperBound; i *= 2 {
		for j := i; j < i*2 && j <= UpperBound; j++ {
			var num = pickNumber()
			var op = pickOperation()
			prefixExpression = fmt.Sprintf("%s %s ", op, num) + prefixExpression
		}
		var newTest = test{
			elements:   len(prefixExpression),
			expression: Clone(prefixExpression),
		}
		tests = append(tests, newTest)
	}
}

var result string

func BenchmarkPrefixEvaluation(b *testing.B) {
	generateTests()
	fmt.Println(UpperBound)
	for _, item := range tests {
		b.Run(fmt.Sprintf("%d", item.elements), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, _ = PrefixEvaluation(item.expression)
			}
		})
	}
}
