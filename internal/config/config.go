package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	URL             string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeLocation, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error retrieving home directory")
		return "", err
	}

	fullURL := homeLocation + "/" + configFileName

	return fullURL, nil
}

func Read() (Config, error) {
	fullURL, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(fullURL)
	if err != nil {
		fmt.Println("Error reading from file")
		return Config{}, err
	}

	returnConfig := Config{}
	err = json.Unmarshal(data, &returnConfig)
	if err != nil {
		fmt.Println("Error unmarshalling data")
		return Config{}, err
	}

	return returnConfig, nil
}

func write(cfg Config) error {
	fullURL, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		fmt.Println("Error marshalling config data")
		return err
	}

	err = os.WriteFile(fullURL, cfgBytes, 0600)
	if err != nil {
		fmt.Println("Error writing to file")
		return err
	}

	return nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username

	return write(*c)
}
