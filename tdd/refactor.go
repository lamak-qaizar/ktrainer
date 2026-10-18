package tdd

import (
	"kata-trainer/ui"
)

type RefactorState struct {
	First    bool
	Advanced bool
}

func (refactor *RefactorState) Phase() Phase {
	return Refactor
}

func (refactor *RefactorState) Advance(event AdvanceEvent) TDDPhase {
	if event == Stay {
		return &RefactorState{}
	}

	return &RedState{Advanced: true}
}

func (refactor *RefactorState) Prompt(userInput string) string {
	return PromptFrom("refactor.tmpl", userInput)
}

func (refacor *RefactorState) ForceAdvance() TDDPhase {
	return &RedState{Advanced: true}
}

func (refactor *RefactorState) PrintInstructions(userInterface ui.UserInterface) {
	if refactor.First {
		instructions(userInterface, refactor.Phase().String(), ui.YELLOW,
			"Remove duplication. Clean up code smells in tests and implementation.\n\nHowever, there isn't anything to improve right now. Type 'next'.", refactor.Advanced)

	} else {
		instructions(userInterface, refactor.Phase().String(), ui.YELLOW,
			"Specify design improvements, or type 'next'.", refactor.Advanced)
	}
}
