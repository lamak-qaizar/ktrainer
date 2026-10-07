package test

import (
	"kata-trainer/app"
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"strings"
	"testing"
)

type Step struct {
	Input       string
	Commands    []app.CommandExpected
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

func (scenario *Scenario) Commands() []app.CommandExpected {
	var commands []app.CommandExpected
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
	commandRunner := app.NewTestCommandRunner(t, scenario.Commands())
	tdd := tdd.NewTestTDDPhase(t, scenario.Phases())

	kata, _ := app.NewKata("../MarsRover", userInterface)
	app.NewAppRunner(*kata, userInterface,
		app.NewClaudeCLIAssistant(app.ConfigWithModelAndEffort(scenario.Model, scenario.Effort), commandRunner),
		app.NewDotnetTestRunner(userInterface.Writer(), commandRunner)).Run(tdd)
}
