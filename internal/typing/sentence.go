package typing

import (
	"fmt"
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

func (s *Sentence) WPM(stats *Stats) string {
	elapsedTime := time.Since(s.StartTime).Minutes()
	if elapsedTime > 0 {
		wpm := float64(s.CorrectCount) / elapsedTime
		return fmt.Sprintf("%.2f", wpm)
	}
	return "0.00"
}

func GetSentences() []Sentence {
	sentences := []Sentence{
		{Text: "hogehoge"},
		{Text: "fuga"},
	}

	rand.Shuffle(len(sentences), func(i, j int) {
		sentences[i], sentences[j] = sentences[j], sentences[i]
	})

	return sentences
}
