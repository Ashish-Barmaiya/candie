package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	AppName        = "candie"
	ConfigFileName = "config.yaml"
)

func configDir() (string, error) {
	baseDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(baseDir, AppName), nil
}

func configPath() (string, error) {
	dir, err := configDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, ConfigFileName), nil
}

func Save(cfg Config) error {
	dir, err := configDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func Load() (Config, error) {
	path, err := configPath()
	if err != nil {
		return Config{}, err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfg := DefaultConfig()

		if err := Save(cfg); err != nil {
			return Config{}, err
		}

		return cfg, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
