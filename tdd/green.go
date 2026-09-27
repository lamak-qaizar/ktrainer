package tdd

import (
	"kata-trainer/ui"
)

type GreenState struct {
}

func (green *GreenState) Phase() Phase {
	return Green
}

func (green *GreenState) Advance(testsPassed bool) TDDPhase {
	if !testsPassed {
		return green
	}

	return &RefactorState{}
}

func (green *GreenState) Prompt(userInput string) string {
	return PromptFrom("green.tmpl", userInput)
}

func (green *GreenState) ForceAdvance() TDDPhase {
	return &RefactorState{}
}

func (green *GreenState) PrintInstructions(userInterface ui.UserInterface) {
	instructions(userInterface, green.Phase().String(), ui.GREEN,
		"How would you like to pass the test? Provide the simplest solution possible.")
}
