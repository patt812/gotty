package typing

import (
	"math/rand"
)

type Sentence struct {
	Text         string
	CorrectCount int
	TotalCount   int
}

func (s *Sentence) UpdateStats(correct bool) {
	s.TotalCount++
	if correct {
		s.CorrectCount++
	}
}

func (s *Sentence) Accuracy() string {
	return CalculateAccuracy(s.CorrectCount, s.TotalCount)
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
