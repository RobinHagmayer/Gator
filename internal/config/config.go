package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	path := getConfigFilePath()

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read file:", err)
		os.Exit(1)
	}

	config := Config{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to unpack JSON")
		os.Exit(1)
	}

	return config
}

func (c *Config) SetUser(userName string) {
	c.CurrentUserName = userName
	path := getConfigFilePath()

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal JSON:", err)
		os.Exit(1)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write file:", err)
		os.Exit(1)
	}
}

func getConfigFilePath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not find $HOME")
		os.Exit(1)
	}

	return userHomeDir + "/" + configFileName
}
