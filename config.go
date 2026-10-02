package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const tokenTTL = 365 * 24 * time.Hour

type Config struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	Model     string    `json:"model"`
	Effort    string    `json:"effort"`
}

func DefaultConfig() *Config {
	return &Config{Token: "", ExpiresAt: time.Now(), Model: "haiku", Effort: "n/a"}
}

func path() string {
	return filepath.Join(os.Getenv("APPDATA"), "ktrainer", "config.json")
}

func (config *Config) Write() {
	path := path()
	os.MkdirAll(filepath.Dir(path), 0700)
	data, _ := json.Marshal(config)
	os.WriteFile(path, data, 0600)
}

func (config *Config) Invalidate() {
	DefaultConfig().Write()
}

func InitConfig() *Config {
	data, err := os.ReadFile(path())
	if err != nil {
		DefaultConfig().Write()
		return DefaultConfig()
	}

	var config Config
	if json.Unmarshal(data, &config) != nil {
		DefaultConfig().Write()
		return DefaultConfig()
	}

	return &config
}

func (config *Config) HasValidToken() bool {
	return config.Token != "" && time.Now().Before(config.ExpiresAt)
}

func (config *Config) UpdateToken(token string) *Config {
	newConfig := Config{Token: token, ExpiresAt: time.Now().Add(tokenTTL), Model: config.Model, Effort: config.Effort}
	newConfig.Write()
	return &newConfig
}
