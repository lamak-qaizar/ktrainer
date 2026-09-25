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

func (red *GreenState) PrintInstructions(userInterface ui.UserInterface) {
	userInterface.Writeln("\nGREEN: Type an instruction for Claude", ui.Style{})
	userInterface.Write("> ", ui.Style{})
}
