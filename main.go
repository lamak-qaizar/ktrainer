package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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
	reader := bufio.NewReader(os.Stdin)
	machine := NewTDDStateMachine()
	assistant, err := NewClaudeCLIAssistant()
	if err != nil {
		fmt.Println("Failed to set up Claude CLI:", err)
		os.Exit(1)
	}

	for {
		fmt.Printf("\n--- Phase: %s ---\n", machine.Phase())
		fmt.Println("Type an instruction for Claude, or 'next' if you edited manually:")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "next" {
			output, err := assistant.ProposeChange(machine.Phase(), input, KATA_DIR)
			if err != nil {
				fmt.Println("Error calling Claude:", output)
				continue
			}
			fmt.Println(output)
		}

		fmt.Println("Running tests...")
		testsPassed := runTests(KATA_DIR)

		fmt.Println("Tests passed:", testsPassed)
		machine.Advance(testsPassed)
	}
}
