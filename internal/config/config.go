package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	Db_url          string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	config := Config{}

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File not found")
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Couldn't Unmarshal JSON!")
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != err {
		return "", err
	}

	path := homePath + "/" + configFileName

	return path, nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}
