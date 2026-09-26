package main

import (
	"fmt"
	"kata-trainer/tdd"
	"testing"
)

func claudeCmd(template string, userPrompt string) string {
	prompt := tdd.PromptFrom(template, userPrompt)
	return fmt.Sprintf("claude -p %s --allowedTools Edit,Create --safe-mode --model haiku --permission-mode acceptEdits", prompt)
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
				ExpectPhase: tdd.Green,
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
				ExpectPhase: tdd.Red,
			},
		},
	}.Run(t)
}

func TestRunApp_StayOnRedWhenClaudeRejectsPrompt(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Input: "write all tests",
				Commands: []CommandExpected{
					{Command: claudeCmd("red.tmpl", "write all tests"), Mock: MockResponse{Output: "REJECTED: write one test only"}},
				},
				ExpectPhase: tdd.Red,
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
				ExpectPhase: tdd.Green,
			},
			{
				Input: "pass the test",
				Commands: []CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
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
				ExpectPhase: tdd.Green,
			},
			{
				Input: "pass the test",
				Commands: []CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{
				Input: "refactor the code",
				Commands: []CommandExpected{
					{Command: claudeCmd("refactor.tmpl", "refactor the code")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{
				Input:       "next",
				ExpectPhase: tdd.Red,
			},
		},
	}.Run(t)
}
