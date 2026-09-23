package tdd

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
