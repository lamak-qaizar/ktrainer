package main

import (
	"errors"
	"fmt"
	"testing"
)

func claudeCommand(prompt string) string {
	return fmt.Sprintf("claude -p %s --allowedTools Edit --safe-mode", prompt)
}

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	// "exit" input is added by runner
	Scenario{}.Run(t)
}

func TestRunApp_RedToGreenWhenTestFails(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write a test",
				Commands: []CommandExpected{
					{Command: claudeCommand("Prompt: write a test"), Mock: MockResponse{Output: "test edited"}},
					{Command: "dotnet test", Mock: MockResponse{Err: errors.New("")}},
				},
				ExpectPhase: Green,
			},
		},
	}.Run(t)
}

func TestRunApp_StayOnRedWhenTestPasses(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write a test",
				Commands: []CommandExpected{
					{Command: claudeCommand("Prompt: write a test"), Mock: MockResponse{Output: "test edited"}},
					{Command: "dotnet test", Mock: MockResponse{Output: ""}},
				},
				ExpectPhase: Red,
			},
		},
	}.Run(t)
}
