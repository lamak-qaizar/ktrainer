package main

import (
	"fmt"
	"os"
	"os/exec"
)

const KATA_DIR string = "mars-rover/MarsRover"

func runTests(dir string) bool {
	cmd := exec.Command("dotnet", "test")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}

func main() {
	machine := NewTDDStateMachine()
	ui := NewUI(os.Stdin, os.Stdout)
	assistant, err := NewClaudeCLIAssistant()
	if err != nil {
		fmt.Println("Failed to set up Claude CLI:", err)
		os.Exit(1)
	}

	for {
		ui.Writeln("\n--- Phase: %s ---", machine.Phase())
		ui.Writeln("Type an instruction for Claude")
		ui.Write("> ")

		input := ui.Read()

		if input == "exit" {
			fmt.Println("Exiting kata trainer")
			break
		}

		output, err := assistant.ProposeChange(machine.Phase(), input, KATA_DIR)
		if err != nil {
			fmt.Println("Error calling Claude:", output)
			continue
		}
		fmt.Println(output)

		fmt.Println("Running tests...")
		testsPassed := runTests(KATA_DIR)

		fmt.Println("Tests passed:", testsPassed)
		machine.Advance(testsPassed)
	}
}
