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

func convertToKana(romaji string) string {
	kana := ""
	i := 0
	for i < len(romaji) {
		if i+2 < len(romaji) {
			if k, ok := romajiToKanaMap[romaji[i:i+3]]; ok {
				kana += k[0]
				i += 3
				continue
			}
		}
		if i+1 < len(romaji) {
			if k, ok := romajiToKanaMap[romaji[i:i+2]]; ok {
				kana += k[0]
				i += 2
				continue
			}
		}
		if k, ok := romajiToKanaMap[romaji[i:i+1]]; ok {
			kana += k[0]
			i++
			continue
		}
		kana += string(romaji[i])
		i++
	}
	return kana
}
