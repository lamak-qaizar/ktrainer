package main

import (
	"io"
	"os"
	"strings"
)

type Assistant interface {
	ProposeChange(phase Phase, instruction string, kataDir string) (string, error)
}

type ClaudeCLIAssistant struct {
	auth   *ClaudeAuth
	runner Runner
}

func NewClaudeCLIAssistant(auth *ClaudeAuth, runner Runner) *ClaudeCLIAssistant {
	return &ClaudeCLIAssistant{auth: auth, runner: runner}
}

func (assistant ClaudeCLIAssistant) invalidatePersistedTokenIfSessionHasExpired(output string, err error) {
	if err != nil {
		if strings.Contains(string(output), "OAuth session expired") {
			assistant.auth.Invalidate()
		}
	}
}

func (assitant ClaudeCLIAssistant) ProposeChange(phase Phase, instructions string, kataDir string) (string, error) {
	env := append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN="+assitant.auth.token)
	output, err := assitant.runner.Run("claude",
		[]string{"-p", promptFrom("red.tmpl", instructions), "--allowedTools", "Edit", "--safe-mode"},
		kataDir, env, io.Discard)

	assitant.invalidatePersistedTokenIfSessionHasExpired(string(output), err)

	return string(output), err
}
