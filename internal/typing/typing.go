package typing

import (
	"bufio"
	"fmt"
	"gotty/config"
	"gotty/pkg/display"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/term"
)

type Play struct {
	Reader         *bufio.Reader
	Judge          Judge
	DisplayManager DisplayManager
	Stats          *Stats
	WaitGroup      *sync.WaitGroup
	Sentences      []Sentence
	CurrentIndex   int
	CurrentInput   string
	CurrentTarget  Sentence
}

func (g *Play) Start(onExit func()) {
	oldState := initializeTerminal()
	defer RestoreTerminal(oldState)

	g.initGame()

	for g.CurrentIndex < len(g.Sentences) {
		g.CurrentTarget = g.Sentences[g.CurrentIndex]
		g.Stats.ResetInterval()
		g.DisplayManager.UpdateDisplay(g.CurrentTarget, g.Judge.GetPatternIndex(), g.Judge.GetCharIndex(), g.Stats)

		userInput := g.handleUserInput(onExit)
		if userInput == "" {
			return
		}

		g.CurrentInput = ""
		g.CurrentIndex++

		if g.CurrentIndex < len(g.Sentences) {
			g.CurrentTarget = g.Sentences[g.CurrentIndex]
			g.Judge = NewJudge(config.Config.InputMode, g.CurrentTarget.RomajiPatterns)
			g.DisplayManager.Initialize()
			g.DisplayManager.UpdateDisplay(g.CurrentTarget, g.Judge.GetPatternIndex(), g.Judge.GetCharIndex(), g.Stats)
			g.DisplayManager.ShowProgress(g.CurrentIndex+1, len(g.Sentences))
			g.updateStats()
		}
	}

	g.Stats.StopTimer()
	g.WaitGroup.Wait()

	ShowResult(g.Sentences, time.Since(g.Stats.StartTime), g.Stats, onExit)
}

func (g *Play) initGame() {
	g.Reader = bufio.NewReader(os.Stdin)
	g.Stats = NewStats()

	g.WaitGroup = &sync.WaitGroup{}
	g.WaitGroup.Add(1)

	var timerLine *display.TerminalLine
	switch dm := g.DisplayManager.(type) {
	case *RomajiDisplayManager:
		timerLine = dm.TimerLine
	case *KanaDisplayManager:
		timerLine = dm.TimerLine
	default:
		timerLine = nil
	}

	go func() {
		defer g.WaitGroup.Done()
		g.Stats.StartTimer(timerLine)
	}()

	g.Sentences = GetSentences()
	g.CurrentIndex = 0
	g.CurrentInput = ""

	g.DisplayManager.Initialize()
	g.updateStats()
}

func (g *Play) handleUserInput(onExit func()) string {
	g.DisplayManager.ShowProgress(g.CurrentIndex+1, len(g.Sentences))

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
			ShowResult(g.Sentences, time.Since(g.Stats.StartTime), g.Stats, onExit)
			return ""
		}

		g.CurrentInput = g.Judge.ProcessInput(char, g.CurrentInput, g.CurrentTarget.Text)

		if g.Judge.IsCorrect(g.CurrentInput, g.CurrentTarget.RomajiPatterns, len(g.CurrentInput)-1) {
			g.Stats.Update(true)
			g.CurrentTarget.UpdateStats(true)

			if g.Judge.ShouldGoNext() {
				g.CurrentIndex++
				if g.CurrentIndex < len(g.Sentences) {
					g.CurrentTarget = g.Sentences[g.CurrentIndex]
					g.CurrentInput = ""
					g.Stats.ResetInterval()
					g.Judge = NewJudge(config.Config.InputMode, g.CurrentTarget.RomajiPatterns)
					g.DisplayManager.Initialize()
					g.DisplayManager.UpdateDisplay(g.CurrentTarget, g.Judge.GetPatternIndex(), g.Judge.GetCharIndex(), g.Stats)
					g.DisplayManager.ShowProgress(g.CurrentIndex+1, len(g.Sentences))
					g.updateStats()
					continue
				} else {
					return g.CurrentInput
				}
			}
		} else {
			g.Stats.Update(false)
			g.CurrentTarget.UpdateStats(false)
			fmt.Println("MISS!")
			g.DisplayManager.ShowMissMessage()
		}

		g.DisplayManager.UpdateDisplay(g.CurrentTarget, g.Judge.GetPatternIndex(), g.Judge.GetCharIndex(), g.Stats)
	}
}

func (g *Play) updateStats() {
	g.DisplayManager.UpdateStats(g.Stats)
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
