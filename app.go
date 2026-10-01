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
	kata          Kata
}

func NewAppRunner(kata Kata, userInterface ui.UserInterface, assistant Assistant, testRunner TestRunner) *AppRunner {
	return &AppRunner{kata: kata, userInterface: userInterface, assistant: assistant, testRunner: testRunner}
}

func (appRunner *AppRunner) Run(tddPhase tdd.TDDPhase) {
	tddPhase.PrintInstructions(appRunner.userInterface)

	input := appRunner.userInterface.Read()

	if input == "exit" {
		appRunner.userInterface.Writeln("Exiting kata trainer.")
		return
	}

	if input == "next" {
		appRunner.Run(tddPhase.ForceAdvance())
		return
	}

	appRunner.userInterface.Writeln("")
	spinner := ui.NewSpinner("Clauding...")
	output, err := appRunner.assistant.ProposeChange(tddPhase.Prompt(input), appRunner.kata.Dir)
	if err != nil {
		spinner.Fail()
		appRunner.userInterface.Writeln("Error calling Claude:" + output)
		appRunner.Run(tddPhase)
		return
	}

	if reason, rejected := extractRejection(output); rejected {
		spinner.Fail()
		appRunner.userInterface.Write("A TDD rule was broken, no changes made. " + reason)
		appRunner.Run(tddPhase)
		return
	}

	spinner.Success()
	appRunner.userInterface.Writeln(output)

	appRunner.userInterface.Writeln("Running tests...")
	testsPassed := appRunner.testRunner.Run(appRunner.kata.Dir)
	appRunner.userInterface.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed))

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
