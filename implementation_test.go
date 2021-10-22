package lab2

import (
	"fmt"
	"math"
	"testing"

	. "gopkg.in/check.v1"
)

//Getting hooked up with Gocheck
func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPrefixEvaluationAdd(c *C) {
	got, _ := PrefixEvaluation("+ 100 2")
	var want float64 = 100 + 2
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluationSub(c *C) {
	got, _ := PrefixEvaluation("- 100 2")
	var want float64 = 100 - 2
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluationMul(c *C) {
	got, _ := PrefixEvaluation("* 100 2")
	var want float64 = 100 * 2
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluationDiv(c *C) {
	got, _ := PrefixEvaluation("/ 100 2")
	var want float64 = 100 / 2
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluationPow(c *C) {
	got, _ := PrefixEvaluation("^ 100 2")
	var want float64 = math.Pow(float64(100), float64(2))
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluation1(c *C) {
	got, _ := PrefixEvaluation("+ 6 - * 7 8 ^ 2 4")
	var want float64 = 6 + 7*8 - math.Pow(float64(2), float64(4))
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluation2(c *C) {
	got, _ := PrefixEvaluation("^ 2 / 16 - 8196 * 2 ^ 2 * 6 2")
	var want float64 = 16
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluation3(c *C) {
	got, _ := PrefixEvaluation("+ / ^ 28 3 8 - * 16 / 3 2 ^ 28 0")
	var want float64 = math.Pow(float64(28), float64(3))/8 + 16*3/2 - math.Pow(float64(28), float64(0))
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestPrefixEvaluation4(c *C) {
	got, _ := PrefixEvaluation("^ 2 + 8 / 624 * 78 / 96 * 3 + 6 2")
	var want float64 = 1024
	c.Assert(got, Equals, fmt.Sprintf("%.0f", want))
}

func (s *TestSuite) TestInvalidInput(c *C) {
	_, err := PrefixEvaluation("+ - * ^ 2 2")
	var want = "invalid input"
	if err != nil {
		c.Assert(fmt.Sprint(err), Equals, want)
	}
}

func (s *TestSuite) TestInvalidArgument(c *C) {
	_, err := PrefixEvaluation("* a c")
	var want = "invalid argument"
	if err != nil {
		c.Assert(fmt.Sprint(err), Equals, want)
	}
}

func (s *TestSuite) TestEmptyString(c *C) {
	_, err := PrefixEvaluation("")
	var want = "empty string"
	if err != nil {
		c.Assert(fmt.Sprint(err), Equals, want)
	}
}

func ExamplePrefixEvaluation() {
	res, _ := PrefixEvaluation("* 3 + 2 2")
	fmt.Println(res)

	// Output:
	//12
}
