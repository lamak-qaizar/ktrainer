package main

import (
	"errors"
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"strings"
	"testing"
)

type MockResponse struct {
	Output string
	Err    error
}

func NewErrorResponse() MockResponse {
	return MockResponse{Err: errors.New("")}
}

type CommandExpected struct {
	Command string
	Mock    MockResponse
}

type Step struct {
	Input       string
	Commands    []CommandExpected
	ExpectPhase tdd.Phase
}

type Scenario struct {
	Steps  []Step
	Model  string
	Effort string
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

func (scenario *Scenario) Phases() []tdd.Phase {
	var phases []tdd.Phase
	for _, step := range scenario.Steps {
		phases = append(phases, step.ExpectPhase)
	}
	return phases
}

func (scenario Scenario) Run(t *testing.T) {
	input := strings.NewReader(scenario.NewlineSeperatedInputs() + "exit\n")
	userInterface := ui.NewTestUI(input)
	commandRunner := NewTestCommandRunner(t, scenario.Commands())
	tdd := tdd.NewTestTDDPhase(t, scenario.Phases())

	kata, _ := NewKata("MarsRover", userInterface.Writer())
	NewAppRunner(*kata, userInterface,
		NewClaudeCLIAssistant(ConfigWithModelAndEffort(scenario.Model, scenario.Effort), commandRunner),
		NewDotnetTestRunner(userInterface.Writer(), commandRunner)).Run(tdd)
}
