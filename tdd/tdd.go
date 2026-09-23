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

type TestTDDStateMachine struct {
	TDD
	t              *testing.T
	expectedPhases []Phase
	pos            int
}

func NewTestTDDStateMachine(t *testing.T, expected []Phase) *TestTDDStateMachine {
	return &TestTDDStateMachine{
		TDD:            &RedState{},
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
