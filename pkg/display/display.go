package display

import (
	"fmt"
)

var ClearTerminal = clearTerminal

var HideCursor = hideCursor

var ShowCursor = showCursor

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}
