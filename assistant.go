package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var oAuthToken string

func configPath() string {
	return filepath.Join(os.Getenv("APPDATA"), "kata-trainer", "config.json")
}

func loadToken() string {
	data, err := os.ReadFile(configPath())
	if err != nil {
		return ""
	}
	var cfg struct {
		Token string `json:"token"`
	}
	if json.Unmarshal(data, &cfg) != nil {
		return ""
	}
	return cfg.Token
}

func saveToken(token string) error {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{token})
	return os.WriteFile(path, data, 0600)
}

func ensureOAuthToken() error {
	if token := loadToken(); token != "" {
		oAuthToken = token
		return nil
	}

	fmt.Println("No Claude token found. Running 'claude setup-token'...")
	cmd := exec.Command("claude", "setup-token")
	cmd.Stdin = os.Stdin

	var buf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &buf)
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	re := regexp.MustCompile(`sk-ant-oat01-[A-Za-z0-9_-]+`)
	match := re.FindString(buf.String())
	if match == "" {
		return fmt.Errorf("Could not find token in 'setup-token' output")
	}

	oAuthToken = match
	return saveToken(match)
}

type ClaudeCLIAssistant struct{}

func (a ClaudeCLIAssistant) ProposeChange(phase Phase, instruction string, kataDir string) (string, error) {
	prompt := fmt.Sprintf(
		"You are helping with the %s phase of TDD. Make ONLY the following small, specific change, nothing else: %s. Do not run any commands, do not run tests, do not verify your change, just make the edit and stop.",
		phase.String(), instruction,
	)

	cmd := exec.Command("claude", "-p", prompt, "--allowedTools", "Edit", "--safe-mode")
	cmd.Dir = kataDir
	cmd.Env = append(os.Environ(), "CLAUDE_CODE_OAUTH_TOKEN="+oAuthToken)

	output, err := cmd.CombinedOutput()
	return string(output), err
}
