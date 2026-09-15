package main

import (
	"io"
)

type TestRunner interface {
	Run(dir string) (passed bool)
}

type DotnetTestRunner struct {
	output        io.Writer
	commandRunner CommandRunner
}

func NewDotnetTestRunner(output io.Writer) *DotnetTestRunner {
	return &DotnetTestRunner{output: output, commandRunner: *NewCommandRunner()}
}

func (runner DotnetTestRunner) Run(dir string) (passed bool) {
	_, err := runner.commandRunner.Run("dotnet", []string{"test"}, dir, nil, runner.output)
	return err == nil
}
