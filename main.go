package main

import (
	"kata-trainer/tdd"
	"kata-trainer/ui"
	"os"
	"path/filepath"
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

func kataDir() string {
	if len(os.Args[1:]) > 0 {
		return os.Args[1:][0]
	}
	return "."
}

func main() {
	userInterface := ui.NewUI()
	commandRunner := NewCommandRunner()
	testRunner := NewDotnetTestRunner(userInterface.Writer(), commandRunner)
	auth := SetupClaudeAuth(userInterface)
	assistant := NewClaudeCLIAssistant(auth, commandRunner)

	userInterface.Title("ktrainer")

	matches, _ := filepath.Glob(kataDir() + "/*.csproj")
	if len(matches) == 0 {
		userInterface.Writeln("Run ktrainer inside a C# project (no *.csproj file found). Exiting...", ui.Style{})
		userInterface.Writeln("To create a new project, run:\n\ndotnet new xunit -n <project_name>\n", ui.Style{})
		return
	}

	NewAppRunner(kataDir(), userInterface, assistant, testRunner).Run(&tdd.RedState{})
}
