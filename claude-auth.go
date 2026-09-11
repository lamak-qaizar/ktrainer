package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type ClaudeAuth struct {
	token string
}

func NewClaudeAuth() *ClaudeAuth {
	return &ClaudeAuth{}
}

func (auth *ClaudeAuth) configPath() string {
	return filepath.Join(os.Getenv("APPDATA"), "kata-trainer", "config.json")
}

func (auth *ClaudeAuth) loadToken() string {
	data, err := os.ReadFile(auth.configPath())
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

func (auth *ClaudeAuth) saveToken(token string) error {
	path := auth.configPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.Marshal(struct {
		Token string `json:"token"`
	}{token})
	return os.WriteFile(path, data, 0600)
}

func (auth *ClaudeAuth) ensureOAuthToken() error {
	if token := auth.loadToken(); token != "" {
		auth.token = token
		return nil
	}

	fmt.Println("No Claude token found. Running 'claude setup-token'...")
	cmd := exec.Command("claude", "setup-token")
	cmd.Stdin = os.Stdin

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Run(); err != nil {
		fmt.Println(buf.String())
		return err
	}

	re := regexp.MustCompile(`sk-ant-oat01-[A-Za-z0-9_-]+`)
	match := re.FindString(buf.String())
	if match == "" {
		return fmt.Errorf("Could not find token in 'setup-token' output")
	}

	auth.token = match
	return auth.saveToken(match)
}
