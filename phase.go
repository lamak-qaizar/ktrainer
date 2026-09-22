package main

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
	Advance(testsPassed bool)
}

type TDDStateMachine struct {
	phase Phase
}

func NewTDDStateMachine() *TDDStateMachine {
	return &TDDStateMachine{phase: Red}
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

type TestTDDStateMachine struct {
	*TDDStateMachine
	t              *testing.T
	expectedPhases []Phase
	pos            int
}

func NewTestTDDStateMachine(t *testing.T, expected []Phase) *TestTDDStateMachine {
	return &TestTDDStateMachine{
		TDDStateMachine: NewTDDStateMachine(),
		t:               t,
		expectedPhases:  expected,
	}
}

func (tdd *TestTDDStateMachine) Advance(testsPassed bool) {
	tdd.TDDStateMachine.Advance(testsPassed)

	expected := tdd.expectedPhases[tdd.pos]
	if tdd.Phase() != expected {
		tdd.t.Fatalf("Expected TDD phase %s, got: %s", expected, tdd.Phase())
	}
}
