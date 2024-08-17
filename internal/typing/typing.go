package typing

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

type Play struct {
	Reader        *bufio.Reader
	Judge         *Judge
	Stats         *Stats
	TextLine      *display.TerminalLine
	MissLine      *display.TerminalLine
	TimerLine     *display.TerminalLine
	StatsLine     *display.TerminalLine
	ProgressLine  *display.TerminalLine
	WaitGroup     *sync.WaitGroup
	Sentences     []Sentence
	CurrentIndex  int
	CurrentInput  string
	CurrentTarget string
}

func (g *Play) Start(onExit func()) {
	oldState := initializeTerminal()
	defer RestoreTerminal(oldState)

	g.initGame()

	for g.CurrentIndex < len(g.Sentences) {
		g.CurrentTarget = g.Sentences[g.CurrentIndex].Text
		g.TextLine.SetText(g.CurrentTarget)

		userInput := g.handleUserInput(onExit)
		if userInput == "" {
			return
		}

		g.CurrentInput = ""
		g.CurrentIndex++

		g.Stats.ResetInterval() // 次の文章に進むときに区間をリセット
	}

	g.Stats.StopTimer()
	g.WaitGroup.Wait()

	ShowResult(g.Sentences, time.Since(g.Stats.Timer.StartTime), g.Stats, onExit)
}

func (g *Play) initGame() {
	g.TextLine = display.NewTerminalLine(1)
	g.MissLine = display.NewTerminalLine(2)
	g.TimerLine = display.NewTerminalLine(3)
	g.StatsLine = display.NewTerminalLine(4)
	g.ProgressLine = display.NewTerminalLine(5)

	g.Reader = bufio.NewReader(os.Stdin)
	g.Stats = NewStats()

	g.WaitGroup = &sync.WaitGroup{}
	g.WaitGroup.Add(1)

	go func() {
		defer g.WaitGroup.Done()
		g.Stats.StartTimer(g.TimerLine)
	}()

	g.Judge = NewJudge()
	g.Sentences = GetSentences()
	g.CurrentIndex = 0
	g.CurrentInput = ""

	g.updateStats()
}

func (g *Play) handleUserInput(onExit func()) string {
	ShowProgressBar(g.CurrentIndex+1, len(g.Sentences), g.ProgressLine)

	for {
		char, _, err := g.Reader.ReadRune()
		if err != nil {
			fmt.Println("\nError reading input:", err)
			os.Exit(1)
		}

		if g.Judge.IsExit(char) {
			fmt.Println("\nGame terminated by Escape key")
			g.Stats.StopTimer()
			g.WaitGroup.Wait()
			ShowResult(g.Sentences, time.Since(g.Stats.Timer.StartTime), g.Stats, onExit)
			return ""
		}

		correct := g.Judge.isCorrect(g.CurrentInput+string(char), g.CurrentTarget, len(g.CurrentInput))
		g.Stats.Update(correct)
		g.Sentences[g.CurrentIndex].UpdateStats(correct)

		g.CurrentInput = g.Judge.ProcessInput(char, g.CurrentInput, g.CurrentTarget)

		g.TextLine.UpdateDisplay(g.CurrentTarget, g.CurrentInput)

		if len(g.CurrentInput) <= len(g.CurrentTarget) && !correct {
			g.MissLine.ShowMissMessage()
			g.CurrentInput = g.CurrentInput[:len(g.CurrentInput)-1]
			g.updateStats()
			continue
		}

		ShowProgressBar(g.CurrentIndex+1, len(g.Sentences), g.ProgressLine)
		g.updateStats()

		if g.CurrentInput == g.CurrentTarget {
			return g.CurrentInput
		}
	}
}

func (g *Play) updateStats() {
	g.StatsLine.SetText(fmt.Sprintf("Accuracy: %s WPM: %s (Current: %s)",
		g.Stats.GetAccuracy(), g.Stats.GetTotalWPM(), g.Stats.GetCurrentWPM()))
}

func initializeTerminal() *term.State {
	display.HideCursor()
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

func RestoreTerminal(oldState *term.State) {
	err := term.Restore(int(syscall.Stdin), oldState)
	display.ShowCursor()
	if err != nil {
		fmt.Println("Error restoring terminal:", err)
	}
}
