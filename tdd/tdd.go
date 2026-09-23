package tdd

import "testing"

type Phase int

const (
	Red Phase = iota
	Green
	Refactor
)

func (p Phase) String() string {
	switch p {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Refactor:
		return "Refactor"
	}
	return "UNKNOWN"
}

type TDD interface {
	Phase() Phase
	Advance(testsPassed bool) TDD
	Prompt(userInput string) string
	ForceAdvance() TDD
}

type TDDStateMachine struct {
	phase Phase
}

func NewTDDStateMachine() TDD {
	return &RedState{}
}

func (m *TDDStateMachine) Phase() Phase {
	return m.phase
}

func (m *TDDStateMachine) Advance(testsPassed bool) {
	switch m.phase {
	case Red:
		if !testsPassed {
			m.phase = Green
		}
	case Green:
		if testsPassed {
			m.phase = Refactor
		}
	case Refactor:
		m.phase = Refactor
	}
}

func (m *TDDStateMachine) RefactorDone(testsPassed bool) {
	if m.phase == Refactor && testsPassed {
		m.phase = Red
	}
}

func (m *TDDStateMachine) Prompt(userInput string) string {
	switch m.phase {
	case Red:
		return PromptFrom("red.tmpl", userInput)
	case Green:
		return PromptFrom("green.tmpl", userInput)
	case Refactor:
		return PromptFrom("refactor.tmpl", userInput)
	}
	panic("Undefined phase :o")
}

type TestTDDStateMachine struct {
	TDD
	t              *testing.T
	expectedPhases []Phase
	pos            int
}

func NewTestTDDStateMachine(t *testing.T, expected []Phase) *TestTDDStateMachine {
	return &TestTDDStateMachine{
		TDD:            NewTDDStateMachine(),
		t:              t,
		expectedPhases: expected,
	}
}

func (tdd *TestTDDStateMachine) Advance(testsPassed bool) TDD {
	nextState := tdd.TDD.Advance(testsPassed)

	expected := tdd.expectedPhases[tdd.pos]
	if nextState.Phase() != expected {
		tdd.t.Fatalf("[Step %d] Expected TDD phase %s, got: %s", tdd.pos, expected, nextState.Phase())
	}

	return &TestTDDStateMachine{
		TDD:            nextState,
		t:              tdd.t,
		expectedPhases: tdd.expectedPhases,
		pos:            tdd.pos + 1,
	}
}
