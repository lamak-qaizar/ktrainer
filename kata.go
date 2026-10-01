package main

import (
	"errors"
	"fmt"
	"io"
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

func NewKata(dir string, output io.Writer) (*Kata, error) {
	matches, _ := filepath.Glob(dir + "/*.csproj")
	if len(matches) == 0 {
		fmt.Fprint(output, "Oops, no *.csproj file found. Run ktrainer inside a C# project or specify path: 'krainer <path>'.\n\n")
		fmt.Fprint(output, "To create a new project: dotnet new xunit -n <name>\n\n")
		fmt.Fprint(output, "Exiting...\n\n")
		return nil, errors.New("Run ktrainer inside a C# project")
	}

	return &Kata{Dir: dir}, nil
}
