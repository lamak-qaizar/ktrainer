package main

import (
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

	spinner := appRunner.userInterface.Spinner("Clauding...")
	output, err := appRunner.assistant.ProposeChange(tddPhase.Prompt(input), appRunner.kata.Dir)
	if err != nil {
		spinner.Fail()
		appRunner.userInterface.Writeln("\nError calling Claude:" + output)
		appRunner.Run(tddPhase.NoAdvance())
		return
	}

	if reason, rejected := extractRejection(output); rejected {
		spinner.Fail()
		appRunner.userInterface.WriteColoured("\nNo changes made: ", ui.RED)
		appRunner.userInterface.Writeln(reason)
		appRunner.Run(tddPhase.NoAdvance())
		return
	}

	spinner.Success()
	appRunner.userInterface.Writeln("\n" + output)

	testsPassed := appRunner.runTests()

	appRunner.Run(tddPhase.Advance(testsPassed))
}

func (appRunner *AppRunner) runTests() (testsPassed bool) {
	spinner := appRunner.userInterface.Spinner("Running tests...")
	testsPassed, testOutput := appRunner.testRunner.Run(appRunner.kata.Dir)
	if !testsPassed {
		spinner.Fail()
		appRunner.userInterface.Writeln(testOutput)
	} else {
		spinner.Success()
	}
	return testsPassed
}

func extractRejection(output string) (reason string, rejected bool) {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "REJECTED:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "REJECTED:")), true
		}
	}
	return "", false
}
