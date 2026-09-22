package main

import (
	"fmt"
	"time"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runApp(ui UserInterface, assistant Assistant, testRunner TestRunner, tdd TDD) {
	for {
		ui.Writeln(fmt.Sprintf("\n %s: Type an instruction for Claude", tdd.Phase()))
		ui.Write("> ")

		input := ui.Read()

		if input == "exit" {
			ui.Writeln("Exiting kata trainer.")
			break
		}

		output, err := assistant.ProposeChange(tdd.Phase(), input, KATA_DIR)
		if err != nil {
			ui.Writeln("Error calling Claude:" + output)
			continue
		}
		ui.Writeln(output)

		ui.Writeln("Running tests...")
		testsPassed := testRunner.Run(KATA_DIR)

		ui.Writeln(fmt.Sprintf("Tests passed: %t", testsPassed))
		tdd.Advance(testsPassed)
	}
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
	runApp(ui, assistant, testRunner, NewTDDStateMachine())
}
