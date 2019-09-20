package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// Config is the config object
type Config struct {
	WebhookURL string
}

func (cfg Config) validate() error {
	if cfg.WebhookURL == "" {
		return errors.New("WebhookURL is empty")
	}
	return nil
}

// LoadConfig reads from `$HOME/.config/sab-notify-discord.json`
func LoadConfig() (*Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filepath.Join(homedir, ".config", "sab-notify-discord.json"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.NewDecoder(fh).Decode(cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.validate()
	if err != nil {
		return nil, err
	}

	err = fh.Close()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
