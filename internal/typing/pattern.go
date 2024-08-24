package typing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var romajiToKanaMap map[string][]string

func LoadPatterns() error {
	patternFilePath := filepath.Join("config", "patterns.json")
	file, err := os.Open(patternFilePath)

	if err != nil {
		return fmt.Errorf("failed to open pattern file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&romajiToKanaMap)
	if err != nil {
		return fmt.Errorf("failed to decode pattern file: %v", err)
	}

	return nil
}

func GenerateRomajiPatterns(text string) [][]string {
	if err := LoadPatterns(); err != nil {
		fmt.Println("Error loading patterns:", err)
		return nil
	}

	var romajiPatterns [][]string
	for _, char := range text {
		if patterns, exists := romajiToKanaMap[string(char)]; exists {
			romajiPatterns = append(romajiPatterns, patterns)
		} else {
			romajiPatterns = append(romajiPatterns, []string{string(char)})
		}
	}
	return romajiPatterns
}
