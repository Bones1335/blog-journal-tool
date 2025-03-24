package config

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Editor       string `yaml:"editor"`
	SaveLocation string `yaml:"save_location"`
}

const configFileName = "/.blog.yaml"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	readYAML, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	decoder := yaml.NewDecoder(bytes.NewReader(readYAML))
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homePath + configFileName, nil
}
