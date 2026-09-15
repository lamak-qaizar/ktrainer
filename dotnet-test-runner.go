package main

import (
	"io"
	"os/exec"
)

type TestRunner interface {
	Run(dir string) (passed bool)
}

type DotnetTestRunner struct {
	output io.Writer
}

func NewDotnetTestRunner(output io.Writer) *DotnetTestRunner {
	return &DotnetTestRunner{output: output}
}

func (runner DotnetTestRunner) Run(dir string) (passed bool) {
	cmd := exec.Command("dotnet", "test")
	cmd.Dir = dir
	cmd.Stdout = runner.output
	cmd.Stderr = runner.output
	err := cmd.Run()
	return err == nil
}
