package main

import (
	"fmt"
	"kata-trainer/tdd"
	"strings"
	"time"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runApp(ui UserInterface, assistant Assistant, testRunner TestRunner, tddStateMachine tdd.TDD) {
	ui.Writeln(fmt.Sprintf("\n %s: Type an instruction for Claude", tddStateMachine.Phase()), Style{})
	ui.Write("> ", Style{})

	input := ui.Read()

	if input == "exit" {
		ui.Writeln("Exiting kata trainer.", Style{})
		return
	}

	if input == "next" {
		runApp(ui, assistant, testRunner, tddStateMachine.ForceAdvance())
		return
	}

	output, err := assistant.ProposeChange(tddStateMachine.Prompt(input), KATA_DIR)
	if err != nil {
		ui.Writeln("Error calling Claude:"+output, Style{})
		runApp(ui, assistant, testRunner, tddStateMachine)
		return
	}

	if reason, rejected := extractRejection(output); rejected {
		ui.Write("A TDD rule was broken, no changes made. "+reason, Style{})
		runApp(ui, assistant, testRunner, tddStateMachine)
		return
	}

	ui.Writeln(output, Style{})

	ui.Writeln("Running tests...", Style{})
	testsPassed := testRunner.Run(KATA_DIR)
	ui.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed), Style{})

	runApp(ui, assistant, testRunner, tddStateMachine.Advance(testsPassed))
}

func extractRejection(output string) (reason string, rejected bool) {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "REJECTED:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "REJECTED:")), true
		}
	}
	return "", false
}

func SetupClaudeAuth(ui UserInterface) *ClaudeAuth {
	token := LoadTokenFromConfig()
	if token != "" {
		return NewClaudeAuth(token)
	}

	ui.Writeln("Run 'claude setup-token' and paste the token here.", Style{})
	ui.Write("> ", Style{})
	token = ui.Read()
	SaveTokenToConfig(token, time.Now().Add(tokenTTL))
	return NewClaudeAuth(token)
}

func main() {
	ui := NewUI()
	commandRunner := NewCommandRunner()
	testRunner := NewDotnetTestRunner(ui.Writer(), commandRunner)
	auth := SetupClaudeAuth(ui)
	assistant := NewClaudeCLIAssistant(auth, commandRunner)

	ui.Writeln("ktrainer", Style{title: true})
	runApp(ui, assistant, testRunner, &tdd.RedState{})
}
