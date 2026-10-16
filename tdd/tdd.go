package tdd

import (
	"kata-trainer/ui"
	"strings"
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
	NoAdvance() TDDPhase
}

type TestTDDPhase struct {
	TDDPhase
	t                   *testing.T
	expectedTransitions []Transition
	pos                 int
}

type Transition struct {
	Current Phase
	Next    Phase
}

func NewTestTDDPhase(t *testing.T, expected []Transition) *TestTDDPhase {
	return &TestTDDPhase{
		TDDPhase:            &RedState{Advanced: true, First: true},
		t:                   t,
		expectedTransitions: expected,
		pos:                 0,
	}
}

func (tdd *TestTDDPhase) Advance(testsPassed bool) TDDPhase {
	nextState := tdd.TDDPhase.Advance(testsPassed)

	expectedTransition := tdd.expectedTransitions[tdd.pos]

	if tdd.TDDPhase.Phase() != expectedTransition.Current {
		tdd.t.Fatalf("[Step %d] Current TDD phase %s, got: %s", tdd.pos, expectedTransition.Current, tdd.TDDPhase.Phase())
	}

	if nextState.Phase() != expectedTransition.Next {
		tdd.t.Fatalf("[Step %d] Expected TDD phase %s, got: %s", tdd.pos, expectedTransition.Next, nextState.Phase())
	}

	return &TestTDDPhase{
		TDDPhase:            nextState,
		t:                   tdd.t,
		expectedTransitions: tdd.expectedTransitions,
		pos:                 tdd.pos + 1,
	}
}

func instructions(userInterface ui.UserInterface, phase string, color ui.Color, senseiSays string, advanced bool) {
	if advanced {
		userInterface.SenseiSays(senseiSays)
	}
	userInterface.WriteColoured("\n"+strings.ToUpper(phase)+" > ", color)
}
