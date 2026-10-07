package app

import (
	"io"
	"os"
	"strings"
)

type Assistant interface {
	ProposeChange(instruction string, kataDir string) (string, error)
}

type ClaudeCLIAssistant struct {
	config *Config
	runner Runner
}

func NewClaudeCLIAssistant(config *Config, runner Runner) *ClaudeCLIAssistant {
	return &ClaudeCLIAssistant{config: config, runner: runner}
}

func (assistant ClaudeCLIAssistant) invalidatePersistedTokenIfSessionHasExpired(output string, err error) {
	if err != nil {
		if strings.Contains(string(output), "OAuth session expired") ||
			strings.Contains(string(output), "Your organization has disabled Claude subscription") {
			assistant.config.Invalidate()
		}
	}
}

func (assitant ClaudeCLIAssistant) ProposeChange(instructions string, kataDir string) (string, error) {
	env := append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN="+assitant.config.Token)
	args := []string{"-p", instructions, "--allowedTools", "Edit,Create", "--safe-mode", "--permission-mode", "acceptEdits", "--model", assitant.config.Model}
	if supportsEffort(assitant.config.Model) && assitant.config.Effort != "" {
		args = append(args, "--effort", assitant.config.Effort)
	}

	output, err := assitant.runner.Run("claude", args, kataDir, env, io.Discard)

	assitant.invalidatePersistedTokenIfSessionHasExpired(string(output), err)

	return string(output), err
}

func supportsEffort(model string) bool {
	return model != "haiku"
}
