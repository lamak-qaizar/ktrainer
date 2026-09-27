package ui

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type Style struct {
	Display Display
	Color   Color
}

type UserInterface interface {
	Write(str string, style Style)
	Writeln(str string, style Style)
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

func (ui *UI) Write(str string, style Style) {
	if style.Display == TITLE {
		pterm.DefaultBigText.WithLetters(
			putils.LettersFromStringWithStyle(str, style.PtermColor().ToStyle()),
		).Render()
		return
	}

	if style.Display == BOX {
		pterm.Println()
		pterm.DefaultBox.WithTitleTopCenter().WithTitle("sensei says").WithTextStyle(style.PtermColor().ToStyle()).Print(str)
		return
	}

	style.PtermColor().Print(str)
}

func (ui *UI) SenseiSays(str string) {
	pterm.Println()
	pterm.DefaultBox.WithTitleTopCenter().WithTitle("sensei says").Print(str)
	pterm.Println()
}

func (ui *UI) Title(str string) {
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle(str, YELLOW.toPterm().ToStyle()),
	).Render()
}

func (ui *UI) Writeln(str string, style Style) {
	ui.Write(str+"\n", style)
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
