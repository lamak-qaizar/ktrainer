package main

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
