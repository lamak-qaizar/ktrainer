package main

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"os"
	"strings"
)

//go:embed prompts/*.tmpl
var promptFiles embed.FS

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
	tmpl, _ := template.ParseFS(promptFiles, "prompts/test.tmpl")
	var prompt bytes.Buffer
	tmpl.Execute(&prompt, struct{ Instructions string }{Instructions: instructions})

	env := append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN="+assitant.auth.token)
	output, err := assitant.runner.Run("claude", []string{"-p", prompt.String(), "--allowedTools", "Edit", "--safe-mode"}, kataDir, env, io.Discard)

	assitant.invalidatePersistedTokenIfSessionHasExpired(string(output), err)

	return string(output), err
}
