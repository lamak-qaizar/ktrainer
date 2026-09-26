package tdd

import (
	"kata-trainer/ui"
)

type RefactorState struct {
}

func (refactor *RefactorState) Phase() Phase {
	return Refactor
}

func (refactor *RefactorState) Advance(testsPassed bool) TDDPhase {
	return refactor
}

func (refactor *RefactorState) Prompt(userInput string) string {
	return PromptFrom("refactor.tmpl", userInput)
}

func (refacor *RefactorState) ForceAdvance() TDDPhase {
	return &RedState{}
}

func (red *RefactorState) PrintInstructions(userInterface ui.UserInterface) {
	userInterface.Writeln("Specify any design improvements.", ui.Style{Display: ui.BOX})
	userInterface.Write("\nREFACTOR > ", ui.Style{Color: ui.YELLOW})
}
