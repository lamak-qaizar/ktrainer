package main

import (
	"os"
	"os/exec"
)

type TestRunner interface {
	Run(dir string) (passed bool)
}

type DotnetTestRunner struct{}

func NewDotnetTestRunner() *DotnetTestRunner {
	return &DotnetTestRunner{}
}

func (runner DotnetTestRunner) Run(dir string) (passed bool) {
	cmd := exec.Command("dotnet", "test")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}
