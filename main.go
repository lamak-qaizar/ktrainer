package main

import (
	"kata-trainer/app"
	"kata-trainer/tdd"
	"kata-trainer/ui"
)

// used when packaging .exe, see build file
var version = "dev"

func SetupConfig(userInterface ui.UserInterface) *app.Config {
	config := app.InitConfig()
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

	commandRunner := app.NewCommandRunner()
	testRunner := app.NewDotnetTestRunner(userInterface.Writer(), commandRunner)
	config := SetupConfig(userInterface)
	userInterface.WriteColoured("v"+version+" | $env:APPDATA\\ktrainer\\config.json { model: "+config.Model+" }\n\n", ui.GRAY)
	userInterface.Writeln("NOTE: Memory is not retained between prompts, please be explicit in your instructions.\n")

	assistant := app.NewClaudeCLIAssistant(config, commandRunner)

	kata, err := app.NewKata(app.ExtractKataDirFromOsArgs(), userInterface)
	if err != nil {
		return
	}

	app.NewAppRunner(*kata, userInterface, assistant, testRunner).Run(&tdd.RedState{Advanced: true})
}
