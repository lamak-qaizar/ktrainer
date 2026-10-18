package tdd

import (
	"kata-trainer/ui"
)

type RedState struct {
	First    bool
	Advanced bool
}

func (red *RedState) Phase() Phase {
	return Red
}

func (red *RedState) Advance(event AdvanceEvent) TDDPhase {
	if event == TestsPassed || event == Stay {
		return &RedState{First: red.First}
	}

	return &GreenState{Advanced: true, First: red.First}
}

func (red *RedState) Prompt(userInput string) string {
	return PromptFrom("red.tmpl", userInput)
}

func (red *RedState) PrintInstructions(userInterface ui.UserInterface) {
	if red.First {
		instructions(userInterface, red.Phase().String(), ui.RED,
			"Describe the first failing test. Tests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).", red.Advanced)
	} else {
		instructions(userInterface, red.Phase().String(), ui.RED,
			"Describe the failing test to write.", red.Advanced)
	}
}
