package game

import (
	"bufio"
	"fmt"
	"gotty/pkg/display"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/term"
)

func Start(onExit func()) {
	display.HideCursor()
	defer display.ShowCursor()

	oldState := initializeTerminal()
	defer display.RestoreTerminal(oldState)

	sentences := GetSentences()
	totalSentences := len(sentences)

	reader := bufio.NewReader(os.Stdin)

	startTime := time.Now()
	stopTimer := make(chan struct{})
	pauseTimer := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		RunTimer(startTime, stopTimer, pauseTimer)
	}()

	for i := 0; i < totalSentences; i++ {
		targetText := sentences[i].Text
		displaySentence(targetText, i+1, totalSentences)

		userInput := handleUserInput(reader, targetText, i+1, totalSentences, stopTimer, pauseTimer)
		if userInput == "" {
			safeClose(stopTimer)
			wg.Wait()
			return
		}
	}

	safeClose(stopTimer)
	wg.Wait()

	totalTime := time.Since(startTime)
	ShowResult(sentences, totalTime, onExit)
}

func initializeTerminal() *term.State {
	display.ClearTerminal()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error setting terminal to raw mode:", err)
		os.Exit(1)
	}
	return oldState
}

func displaySentence(sentence string, currentSentence, totalSentences int) {
	display.ClearTerminal()
	fmt.Print(sentence)
	fmt.Print("\r")

	fmt.Print("\033[3B")
	ShowProgressBar(currentSentence, totalSentences)
	fmt.Print("\033[3A")
}

func safeClose(ch chan struct{}) {
	select {
	case <-ch:

	default:
		close(ch)
	}
}

func handleUserInput(reader *bufio.Reader, targetText string, currentSentence, totalSentences int, stopTimer chan struct{}, pauseTimer chan bool) string {
	userInput := ""

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("\nError reading input:", err)
			os.Exit(1)
		}

		if char == 27 {
			fmt.Println("\nGame terminated by Escape key")
			safeClose(stopTimer)
			return ""
		}

		if char == 127 && len(userInput) > 0 {
			userInput = userInput[:len(userInput)-1]
		} else if len(userInput) < len(targetText) {
			userInput += string(char)
		}

		if len(userInput) <= len(targetText) && userInput[len(userInput)-1] != targetText[len(userInput)-1] {

			pauseTimer <- true
			display.ShowMissMessage()
			userInput = userInput[:len(userInput)-1]

			pauseTimer <- false
			continue
		}

		fmt.Print("\r\033[K")
		display.UpdateDisplay(targetText, userInput)

		fmt.Print("\033[3B")
		ShowProgressBar(currentSentence, totalSentences)
		fmt.Print("\033[3A")

		if userInput == targetText {
			return userInput
		}
	}
}
