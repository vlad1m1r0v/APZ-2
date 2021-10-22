package lab2

import (
	"bytes"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestInputFile(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("^ 2 / 16 - 8196 * 2 ^ 2 * 6 2"),
		Output: &buff,
	}
	var want = "16"

	handler.Compute()
	c.Assert(buff.String(), Equals, want)
}

func (s *TestSuite) TestInputErr(c *C) {
	var buff bytes.Buffer

	handler := &ComputeHandler{
		Input:  strings.NewReader("* a c"),
		Output: &buff,
	}
	want := "invalid argument"

	err := handler.Compute()
	c.Assert(fmt.Sprint(err), Equals, want)
}