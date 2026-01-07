package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	path, err := getConfigFilePath()

	if err != nil {
		return nil, err
	}

	data, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var config Config

	err = json.NewDecoder(data).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	return write(cfg)
}

func write(cfg *Config) error {
	path, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "  ")

	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	home, error := os.UserHomeDir()

	if error != nil {
		return "", error
	}

	return fmt.Sprintf("%s/%s", home, configFileName), nil
}
