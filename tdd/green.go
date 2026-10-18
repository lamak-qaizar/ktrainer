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

func (green *GreenState) Advance(event AdvanceEvent) TDDPhase {
	if event == TestsFailed || event == Stay {
		return &GreenState{First: green.First}
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
			"Congratulations! We have our first failing test.\n\nSpecify the simplest solution possible to pass the test: if the test can be passed by returning a constant value, do that. To generalise the solution, write more tests.", green.Advanced)
	} else {
		instructions(userInterface, green.Phase().String(), ui.GREEN,
			"How would you like to pass the test?", green.Advanced)
	}
}
