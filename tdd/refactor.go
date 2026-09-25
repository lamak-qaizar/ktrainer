package tdd

import (
	"kata-trainer/ui"
)

type RefactorState struct {
}

func (refactor *RefactorState) Phase() Phase {
	return Refactor
}

func (refactor *RefactorState) Advance(testsPassed bool) TDD {
	return refactor
}

func (refactor *RefactorState) Prompt(userInput string) string {
	return PromptFrom("refactor.tmpl", userInput)
}

func (refacor *RefactorState) ForceAdvance() TDD {
	return &RedState{}
}

func (red *RefactorState) PrintInstructions(userInterface ui.UserInterface) {
	userInterface.Writeln("\nREFACTOR: Type an instruction for Claude", ui.Style{})
	userInterface.Write("> ", ui.Style{})
}
