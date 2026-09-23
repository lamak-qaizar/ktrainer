package tdd

type GreenState struct {
}

func (green *GreenState) Phase() Phase {
	return Green
}

func (green *GreenState) Advance(testsPassed bool) TDD {
	if !testsPassed {
		return green
	}

	return &RefactorState{}
}

func (green *GreenState) Prompt(userInput string) string {
	return PromptFrom("green.tmpl", userInput)
}

func (green *GreenState) ForceAdvance() TDD {
	return &RefactorState{}
}
