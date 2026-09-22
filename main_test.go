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
				Input: "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_Prompt(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write a test",
				Commands: []CommandExpected{
					{Command: claudeCommand("Prompt: write a test"), Mock: MockResponse{Output: "test edited"}},
					{Command: "dotnet test", Mock: MockResponse{Output: "Tests passed."}},
				},
				ExpectPhase: Green,
			},
			{
				Input: "exit",
			},
		},
	}.Run(t)
}
