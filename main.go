package main

import (
	"fmt"
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"strings"
	"time"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runApp(userInterface ui.UserInterface, assistant Assistant, testRunner TestRunner, tddPhase tdd.TDDPhase) {
	tddPhase.PrintInstructions(userInterface)

	input := userInterface.Read()

	if input == "exit" {
		userInterface.Writeln("Exiting kata trainer.", ui.Style{})
		return
	}

	if input == "next" {
		runApp(userInterface, assistant, testRunner, tddPhase.ForceAdvance())
		return
	}

	output, err := assistant.ProposeChange(tddPhase.Prompt(input), KATA_DIR)
	if err != nil {
		userInterface.Writeln("Error calling Claude:"+output, ui.Style{})
		runApp(userInterface, assistant, testRunner, tddPhase)
		return
	}

	if reason, rejected := extractRejection(output); rejected {
		userInterface.Write("A TDD rule was broken, no changes made. "+reason, ui.Style{})
		runApp(userInterface, assistant, testRunner, tddPhase)
		return
	}

	userInterface.Writeln(output, ui.Style{})

	userInterface.Writeln("Running tests...", ui.Style{})
	testsPassed := testRunner.Run(KATA_DIR)
	userInterface.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed), ui.Style{})

	runApp(userInterface, assistant, testRunner, tddPhase.Advance(testsPassed))
}

func extractRejection(output string) (reason string, rejected bool) {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "REJECTED:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "REJECTED:")), true
		}
	}
	return "", false
}

func SetupClaudeAuth(userInterface ui.UserInterface) *ClaudeAuth {
	token := LoadTokenFromConfig()
	if token != "" {
		return NewClaudeAuth(token)
	}

	userInterface.Writeln("Run 'claude setup-token' and paste the token here.", ui.Style{})
	userInterface.Write("> ", ui.Style{})
	token = userInterface.Read()
	SaveTokenToConfig(token, time.Now().Add(tokenTTL))
	return NewClaudeAuth(token)
}

func main() {
	userInterface := ui.NewUI()
	commandRunner := NewCommandRunner()
	testRunner := NewDotnetTestRunner(userInterface.Writer(), commandRunner)
	auth := SetupClaudeAuth(userInterface)
	assistant := NewClaudeCLIAssistant(auth, commandRunner)

	userInterface.Title("ktrainer")
	runApp(userInterface, assistant, testRunner, &tdd.RedState{})
}
