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

func (red *RedState) Advance(testsPassed bool) TDDPhase {
	if testsPassed {
		return red.NoAdvance()
	}

	return &GreenState{Advanced: true}
}

func (red *RedState) Prompt(userInput string) string {
	return PromptFrom("red.tmpl", userInput)
}

func (red *RedState) ForceAdvance() TDDPhase {
	return &GreenState{Advanced: true}
}

func (red *RedState) PrintInstructions(userInterface ui.UserInterface) {
	if red.First {
		instructions(userInterface, red.Phase().String(), ui.RED,
			"Describe the first failing test.\n\nTests are declarative: they describe the 'what'. Usually in the form of GWT (given, when, then).", red.Advanced)
	} else {
		instructions(userInterface, red.Phase().String(), ui.RED,
			"Describe the failing test to write.", red.Advanced)
	}
}

func (red *RedState) NoAdvance() TDDPhase {
	return &RedState{}
}
