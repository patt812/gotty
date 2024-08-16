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

	display.HideCursor()
	defer display.ShowCursor()

	displaySentences(sentences)
	displaySummary(len(sentences), totalTime)

	waitForEscapeKey(onExit)
}

func displaySentences(sentences []Sentence) {
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Print("\033[K")
	fmt.Println("All Sentences:")

	for _, sentence := range sentences {
		fmt.Print("\r\033[K")
		fmt.Printf("%s\n", cyan(sentence.Text))
	}

	fmt.Print("\033[1B")
}

func displaySummary(sentenceCount int, totalTime time.Duration) {
	fmt.Print("\r\033[K")
	fmt.Printf("Number of sentences: %d\n", sentenceCount)
	fmt.Print("\r")
	fmt.Printf("Total time: %02d.%03d seconds\n", int(totalTime.Seconds()), int(totalTime.Milliseconds()%1000))

	fmt.Print("\r\033[K")
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
