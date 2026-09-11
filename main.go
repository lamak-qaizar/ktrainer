package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runTests(dir string) bool {
	cmd := exec.Command("dotnet", "test")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}

func main() {
	if err := ensureOAuthToken(); err != nil {
		fmt.Println("Failed to set up Claude token:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	phase := Red
	kataDir := "mars-rover/MarsRover"
	assistant := ClaudeCLIAssistant{}

	for {
		fmt.Printf("\n--- Phase: %s ---\n", phase)
		fmt.Println("Type an instruction for Claude, or 'next' if you edited manually:")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input != "next" {
			output, err := assistant.ProposeChange(phase, input, kataDir)
			if err != nil {
				if strings.Contains(output, "OAuth session expired") {
					fmt.Println("Your Claude session has expired. Please run 'claude' to log in, then restart this tool.")
					os.Exit(1)
				}
				fmt.Println("Error calling Claude:", output)
				continue
			}
			fmt.Println(output)
		}

		fmt.Println("Running tests...")
		testsPassed := runTests(kataDir)

		fmt.Println("Tests passed:", testsPassed)
		phase = nextPhase(phase, testsPassed)
	}
}
