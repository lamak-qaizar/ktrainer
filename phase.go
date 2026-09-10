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

func nextPhase(current Phase, testsPassed bool) Phase {
	switch current {
	case Red:
		if !testsPassed {
			return Green
		}
		return Red
	case Green:
		if testsPassed {
			return Refactor
		}
	case Refactor:
		return Refactor
	}
	return current
}
