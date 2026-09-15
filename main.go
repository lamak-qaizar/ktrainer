package main

import (
	"fmt"
	"os"
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

func main() {
	ui := NewCommandLine(os.Stdin, os.Stdout)
	testRunner := NewDotnetTestRunner(ui.Writer())
	assistant, err := NewClaudeCLIAssistant()
	if err != nil {
		fmt.Println("Failed to set up Claude CLI:", err)
		os.Exit(1)
	}

	runApp(ui, assistant, testRunner)
}
