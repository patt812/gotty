package typing

import (
	"fmt"
	"gotty/config"
	"math/rand"
	"time"
)

type Sentence struct {
	Text           string
	RomajiPatterns [][]string
	CorrectCount   int
	TotalCount     int
	StartTime      time.Time
}

func (s *Sentence) UpdateStats(correct bool) {
	if s.TotalCount == 0 {
		s.StartTime = time.Now()
	}
	s.TotalCount++
	if correct {
		s.CorrectCount++
	}
}

func (s *Sentence) Accuracy() string {
	if s.TotalCount == 0 {
		return "---"
	}
	return CalculateAccuracy(s.CorrectCount, s.TotalCount)
}

func (s *Sentence) WPM() string {
	if s.TotalCount == 0 || s.StartTime.IsZero() {
		return "0.00"
	}

	elapsedTime := time.Since(s.StartTime).Minutes()
	if elapsedTime > 0 {
		wpm := float64(s.CorrectCount) / elapsedTime
		return fmt.Sprintf("%.2f", wpm)
	}
	return "0.00"
}

func GetSentences(generatePatterns func(string) [][]string, rng *rand.Rand) []Sentence {
	totalSentences := config.Config.NumberOfSentences
	indices := rng.Perm(len(config.Sentences))

	selectedSentences := make([]Sentence, totalSentences)

	for i := 0; i < totalSentences; i++ {
		randomIndex := indices[i]
		text := config.Sentences[randomIndex]
		selectedSentences[i] = Sentence{
			Text:           text,
			RomajiPatterns: generatePatterns(text),
		}
	}

	return selectedSentences
}
