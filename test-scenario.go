package main

import (
	"strings"
	"testing"
)

type MockResponse struct {
	Output string
	Err    error
}

type CommandExpected struct {
	Command string
	Mock    MockResponse
}

type Step struct {
	Input       string
	Commands    []CommandExpected
	ExpectPhase Phase
}

type Scenario struct {
	Steps []Step
}

func (scenario *Scenario) NewlineSeperatedInputs() string {
	if len(scenario.Steps) == 0 {
		return ""
	}

	var inputs []string
	for _, step := range scenario.Steps {
		inputs = append(inputs, step.Input)
	}
	return strings.Join(inputs, "\n") + "\n"
}

func (scenario *Scenario) Commands() []CommandExpected {
	var commands []CommandExpected
	for _, step := range scenario.Steps {
		commands = append(commands, step.Commands...)
	}
	return commands
}

func (scenario *Scenario) Phases() []Phase {
	var phases []Phase
	for _, step := range scenario.Steps {
		phases = append(phases, step.ExpectPhase)
	}
	return phases
}

func (scenario Scenario) Run(t *testing.T) {
	input := strings.NewReader(scenario.NewlineSeperatedInputs() + "exit\n")
	ui := NewTestUI(input)
	commandRunner := NewTestCommandRunner(t, scenario.Commands())
	tdd := NewTestTDDStateMachine(t, scenario.Phases())

	runApp(ui,
		NewClaudeCLIAssistant(NewClaudeAuth(""), commandRunner),
		NewDotnetTestRunner(ui.Writer(), commandRunner),
		tdd)
}
