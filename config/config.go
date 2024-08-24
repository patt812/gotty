package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigData struct {
	NumberOfSentences int    `json:"number_of_sentences"`
	InputMode         string `json:"input_mode"`
}

type SentenceConfig struct {
	Sentences []string `json:"sentences"`
}

type PatternConfig map[string][]string

var Config ConfigData
var Sentences []string
var Patterns PatternConfig

func LoadConfig() error {
	configPath := filepath.Join("config", "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config.json not found")
	}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		return err
	}

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
	filePath := filepath.Join("config", "sentences.json")
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read sentences.json: %v", err)
	}

	var sentenceConfig SentenceConfig
	err = json.Unmarshal(file, &sentenceConfig)
	if err != nil {
		return fmt.Errorf("failed to unmarshal sentences.json: %v", err)
	}

	Sentences = sentenceConfig.Sentences
	return nil
}

func LoadPatterns() error {
	filePath := filepath.Join("config", "patterns.json")
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
	configPath := filepath.Join("config", "config.json")

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
