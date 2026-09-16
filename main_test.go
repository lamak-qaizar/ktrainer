package main

import (
	"strings"
	"testing"
)

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	input := strings.NewReader("exit\n")
	ui := NewTestUI(input)

	runApp(ui, fakeAssistant{}, fakeRunner{})

	ui.AssertOutputContains(t, "Exiting kata trainer.")
}

type fakeAssistant struct{}

func (f fakeAssistant) ProposeChange(Phase, string, string) (string, error) { return "", nil }

type fakeRunner struct{}

func (f fakeRunner) Run(string) bool { return true }
