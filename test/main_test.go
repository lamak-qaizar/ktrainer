package test

import (
	"fmt"
	"kata-trainer/app"
	"kata-trainer/tdd"
	"kata-trainer/test/prompts"
	"testing"
)

func claudeCmd(template string, userPrompt string) string {
	prompt := prompts.PromptFrom(template, userPrompt)
	return fmt.Sprintf("claude -p %s --allowedTools Edit,Create --safe-mode --permission-mode acceptEdits --model haiku", prompt)
}

func claudeCmdWithModelAndEffort(template string, userPrompt string, model string, effort string) string {
	prompt := tdd.PromptFrom(template, userPrompt)
	return fmt.Sprintf("claude -p %s --allowedTools Edit,Create --safe-mode --permission-mode acceptEdits --model %s --effort %s", prompt, model, effort)
}

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_RedToGreenWhenTestFails(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.",
				Input:  "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_StayOnRedWhenTestPasses(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Red,
			},
			{
				Input: "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_StayOnRedWhenClaudeRejectsPrompt(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write all tests",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write all tests"), Mock: app.MockResponse{Output: "REJECTED: write one test only"}},
				},
				ExpectPhase: tdd.Red,
			},
			{
				Input: "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_RefactorAfterTestsPassOnGreen(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.",
				Input:  "pass the test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{
				Sensei: "Specify design improvements",
				Input:  "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_RefactorStaysOnRefactorUnlessUserExplicitySelectsNext(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.",
				Input:  "pass the test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{

				Sensei: "Specify design improvements",
				Input:  "refactor the code",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("refactor.tmpl", "refactor the code")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{
				Input:       "next",
				ExpectPhase: tdd.Red,
			},
			{
				Sensei: "Describe the failing test to write.",
				Input:  "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_ModelAndEffortShouldReflectInClaudeCommand(t *testing.T) {
	Scenario{
		Model:  "sonnet",
		Effort: "low",
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmdWithModelAndEffort("red.tmpl", "write a test", "sonnet", "low")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.",
				Input:  "exit",
			},
		},
	}.Run(t)
}

func TestRunApp_FirstPromptShouldBeMoreDescriptive(t *testing.T) {
	Scenario{
		Steps: []Step{
			{
				Sensei: "Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.",
				Input:  "pass the test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{

				Sensei:      "Specify design improvements",
				Input:       "next",
				ExpectPhase: tdd.Red,
			},
			{
				Sensei: "Describe the failing test to write.",
				Input:  "write a test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("red.tmpl", "write a test")},
					{Command: "dotnet test", Mock: app.NewErrorResponse()},
				},
				ExpectPhase: tdd.Green,
			},
			{
				Sensei: "How would you like to pass the test?",
				Input:  "pass the test",
				Commands: []app.CommandExpected{
					{Command: claudeCmd("green.tmpl", "pass the test")},
					{Command: "dotnet test"},
				},
				ExpectPhase: tdd.Refactor,
			},
			{
				Sensei: "Specify design improvements",
				Input:  "exit",
			},
		},
	}.Run(t)
}
