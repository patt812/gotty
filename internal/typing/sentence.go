package typing

import (
	"math/rand"
)

type Sentence struct {
	Text string
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
