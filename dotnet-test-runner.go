package main

import (
	"io"
)

type TestRunner interface {
	Run(dir string) (passed bool, output string)
}

type DotnetTestRunner struct {
	output io.Writer
	runner Runner
}

func NewDotnetTestRunner(output io.Writer, runner Runner) *DotnetTestRunner {
	return &DotnetTestRunner{output: output, runner: runner}
}

func (runner DotnetTestRunner) Run(dir string) (passed bool, output string) {
	output, err := runner.runner.Run("dotnet", []string{"test"}, dir, nil, io.Discard)
	return err == nil, output
}
