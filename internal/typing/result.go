package typing

import (
	"fmt"
	"gotty/pkg/display"
	"os"

	"github.com/fatih/color"
)

type SentenceResult struct {
	Text     string
	Accuracy string
	WPM      string
}

type Result struct {
	Sentences     []SentenceResult
	TotalWPM      string
	TotalAccuracy string
	TotalTime     string
}

func ShowResult(result Result, onExit func()) {
	display.ClearTerminal()

	lineNumber := 1
	titleLine := display.NewTerminalLine(lineNumber)
	titleLine.SetText("All Sentences:")

	for _, sentence := range result.Sentences {
		lineNumber++
		textLine := display.NewTerminalLine(lineNumber)
		cyan := color.New(color.FgCyan).SprintFunc()
		textLine.SetText(cyan(fmt.Sprintf("%s Accuracy: %s WPM: %s", sentence.Text, sentence.Accuracy, sentence.WPM)))
	}

	lineNumber++
	totalAccuracyLine := display.NewTerminalLine(lineNumber)
	totalAccuracyLine.SetText(fmt.Sprintf("Total Accuracy: %s", result.TotalAccuracy))

	lineNumber++
	totalWPMLine := display.NewTerminalLine(lineNumber)
	totalWPMLine.SetText(fmt.Sprintf("Total WPM: %s", result.TotalWPM))

	lineNumber++
	timeLine := display.NewTerminalLine(lineNumber)
	timeLine.SetText(fmt.Sprintf("Total time: %s seconds", result.TotalTime))

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
