package main

import (
	"fmt"
	"kata-trainer/tdd"
	"time"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runApp(ui UserInterface, assistant Assistant, testRunner TestRunner, tddStateMachine tdd.TDD) {
	ui.Writeln(fmt.Sprintf("\n %s: Type an instruction for Claude", tddStateMachine.Phase()))
	ui.Write("> ")

	input := ui.Read()

	if input == "exit" {
		ui.Writeln("Exiting kata trainer.")
		return
	}

	if input == "next" {
		runApp(ui, assistant, testRunner, tddStateMachine.ForceAdvance())
		return
	}

	output, err := assistant.ProposeChange(tddStateMachine.Prompt(input), KATA_DIR)
	if err != nil {
		ui.Writeln("Error calling Claude:" + output)
		runApp(ui, assistant, testRunner, tddStateMachine)
		return
	}
	ui.Writeln(output)

	ui.Writeln("Running tests...")
	testsPassed := testRunner.Run(KATA_DIR)
	ui.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed))

	runApp(ui, assistant, testRunner, tddStateMachine.Advance(testsPassed))
}

func SetupClaudeAuth(ui UserInterface) *ClaudeAuth {
	token := LoadTokenFromConfig()
	if token != "" {
		return NewClaudeAuth(token)
	}

	ui.Writeln("Run 'claude setup-token' and paste the token here.")
	ui.Write("> ")
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
	runApp(ui, assistant, testRunner, &tdd.RedState{})
}
