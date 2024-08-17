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
