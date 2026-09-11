package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

const tokenTTL = 365 * 24 * time.Hour

type ClaudeAuth struct {
	token string
}

type tokenJson struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

func NewClaudeAuth() *ClaudeAuth {
	return &ClaudeAuth{}
}

func (auth *ClaudeAuth) configPath() string {
	return filepath.Join(os.Getenv("APPDATA"), "kata-trainer", "config.json")
}

func (auth *ClaudeAuth) loadToken() *tokenJson {
	data, err := os.ReadFile(auth.configPath())
	if err != nil {
		return nil
	}
	var tj tokenJson
	if json.Unmarshal(data, &tj) != nil {
		return nil
	}
	if time.Since(tj.CreatedAt) > tokenTTL {
		return nil
	}
	return &tj
}

func (auth *ClaudeAuth) saveToken(token string) error {
	path := auth.configPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.Marshal(tokenJson{Token: token, CreatedAt: time.Now()})
	return os.WriteFile(path, data, 0600)
}

func (auth *ClaudeAuth) EnsureToken() error {
	if tokenJson := auth.loadToken(); tokenJson != nil {
		auth.token = tokenJson.Token
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
