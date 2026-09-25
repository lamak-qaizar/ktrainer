package tdd

import (
	"kata-trainer/ui"
)

type RedState struct {
}

func (red *RedState) Phase() Phase {
	return Red
}

func (red *RedState) Advance(testsPassed bool) TDD {
	if testsPassed {
		return red
	}

	return &GreenState{}
}

func (red *RedState) Prompt(userInput string) string {
	return PromptFrom("red.tmpl", userInput)
}

func (red *RedState) ForceAdvance() TDD {
	return &GreenState{}
}

func (red *RedState) PrintInstructions(userInterface ui.UserInterface) {
	userInterface.Writeln("\nRED: Type an instruction for Claude", ui.Style{})
	userInterface.Write("> ", ui.Style{})
}
