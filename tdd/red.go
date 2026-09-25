package tdd

import (
	"kata-trainer/ui"
)

type RedState struct {
}

func (red *RedState) Phase() Phase {
	return Red
}

func (red *RedState) Advance(testsPassed bool) TDDPhase {
	if testsPassed {
		return red
	}

	return &GreenState{}
}

func (red *RedState) Prompt(userInput string) string {
	return PromptFrom("red.tmpl", userInput)
}

func (red *RedState) ForceAdvance() TDDPhase {
	return &GreenState{}
}

func (red *RedState) PrintInstructions(userInterface ui.UserInterface) {
	userInterface.Writeln("Describe the failing test to write e.g. 'add 1 and 2 and assert that output equals 3'", ui.Style{Display: ui.BOX})
	userInterface.Write("\nRED > ", ui.Style{Color: ui.RED})
}
