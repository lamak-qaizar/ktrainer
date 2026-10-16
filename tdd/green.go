package tdd

import (
	"kata-trainer/ui"
)

type GreenState struct {
	First    bool
	Advanced bool
}

func (green *GreenState) Phase() Phase {
	return Green
}

func (green *GreenState) Advance(testsPassed bool) TDDPhase {
	if !testsPassed {
		return green.NoAdvance()
	}

	return &RefactorState{Advanced: true, First: green.First}
}

func (green *GreenState) Prompt(userInput string) string {
	return PromptFrom("green.tmpl", userInput)
}

func (green *GreenState) ForceAdvance() TDDPhase {
	return &RefactorState{Advanced: true, First: green.First}
}

func (green *GreenState) PrintInstructions(userInterface ui.UserInterface) {
	if green.First {
		instructions(userInterface, green.Phase().String(), ui.GREEN,
			"Specify the simplest solution possible to pass the test.\n\nIf the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.", green.Advanced)
	} else {
		instructions(userInterface, green.Phase().String(), ui.GREEN,
			"How would you like to pass the test?", green.Advanced)
	}
}

func (green *GreenState) NoAdvance() TDDPhase {
	return &GreenState{First: green.First}
}
