package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
)

var Config ConfigData

type ConfigData struct {
	NumberOfSentences int `json:"number_of_sentences"`
}

var defaultConfig = ConfigData{
	NumberOfSentences: 2,
}

func LoadConfig() error {
	var configPath string

	env := os.Getenv("GOTTY_ENV")
	if env == "production" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configPath = filepath.Join(homeDir, ".gotty", "config.json")
	} else {
		configPath = "config/config.json"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfigData, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return err
		}
		err = os.WriteFile(configPath, defaultConfigData, 0644)
		if err != nil {
			return err
		}
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		return err
	}

	applyDefaultValues(&Config, defaultConfig)

	return nil
}

func SaveConfig() error {
	var configPath string

	env := os.Getenv("GOTTY_ENV")
	if env == "production" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configPath = filepath.Join(homeDir, ".gotty", "config.json")
	} else {
		configPath = "config/config.json"
	}

	configData, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, configData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func applyDefaultValues(config *ConfigData, defaults ConfigData) {
	vConfig := reflect.ValueOf(config).Elem()
	vDefaults := reflect.ValueOf(defaults)

	for i := 0; i < vConfig.NumField(); i++ {
		field := vConfig.Field(i)
		if field.IsZero() {
			field.Set(vDefaults.Field(i))
		}
	}
}
