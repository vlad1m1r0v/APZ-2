package lab2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	fileScanner := bufio.NewScanner(ch.Input)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		res, err := PrefixEvaluation(input)
		if err != nil {
			return err
		}
		ch.Output.Write([]byte(res))
	}

	return nil
}
