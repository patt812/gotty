package display

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type TerminalLine struct {
	LineNumber int
	mu         sync.Mutex
}

func NewTerminalLine(lineNumber int) *TerminalLine {
	return &TerminalLine{LineNumber: lineNumber}
}

func (tl *TerminalLine) SetText(text string) {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.moveToLine()
	fmt.Print("\r\033[K")
	fmt.Print(text)
}

func (tl *TerminalLine) Clear() {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.moveToLine()
	fmt.Print("\r\033[K")
}

func (tl *TerminalLine) ShowMissMessage() {
	tl.mu.Lock()
	defer tl.mu.Unlock()
	tl.moveToLine()
	red := color.New(color.FgRed)
	fmt.Print("\r\033[K")
	red.Print("MISS!")

	go func() {
		time.Sleep(1 * time.Second)
		tl.Clear()
	}()
}

func (tl *TerminalLine) moveToLine() {
	fmt.Printf("\033[%d;0H", tl.LineNumber)
}

func (tl *TerminalLine) UpdateDisplay(target, userInput string) {
	tl.moveToLine()
	fmt.Print("\r\033[K")

	for i, c := range target {
		if i < len(userInput) {
			if c == rune(userInput[i]) {
				color.New(color.FgCyan).Print(string(c))
			} else {
				fmt.Print(string(c))
			}
		} else {
			fmt.Print(string(c))
		}
	}
	fmt.Print("\033[K")
}

func ShowProgressBar(current, total int, progressLine *TerminalLine) {
	progress := int(float64(current) / float64(total) * 20)
	bar := fmt.Sprintf("[%s>%s]",
		string(repeatRune('=', progress)),
		string(repeatRune('-', 20-progress)))
	progressLine.SetText(fmt.Sprintf("%d / %d %s", current, total, bar))
}

func repeatRune(r rune, count int) []rune {
	result := make([]rune, count)
	for i := range result {
		result[i] = r
	}
	return result
}
