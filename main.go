package main

import (
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"time"
)

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
	RunApp("MarsRover", userInterface, assistant, testRunner, &tdd.RedState{})
}
