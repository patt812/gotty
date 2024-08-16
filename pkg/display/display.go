package display

import (
	"fmt"
	"syscall"
	"time"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func UpdateDisplay(target, userInput string) {

	fmt.Print("\r")

	for i, c := range target {
		if i < len(userInput) {
			if c == rune(userInput[i]) {

				color.New(color.FgCyan).Print(string(c))
			} else {

				color.New(color.FgHiYellow).Print(string(userInput[i]))
			}
		} else {

			fmt.Print(string(c))
		}
	}

	fmt.Print("\033[K")
}

func RestoreTerminal(oldState *term.State) {
	err := term.Restore(int(syscall.Stdin), oldState)
	if err != nil {
		fmt.Println("Error restoring terminal:", err)
	}
}

func ShowMissMessage() {
	red := color.New(color.FgRed)
	fmt.Print("\033[1B\r\033[K")
	red.Print("MISS!")
	time.Sleep(1 * time.Second)
	fmt.Print("\r\033[K")
	fmt.Print("\033[1A")
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
