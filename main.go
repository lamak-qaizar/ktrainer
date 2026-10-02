package main

import (
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"time"
)

// used when packaging .exe, see build file
var version = "dev"

func SetupClaudeAuth(userInterface ui.UserInterface) *ClaudeAuth {
	token := LoadTokenFromConfig()
	if token != "" {
		return NewClaudeAuth(token)
	}

	userInterface.Writeln("Run 'claude setup-token' in another Terminal and paste the token here.")
	userInterface.Write("> ")
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
	userInterface.WriteColoured("v"+version+"\n\n", ui.GRAY)

	kata, err := NewKata(ExtractKataDirFromOsArgs(), userInterface.Writer())
	if err != nil {
		return
	}

	NewAppRunner(*kata, userInterface, assistant, testRunner).Run(&tdd.RedState{})
}
