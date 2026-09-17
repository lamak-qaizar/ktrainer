package main

import (
	"bytes"
	"io"
	"os/exec"
	"reflect"
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
	calls []string
}

func NewTestCommandRunner() *TestCommandRunner {
	return &TestCommandRunner{calls: []string{}}
}

func (runner *TestCommandRunner) Run(name string, args []string, dir string, env []string, output io.Writer) (string, error) {
	command := strings.Join(append([]string{name}, args...), " ")
	runner.calls = append(runner.calls, command)
	return "", nil
}

func (runner *TestCommandRunner) AssertCommands(test *testing.T, want []string) {
	test.Helper()
	got := runner.calls
	if !reflect.DeepEqual(got, want) {
		test.Errorf("Expected commands: %q, got: %q", want, got)
	}
}
