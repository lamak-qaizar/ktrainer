package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

type UI interface {
	Write(format string, a ...interface{})
	Writeln(format string, a ...interface{})
	Read() string
}

type CommandLine struct {
	reader *bufio.Reader
	output io.Writer
}

func NewCommandLine() *CommandLine {
	return &CommandLine{
		reader: bufio.NewReader(os.Stdin),
		output: os.Stdout,
	}
}

func (ui *CommandLine) Write(format string, a ...interface{}) {
	fmt.Fprintf(ui.output, format, a...)
}

func (ui *CommandLine) Writeln(format string, a ...interface{}) {
	fmt.Fprintf(ui.output, format+"\n", a...)
}

func (ui *CommandLine) Read() string {
	raw, _ := ui.reader.ReadString('\n')
	return strings.TrimSpace(raw)
}

func (ui *CommandLine) Writer() io.Writer {
	return ui.output
}

type TestCommandLine struct {
	*CommandLine
	output *bytes.Buffer
}

func NewTestCommandLine(input io.Reader) *TestCommandLine {
	var buf bytes.Buffer
	return &TestCommandLine{
		CommandLine: &CommandLine{reader: bufio.NewReader(input), output: &buf},
		output:      &buf,
	}
}

func (t *TestCommandLine) AssertOutputContains(test *testing.T, want string) {
	test.Helper()
	got := t.output.String()
	if !strings.Contains(got, want) {
		test.Errorf("Expected output to contain %q, got: %s", want, got)
	}
}
