package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	phase := Red

	for {
		fmt.Printf("\n--- Phase: %s ---\n", phase)
		fmt.Printf("Type an instruction for Claude, or 'next' if you edited manually:")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		fmt.Println("You typed:", input)

		testsPassed := true

		phase = nextPhase(phase, testsPassed)
	}
}
