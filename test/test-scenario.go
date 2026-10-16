package test

import (
	"kata-trainer/app"
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"strings"
	"testing"
)

type Step struct {
	Sensei      string
	Input       string
	Commands    []app.CommandExpected
	ExpectPhase tdd.Transition
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

func (scenario *Scenario) Phases() []tdd.Transition {
	var transitions []tdd.Transition
	for _, step := range scenario.Steps {
		transitions = append(transitions, step.ExpectPhase)
	}
	return transitions
}

func (scenario *Scenario) SenseiSaid() []string {
	var senseiSaid []string
	for _, step := range scenario.Steps {
		if step.Sensei != "" {
			senseiSaid = append(senseiSaid, step.Sensei)
		}
	}
	return senseiSaid
}

func (scenario Scenario) Run(t *testing.T) {
	input := strings.NewReader(scenario.NewlineSeperatedInputs())
	userInterface := ui.NewTestUI(input, t, scenario.SenseiSaid())
	commandRunner := app.NewTestCommandRunner(t, scenario.Commands())
	tdd := tdd.NewTestTDDPhase(t, scenario.Phases())

	kata, _ := app.NewKata("../MarsRover", userInterface)
	app.NewAppRunner(*kata, userInterface,
		app.NewClaudeCLIAssistant(app.ConfigWithModelAndEffort(scenario.Model, scenario.Effort), commandRunner),
		app.NewDotnetTestRunner(userInterface.Writer(), commandRunner)).Run(tdd)

	userInterface.VerifyAllReceived()
}
