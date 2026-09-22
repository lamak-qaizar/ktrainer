package main

import (
	"fmt"
	"testing"
)

func claudeCommand(prompt string) string {
	return fmt.Sprintf("claude -p %s --allowedTools Edit --safe-mode", prompt)
}

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input:       "exit\n",
				Commands:    []CommandExpected{},
				ExpectPhase: Red,
			},
		},
	}.Run(t)
}

// func TestRunApp_Prompt(t *testing.T) {
// 	input := strings.NewReader("write a test\n" + "exit\n")
// 	ui := NewTestUI(input)
// 	commandRunner := NewTestCommandRunner()

// 	runApp(ui, NewClaudeCLIAssistant(NewClaudeAuth(""), commandRunner),
// 		NewDotnetTestRunner(ui.Writer(), commandRunner))

// 	commandRunner.AssertCommands(t, []string{
// 		claudeCommand("Prompt: write a test"),
// 		"dotnet test"})
// }
