package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

//Getting hooked up with Gocheck
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPrefixEvaluation(c *C) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "+ 2 2", want: "4"},
		{input: "* 3 + 4 1", want: "15"},
		{input: "^ 2 * 4 7", want: "268435456"},
		{input: "^ 3 / 8 4", want: "9"},
		{input: "3 * 2 * + 4 4 ^ 3 / 8 4", want: "144"},
		{input: "* 6 + 2 * 3 + 2 ^ * + 1 3 4 5", want: "18874416"},
		{input: "+ + 4 5", want: "invalid input"},
		{input: "+ g 5", want: "invalid input"},
		{input: "+ + g 5", want: "invalid input"},
		{input: "gg", want: "invalid input"},
		{input: "", want: "invalid input"},
	}

	for _, test := range tests {
		got, err := PrefixEvaluation(test.input)
		if test.want == "invalid input" {
			c.Assert(err, Not(Equals), nil)
			c.Assert(got, Equals, "")
		} else {
			c.Assert(test.want, Equals, got)
		}
	}

}

func ExamplePrefixEvaluation() {
	res, _ := PrefixEvaluation("* 3 + 2 2")
	fmt.Println(res)

	// Output:
	//12
}
