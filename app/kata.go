package app

import (
	"errors"
	"kata-trainer/ui"
	"os"
	"path/filepath"
)

type Kata struct {
	Dir string
}

func ExtractKataDirFromOsArgs() string {
	if len(os.Args[1:]) > 0 {
		return os.Args[1:][0]
	}
	return "."
}

func NewKata(dir string, userInterface ui.UserInterface) (*Kata, error) {
	matches, _ := filepath.Glob(dir + "/*.csproj")
	if len(matches) == 0 {
		userInterface.WriteColoured("Oops, no *.csproj file found. Run ktrainer inside a C# project or specify path: 'krainer <path>'.\n\n", ui.RED)
		userInterface.Writeln("To create a new project: dotnet new xunit -n <name>\n")
		userInterface.Writeln("Exiting...\n")
		return nil, errors.New("Run ktrainer inside a C# project")
	}

	return &Kata{Dir: dir}, nil
}
