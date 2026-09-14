package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type UI struct {
	reader *bufio.Reader
	output io.Writer
}

func NewUI(input io.Reader, output io.Writer) *UI {
	return &UI{
		reader: bufio.NewReader(input),
		output: output,
	}
}

func (ui *UI) Write(format string, a ...interface{}) {
	fmt.Fprintf(ui.output, format, a...)
}

func (ui *UI) Writeln(format string, a ...interface{}) {
	fmt.Fprintf(ui.output, format+"\n", a...)
}

func (ui *UI) Read() string {
	raw, _ := ui.reader.ReadString('\n')
	return strings.TrimSpace(raw)

}
