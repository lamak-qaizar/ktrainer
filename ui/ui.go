package ui

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type UserInterface interface {
	Write(str string)
	Writeln(str string)
	WriteColoured(str string, color Color)
	SenseiSays(str string)
	Title(str string)
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

func (ui *UI) SenseiSays(str string) {
	says := pterm.DefaultBox.WithHorizontalPadding(2).WithVerticalPadding(1).WithTitleTopCenter().WithTitle(pterm.Cyan(" sensei says ")).Sprint(str)
	ui.Write("\n" + says + "\n")
}

func (ui *UI) Title(str string) {
	title, _ := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(str, YELLOW.toPterm().ToStyle()),
	).Srender()
	ui.Write("\n" + title + "\n")
}

func (ui *UI) Write(str string) {
	fmt.Fprint(ui.output, str)
}

func (ui *UI) Writeln(str string) {
	ui.Write(str + "\n")
}

func (ui *UI) WriteColoured(str string, color Color) {
	ui.Write(color.toPterm().Sprint(str))
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
	pterm.SetDefaultOutput(&buf)
	pterm.DefaultSpinner.Writer = &buf
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
