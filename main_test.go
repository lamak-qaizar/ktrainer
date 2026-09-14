package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunApp_QuitExitsCleanly(t *testing.T) {
	input := strings.NewReader("exit\n")
	var output bytes.Buffer

	runApp(NewCommandLine(input, &output), fakeAssistant{}, fakeRunner{})

	if !strings.Contains(output.String(), "Exiting kata trainer.") {
		t.Errorf("Expected to exit.")
	}
}

type fakeAssistant struct{}

func (f fakeAssistant) ProposeChange(Phase, string, string) (string, error) { return "", nil }

type fakeRunner struct{}

func (f fakeRunner) Run(string) bool { return true }
