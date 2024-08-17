package typing

import (
	"fmt"
	"gotty/pkg/display"
	"os"
	"time"

	"github.com/fatih/color"
)

func ShowResult(sentences []Sentence, totalTime time.Duration, onExit func()) {
	display.ClearTerminal()

	lineNumber := 1

	titleLine := display.NewTerminalLine(lineNumber)
	titleLine.SetText("All Sentences:")

	for _, sentence := range sentences {
		lineNumber++
		textLine := display.NewTerminalLine(lineNumber)
		cyan := color.New(color.FgCyan).SprintFunc()
		textLine.SetText(cyan(sentence.Text))
	}

	lineNumber++
	summaryLine := display.NewTerminalLine(lineNumber)
	summaryLine.SetText(fmt.Sprintf("Number of sentences: %d", len(sentences)))

	lineNumber++
	timeLine := display.NewTerminalLine(lineNumber)
	timeLine.SetText(fmt.Sprintf("Total time: %02d.%03d seconds", int(totalTime.Seconds()), int(totalTime.Milliseconds()%1000)))

	lineNumber++
	exitLine := display.NewTerminalLine(lineNumber)
	exitLine.SetText("Press ESC to return to the menu")

	waitForEscapeKey(onExit)
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
