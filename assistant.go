package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Assistant interface {
	ProposeChange(phase Phase, instruction string, kataDir string) (string, error)
}

type ClaudeCLIAssistant struct {
	auth          *ClaudeAuth
	commandRunner CommandRunner
}

func NewClaudeCLIAssistant(auth *ClaudeAuth) *ClaudeCLIAssistant {
	return &ClaudeCLIAssistant{auth: auth, commandRunner: *NewCommandRunner()}
}

func (assistant ClaudeCLIAssistant) invalidatePersistedTokenIfSessionHasExpired(output string, err error) {
	if err != nil {
		if strings.Contains(string(output), "OAuth session expired") {
			assistant.auth.Invalidate()
		}
	}
}

func (assitant ClaudeCLIAssistant) ProposeChange(phase Phase, instruction string, kataDir string) (string, error) {
	prompt := fmt.Sprintf(
		"You are helping with the %s phase of TDD. Make ONLY the following small, specific change, nothing else: %s. Do not run any commands, do not run tests, do not verify your change, just make the edit and stop.",
		phase.String(), instruction,
	)

	env := append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN="+assitant.auth.token)
	output, err := assitant.commandRunner.Run("claude", []string{"-p", prompt, "--allowedTools", "Edit", "--safe-mode"}, kataDir, env, io.Discard)

	assitant.invalidatePersistedTokenIfSessionHasExpired(string(output), err)

	return string(output), err
}
