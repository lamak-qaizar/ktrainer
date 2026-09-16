package main

import (
	"strings"
	"testing"
)

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	input := strings.NewReader("exit\n")
	ui := NewTestUI(input)
	commandRunner := NewTestCommandRunner()

	runApp(ui, NewClaudeCLIAssistant(NewClaudeAuth(""), commandRunner),
		NewDotnetTestRunner(ui.Writer(), commandRunner))

	ui.AssertOutputContains(t, "Exiting kata trainer.")
	commandRunner.AssertCommands(t, []string{})
}
