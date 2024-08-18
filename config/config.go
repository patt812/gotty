package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type ConfigData struct {
	NumberOfSentences int    `json:"number_of_sentences"`
	InputMode         string `json:"input_mode"`
}

var defaultConfig = ConfigData{
	NumberOfSentences: 2,
	InputMode:         "romaji",
}

type SentenceConfig struct {
	Sentences []string `json:"sentences"`
}

type PatternConfig map[string][]string

var Config ConfigData
var Sentences []string
var Patterns PatternConfig

func LoadConfig() error {
	var configPath string

	env := os.Getenv("APP_ENV")
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

	err = LoadSentences()
	if err != nil {
		return err
	}

	err = LoadPatterns()
	if err != nil {
		return err
	}

	return nil
}

func LoadSentences() error {
	var sentenceConfig SentenceConfig

	filePath := filepath.Join("config", "sentences.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		defaultFilePath := filepath.Join("config", "default_sentences.json")
		defaultData, err := os.ReadFile(defaultFilePath)
		if err != nil {
			return fmt.Errorf("failed to read default_sentences.json: %v", err)
		}

		err = os.WriteFile(filePath, defaultData, 0644)
		if err != nil {
			return fmt.Errorf("failed to create sentences.json: %v", err)
		}
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read sentences.json: %v", err)
	}

	err = json.Unmarshal(file, &sentenceConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal sentences.json: %v", err)
	}

	if len(sentenceConfig.Sentences) == 0 {
		return fmt.Errorf("sentences.json contains no sentences")
	}

	Sentences = sentenceConfig.Sentences
	return nil
}

func LoadPatterns() error {
	filePath := filepath.Join("config", "patterns.json")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		defaultFilePath := filepath.Join("config", "default_patterns.json")
		defaultData, err := os.ReadFile(defaultFilePath)
		if err != nil {
			return fmt.Errorf("failed to read default_patterns.json: %v", err)
		}

		err = os.WriteFile(filePath, defaultData, 0644)
		if err != nil {
			return fmt.Errorf("failed to create patterns.json: %v", err)
		}
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read patterns.json: %v", err)
	}

	err = json.Unmarshal(file, &Patterns)
	if err != nil {
		return fmt.Errorf("failed to unmarshal patterns.json: %v", err)
	}

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
