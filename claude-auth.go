package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const tokenTTL = 365 * 24 * time.Hour

type ClaudeAuth struct {
	token string
}

type tokenJson struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewClaudeAuth(token string) *ClaudeAuth {
	return &ClaudeAuth{token: token}
}

func configPath() string {
	return filepath.Join(os.Getenv("APPDATA"), "kata-trainer", "config.json")
}

func LoadTokenFromConfig() string {
	data, err := os.ReadFile(configPath())
	if err != nil {
		return ""
	}
	var tj tokenJson
	if json.Unmarshal(data, &tj) != nil {
		return ""
	}
	if time.Now().After(tj.ExpiresAt) {
		return ""
	}
	return tj.Token
}

func SaveTokenToConfig(token string, expires_at time.Time) {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.Marshal(tokenJson{
		Token:     token,
		ExpiresAt: expires_at})
	os.WriteFile(path, data, 0600)
}

func (auth *ClaudeAuth) Invalidate() {
	SaveTokenToConfig(auth.token, time.Now())
}
