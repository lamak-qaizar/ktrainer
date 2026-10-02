package tdd

import (
	"kata-trainer/ui"
)

type GreenState struct {
	Advanced bool
}

func (green *GreenState) Phase() Phase {
	return Green
}

func (green *GreenState) Advance(testsPassed bool) TDDPhase {
	if !testsPassed {
		return green.NoAdvance()
	}

	return &RefactorState{Advanced: true}
}

func (green *GreenState) Prompt(userInput string) string {
	return PromptFrom("green.tmpl", userInput)
}

func (green *GreenState) ForceAdvance() TDDPhase {
	return &RefactorState{Advanced: true}
}

func (green *GreenState) PrintInstructions(userInterface ui.UserInterface) {
	instructions(userInterface, green.Phase().String(), ui.GREEN,
		"How would you like to pass the test? Provide the simplest solution possible.", green.Advanced)
}

func (red *GreenState) NoAdvance() TDDPhase {
	return &GreenState{}
}
