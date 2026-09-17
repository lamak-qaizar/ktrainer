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

type UserInterface interface {
	Write(format string)
	Writeln(format string)
	Read() string
}

type UI struct {
	reader *bufio.Reader
	output io.Writer
}

func NewUI() *UI {
	return &UI{
		reader: bufio.NewReader(os.Stdin),
		output: os.Stdout,
	}
}

func (ui *UI) Write(str string) {
	fmt.Fprintf(ui.output, "%s", str)
}

func (ui *UI) Writeln(str string) {
	fmt.Fprintf(ui.output, "%s", str+"\n")
}

func (ui *UI) Read() string {
	raw, _ := ui.reader.ReadString('\n')
	return strings.TrimSpace(raw)
}

func (ui *UI) Writer() io.Writer {
	return ui.output
}

type TestUI struct {
	*UI
	output *bytes.Buffer
}

func NewTestUI(input io.Reader) *TestUI {
	var buf bytes.Buffer
	return &TestUI{
		UI:     &UI{reader: bufio.NewReader(input), output: &buf},
		output: &buf,
	}
}

func (t *TestUI) AssertOutputContains(test *testing.T, want string) {
	test.Helper()
	got := t.output.String()
	if !strings.Contains(got, want) {
		test.Errorf("Expected output to contain %q, got: %s", want, got)
	}
}
