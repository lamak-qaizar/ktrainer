package tdd

import (
	"kata-trainer/ui"
	"testing"
)

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

type TDDPhase interface {
	Phase() Phase
	Advance(testsPassed bool) TDDPhase
	Prompt(userInput string) string
	ForceAdvance() TDDPhase
	PrintInstructions(userInterface ui.UserInterface)
}

type TestTDDPhase struct {
	TDDPhase
	t              *testing.T
	expectedPhases []Phase
	pos            int
}

func NewTestTDDPhase(t *testing.T, expected []Phase) *TestTDDPhase {
	return &TestTDDPhase{
		TDDPhase:       &RedState{},
		t:              t,
		expectedPhases: expected,
	}
}

func (tdd *TestTDDPhase) Advance(testsPassed bool) TDDPhase {
	nextState := tdd.TDDPhase.Advance(testsPassed)

	expected := tdd.expectedPhases[tdd.pos]
	if nextState.Phase() != expected {
		tdd.t.Fatalf("[Step %d] Expected TDD phase %s, got: %s", tdd.pos, expected, nextState.Phase())
	}

	return &TestTDDPhase{
		TDDPhase:       nextState,
		t:              tdd.t,
		expectedPhases: tdd.expectedPhases,
		pos:            tdd.pos + 1,
	}
}
