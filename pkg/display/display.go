package display

import (
	"fmt"

	"github.com/fatih/color"
)

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

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}
