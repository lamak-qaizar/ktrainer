package main

import (
	"fmt"
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"strings"
)

type AppRunner struct {
	userInterface ui.UserInterface
	assistant     Assistant
	testRunner    TestRunner
	kataDir       string
}

func NewAppRunner(kataDir string, userInterface ui.UserInterface, assistant Assistant, testRunner TestRunner) *AppRunner {
	return &AppRunner{kataDir: kataDir, userInterface: userInterface, assistant: assistant, testRunner: testRunner}
}

func (appRunner *AppRunner) Run(tddPhase tdd.TDDPhase) {
	tddPhase.PrintInstructions(appRunner.userInterface)

	input := appRunner.userInterface.Read()

	if input == "exit" {
		appRunner.userInterface.Writeln("Exiting kata trainer.", ui.Style{})
		return
	}

	if input == "next" {
		appRunner.Run(tddPhase.ForceAdvance())
		return
	}

	spinner := ui.NewSpinner("Clauding...")
	output, err := appRunner.assistant.ProposeChange(tddPhase.Prompt(input), KATA_DIR)
	if err != nil {
		spinner.Fail()
		appRunner.userInterface.Writeln("Error calling Claude:"+output, ui.Style{})
		appRunner.Run(tddPhase)
		return
	}

	if reason, rejected := extractRejection(output); rejected {
		spinner.Fail()
		appRunner.userInterface.Write("A TDD rule was broken, no changes made. "+reason, ui.Style{})
		appRunner.Run(tddPhase)
		return
	}

	spinner.Success()
	appRunner.userInterface.Writeln(output, ui.Style{})

	appRunner.userInterface.Writeln("Running tests...", ui.Style{})
	testsPassed := appRunner.testRunner.Run(KATA_DIR)
	appRunner.userInterface.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed), ui.Style{})

	appRunner.Run(tddPhase.Advance(testsPassed))
}

func extractRejection(output string) (reason string, rejected bool) {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "REJECTED:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "REJECTED:")), true
		}
	}
	return "", false
}
