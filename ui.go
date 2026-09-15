package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
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

func NewCommandLine(input io.Reader, output io.Writer) *CommandLine {
	return &CommandLine{
		reader: bufio.NewReader(input),
		output: output,
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
