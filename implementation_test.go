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
	got,_ := PrefixEvaluation("+ 100 2")
	var want float64 = 100 + 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationSub(c *C) {
	got,_ := PrefixEvaluation("- 100 2")
	var want float64 = 100 - 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationMul(c *C) {
	got,_ := PrefixEvaluation("* 100 2")
	var want float64 = 100 * 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationDiv(c *C) {
	got,_ := PrefixEvaluation("/ 100 2")
	var want float64 = 100 / 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationPow(c *C) {
	got,_ := PrefixEvaluation("^ 100 2")
	var want float64 = math.Pow(float64(100), float64(2))
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluation1(c *C) {
	got,_ := PrefixEvaluation("+ 6 - * 7 8 ^ 2 4")
	var want float64 = 6 + 7 * 8 - math.Pow(float64(2), float64(4))
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluation2(c *C) {
	got,_ := PrefixEvaluation("/ - ^ 4 + * 1 4 * 4 17 3")
	var want float64 = (math.Pow(float64(4), float64(-1)) * 4 + 4 * 17) / 3
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluation3(c *C) {
	got,_ := PrefixEvaluation("+ / ^ 28 3 8 - * 16 / 3 2 ^ 28 0")
	var want float64 = math.Pow(float64(28), float64(3)) / 8 + 16 * 3 / 2 - math.Pow(float64(28), float64(0))
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluation4(c *C) {
	got,_ := PrefixEvaluation("- * 2 17 + / 6 * 3 4 1")
	var want float64 = 2 * 17 - 6 / 3 * 4 + 1
	c.Assert(got, Equals, want)
}

func ExamplePrefixEvaluation() {
	res, _ := PrefixEvaluation("* 3 + 2 2")
	fmt.Println(res)

	// Output:
	//12
}
