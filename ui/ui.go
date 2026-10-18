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
	Spinner(str string) *Spinner
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

func (ui *UI) Spinner(str string) *Spinner {
	ui.Writeln("")
	return NewSpinner(str)
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
	output     *bytes.Buffer
	t          *testing.T
	senseiSaid []string
	pos        int
}

func NewTestUI(input io.Reader, t *testing.T, senseiSaid []string) *TestUI {
	var buf bytes.Buffer
	pterm.SetDefaultOutput(&buf)
	pterm.DefaultSpinner.Writer = &buf
	return &TestUI{
		UI:         &UI{reader: bufio.NewReader(input), output: &buf},
		output:     &buf,
		t:          t,
		senseiSaid: senseiSaid,
		pos:        0,
	}
}

func (ui *TestUI) SenseiSays(str string) {
	if ui.pos >= len(ui.senseiSaid) {
		ui.t.Fatalf("Sensei is out of things to say")
	}

	exp := ui.senseiSaid[ui.pos]
	if !strings.Contains(str, exp) {
		ui.t.Fatalf("Expected sensei to say: %q, got: %q, step: %d", exp, str, ui.pos)
	}

	ui.pos++
}

func (ui *TestUI) VerifyAllReceived() {
	if ui.pos < len(ui.senseiSaid) {
		ui.t.Fatalf("Sensei did not say all that was expected to be said.")
	}
}
