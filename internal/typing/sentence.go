package typing

import (
	"fmt"
	"gotty/config"
	"math/rand"
	"time"
)

type Sentence struct {
	Text         string
	CorrectCount int
	TotalCount   int
	StartTime    time.Time
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
	return CalculateAccuracy(s.CorrectCount, s.TotalCount)
}

func (s *Sentence) WPM() string {
	elapsedTime := time.Since(s.StartTime).Minutes()
	if elapsedTime > 0 {
		wpm := float64(s.CorrectCount) / elapsedTime
		return fmt.Sprintf("%.2f", wpm)
	}
	return "0.00"
}

func GetSentences() []Sentence {
	baseSentences := []string{"hogehoge", "fuga"}

	totalSentences := config.Config.NumberOfSentences
	sentences := make([]Sentence, totalSentences)

	for i := 0; i < totalSentences; i++ {
		randomIndex := rand.Intn(len(baseSentences))
		sentences[i] = Sentence{Text: baseSentences[randomIndex]}
	}

	return sentences
}
