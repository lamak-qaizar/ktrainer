package main

import (
	"time"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runApp(ui UI, assistant Assistant, testRunner TestRunner) {
	machine := NewTDDStateMachine()

	for {
		ui.Writeln("\n--- Phase: %s ---", machine.Phase())
		ui.Writeln("Type an instruction for Claude")
		ui.Write("> ")

		input := ui.Read()

		if input == "exit" {
			ui.Writeln("Exiting kata trainer.")
			break
		}

		output, err := assistant.ProposeChange(machine.Phase(), input, KATA_DIR)
		if err != nil {
			ui.Writeln("Error calling Claude:", output)
			continue
		}
		ui.Writeln(output)

		ui.Writeln("Running tests...")
		testsPassed := testRunner.Run(KATA_DIR)

		ui.Writeln("Tests passed:", testsPassed)
		machine.Advance(testsPassed)
	}
}

func SetupClaudeAuth(ui UI) *ClaudeAuth {
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
	ui := NewStdIOCommandLine()
	commandRunner := NewCommandRunner()
	testRunner := NewDotnetTestRunner(ui.Writer(), commandRunner)
	auth := SetupClaudeAuth(ui)
	assistant := NewClaudeCLIAssistant(auth, commandRunner)
	runApp(ui, assistant, testRunner)
}
