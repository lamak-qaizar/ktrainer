package main

import (
	"fmt"
	"os"
	"os/exec"
)

type ClaudeCLIAssistant struct{}

func (a ClaudeCLIAssistant) ProposeChange(phase Phase, instruction string, kataDir string) (string, error) {
	prompt := fmt.Sprintf(
		"You are helping with the %s phase of TDD. Makr ONLY the following small, specific change, nothing else: %s",
		phase.String(), instruction,
	)

	cmd := exec.Command("claude", "-p", prompt, "--allowedTools", "Edit")
	cmd.Dir = kataDir
	cmd.Env = append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN=...")

	output, err := cmd.CombinedOutput()
	return string(output), err
}
