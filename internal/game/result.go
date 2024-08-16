package game

import (
	"fmt"
	"gotty/pkg/display"
	"os"
	"time"

	"github.com/fatih/color"
)

func ShowResult(sentences []Sentence, totalTime time.Duration, onExit func()) {
	display.ClearTerminal()
	displaySentences(sentences)
	displaySummary(len(sentences), totalTime)

	waitForEscapeKey(onExit)
}

func displaySentences(sentences []Sentence) {
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println("All Sentences:")
	for _, sentence := range sentences {
		fmt.Printf("  %s\n", cyan(sentence.Text))
	}
	fmt.Println()
}

func displaySummary(sentenceCount int, totalTime time.Duration) {
	fmt.Printf("Number of sentences: %d\n", sentenceCount)
	fmt.Printf("Total time: %02d.%03d seconds\n", int(totalTime.Seconds()), int(totalTime.Milliseconds()%1000))
	fmt.Println()
	fmt.Println("Press ESC to return to the menu...")
}

func waitForEscapeKey(onExit func()) {
	for {
		var b = make([]byte, 1)
		os.Stdin.Read(b)
		if b[0] == 27 {
			onExit()
			return
		}
	}
}
