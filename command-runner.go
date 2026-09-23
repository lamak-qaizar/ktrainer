package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
)

type Runner interface {
	Run(name string, args []string, dir string, env []string, output io.Writer) (string, error)
}

type CommandRunner struct{}

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}

func (CommandRunner) Run(name string, args []string, dir string, env []string, output io.Writer) (string, error) {
	var buf bytes.Buffer
	writer := io.MultiWriter(&buf, output)

	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	if env != nil {
		cmd.Env = env
	}

	cmd.Stdout = writer
	cmd.Stderr = writer

	err := cmd.Run()
	return buf.String(), err
}

type TestCommandRunner struct {
	t        *testing.T
	expected []CommandExpected
	pos      int
}

func NewTestCommandRunner(t *testing.T, expected []CommandExpected) *TestCommandRunner {
	return &TestCommandRunner{t: t, expected: expected, pos: 0}
}

func (runner *TestCommandRunner) Run(name string, args []string, dir string, env []string, output io.Writer) (string, error) {
	command := strings.Join(append([]string{name}, args...), " ")

	if runner.pos >= len(runner.expected) {
		runner.t.Fatalf("Out of commands")
	}

	exp := runner.expected[runner.pos]
	if command != exp.Command {
		runner.t.Fatalf("Expected command:\n%q\nGot:\n%q", exp.Command, command)
	}

	runner.pos++
	return exp.Mock.Output, exp.Mock.Err
}
