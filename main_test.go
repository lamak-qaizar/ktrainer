package main

import (
	"fmt"
	"testing"
)

func claudeCmd(template string, userPrompt string) string {
	prompt := promptFrom(template, userPrompt)
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
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: NewErrorResponse()},
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
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: Red,
			},
		},
	}.Run(t)
}

func TestRunApp_RefactorAfterTestsPassOnGreen(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write a test",
				Commands: []CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: NewErrorResponse()},
				},
				ExpectPhase: Green,
			},
			{
				Input: "pass the test",
				Commands: []CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: Refactor,
			},
		},
	}.Run(t)
}

func TestRunApp_RefactorStaysOnRefactorUnlessUserExplicitySelectsNext(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write a test",
				Commands: []CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: NewErrorResponse()},
				},
				ExpectPhase: Green,
			},
			{
				Input: "pass the test",
				Commands: []CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: Refactor,
			},
			{
				Input: "refactor the code",
				Commands: []CommandExpected{
					{Command: claudeCmd("refactor.tmpl", "refactor the code")},
					{Command: "dotnet test"},
				},
				ExpectPhase: Refactor,
			},
			{
				Input: "next",
				Commands: []CommandExpected{
					{Command: "dotnet test"},
				},
				ExpectPhase: Red,
			},
		},
	}.Run(t)
}
