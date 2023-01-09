package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const Version = "1.0.0"

type GitRepoConfig struct {
	Name       string
	Version    SymVer
	Visibility Visibility
	Backends   []Backend
}

func GetRepoConfig(path string) (GitRepoConfig, error) {
	config := GitRepoConfig{}

	jsonFile, err := os.Open(path)
	if err != nil {
		return config, fmt.Errorf("parseRepoFile: Failed to open file: %v", err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return config, fmt.Errorf("parseRepoFile: Failed to read file: %v", err)
	}

	err = json.Unmarshal(byteValue, &config)

	if err != nil {
		return config, fmt.Errorf("parseRepoFile: Failed to parse json: %v", err)
	}

	if config.Version.String() != Version {
		fmt.Println("Config version does not match, some things may not work")
	}

	return config, nil
}
