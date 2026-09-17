package main

import (
	"fmt"
	"strings"
	"testing"
)

func claudeCommand(prompt string) string {
	return fmt.Sprintf("claude -p %s --allowedTools Edit --safe-mode", prompt)
}

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	input := strings.NewReader("exit\n")
	ui := NewTestUI(input)
	commandRunner := NewTestCommandRunner()

	runApp(ui, NewClaudeCLIAssistant(NewClaudeAuth(""), commandRunner),
		NewDotnetTestRunner(ui.Writer(), commandRunner))

	ui.AssertOutputContains(t, "Exiting kata trainer.")
	commandRunner.AssertCommands(t, []string{})
}

func TestRunApp_Prompt(t *testing.T) {
	input := strings.NewReader("write a test\n" + "exit\n")
	ui := NewTestUI(input)
	commandRunner := NewTestCommandRunner()

	runApp(ui, NewClaudeCLIAssistant(NewClaudeAuth(""), commandRunner),
		NewDotnetTestRunner(ui.Writer(), commandRunner))

	commandRunner.AssertCommands(t, []string{
		claudeCommand("Prompt: write a test"),
		"dotnet test"})
}
