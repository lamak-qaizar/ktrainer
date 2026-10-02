package tdd

import (
	"kata-trainer/ui"
)

type RefactorState struct {
	Advanced bool
}

func (refactor *RefactorState) Phase() Phase {
	return Refactor
}

func (refactor *RefactorState) Advance(testsPassed bool) TDDPhase {
	return refactor.NoAdvance()
}

func (refactor *RefactorState) Prompt(userInput string) string {
	return PromptFrom("refactor.tmpl", userInput)
}

func (refacor *RefactorState) ForceAdvance() TDDPhase {
	return &RedState{Advanced: true}
}

func (refactor *RefactorState) PrintInstructions(userInterface ui.UserInterface) {
	instructions(userInterface, refactor.Phase().String(), ui.YELLOW,
		"Specify design improvements, or type 'next'.", refactor.Advanced)
}

func (red *RefactorState) NoAdvance() TDDPhase {
	return &RefactorState{}
}
