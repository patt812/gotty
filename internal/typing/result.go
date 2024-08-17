package typing

import (
	"fmt"
	"gotty/pkg/display"
	"os"
	"time"

	"github.com/fatih/color"
)

func ShowResult(sentences []Sentence, totalTime time.Duration, stats *Stats, onExit func()) {
	display.ClearTerminal()

	lineNumber := 1

	titleLine := display.NewTerminalLine(lineNumber)
	titleLine.SetText("All Sentences:")

	for _, sentence := range sentences {
		lineNumber++
		textLine := display.NewTerminalLine(lineNumber)
		cyan := color.New(color.FgCyan).SprintFunc()
		textLine.SetText(cyan(fmt.Sprintf("%s %s %s", sentence.Text, sentence.Accuracy(), sentence.WPM(stats))))
	}

	lineNumber++
	totalAccuracyLine := display.NewTerminalLine(lineNumber)
	totalAccuracyLine.SetText(fmt.Sprintf("Total Accuracy: %s", stats.GetAccuracy()))

	lineNumber++
	totalWPMLine := display.NewTerminalLine(lineNumber)
	totalWPMLine.SetText(fmt.Sprintf("Total WPM: %s", stats.GetTotalWPM()))

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
