package main

import (
	"kata-trainer/tdd"
	"kata-trainer/ui"
)

// used when packaging .exe, see build file
var version = "dev"

func SetupConfig(userInterface ui.UserInterface) *Config {
	config := InitConfig()
	if config.HasValidToken() {
		return config
	}

	userInterface.Writeln("Run 'claude setup-token' in another Terminal and paste the token here.")
	userInterface.Write("> ")
	token := userInterface.Read()

	return config.UpdateToken(token)
}

func main() {
	userInterface := ui.NewUI()
	userInterface.Title("ktrainer")
	userInterface.WriteColoured("v"+version+"\n\n", ui.GRAY)

	commandRunner := NewCommandRunner()
	testRunner := NewDotnetTestRunner(userInterface.Writer(), commandRunner)
	auth := SetupConfig(userInterface)
	assistant := NewClaudeCLIAssistant(auth, commandRunner)

	kata, err := NewKata(ExtractKataDirFromOsArgs(), userInterface.Writer())
	if err != nil {
		return
	}

	NewAppRunner(*kata, userInterface, assistant, testRunner).Run(&tdd.RedState{})
}
