package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/vlad1m1r0v/APZ-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func getInput(inputData *string, inputExpression *string) (io.Reader, error) {
	if *inputData != "" {
		input, err := os.Open(*inputData)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		return input, nil
	}
	return strings.NewReader(*inputExpression), nil
}

func getOutput(outputData *string) (io.Writer, error) {
	if *outputData != "" {
		output, err := os.Create(*outputData)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil, err
		}
		return output, nil
	}
	return os.Stdout, nil
}

func main() {
	flag.Parse()

	input, err := getInput(inputFile, inputExpression)

	if err != nil {
		return
	}

	output, err := getOutput(outputFile)

	if err != nil {
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err = handler.Compute()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
