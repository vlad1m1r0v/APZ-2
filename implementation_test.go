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
	got,_ := PrefixEvaluation("+ 2 100")
	var want float64 = 100 + 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationSub(c *C) {
	got,_ := PrefixEvaluation("- 2 100")
	var want float64 = 100 - 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationMul(c *C) {
	got,_ := PrefixEvaluation("* 2 100")
	var want float64 = 100 * 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationDiv(c *C) {
	got,_ := PrefixEvaluation("/ 2 100")
	var want float64 = 100 / 2
	c.Assert(got, Equals, want)
}

func (s *TestSuite) TestPrefixEvaluationPow(c *C) {
	got,_ := PrefixEvaluation("^ 100 2")
	var want float64 = math.Pow(float64(100), float64(2))
	c.Assert(got, Equals, want)
}

func ExamplePrefixEvaluation() {
	res, _ := PrefixEvaluation("* 3 + 2 2")
	fmt.Println(res)

	// Output:
	//12
}
