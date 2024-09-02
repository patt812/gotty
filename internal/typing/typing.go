package typing

import (
	"bufio"
	"fmt"
	"gotty/config"
	"gotty/pkg/display"
	"math/rand"
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
	SentenceIndex  int
	CurrentInput   string
	CurrentTarget  Sentence
	Results        Result
}

func (g *Play) Start(onExit func()) {
	oldState := initializeTerminal()
	defer RestoreTerminal(oldState)

	g.initGame()

	for g.SentenceIndex < len(g.Sentences) {
		g.CurrentTarget = g.Sentences[g.SentenceIndex]
		g.Judge = NewJudge(config.Config.InputMode, g.Sentences[g.SentenceIndex].RomajiPatterns)
		g.Stats.ResetInterval()
		done, yet := g.Judge.ToString()
		g.DisplayManager.UpdateDisplay(g.CurrentTarget.Text, done, yet, g.Stats)

		userInput := g.handleUserInput(onExit)
		if userInput == "" {
			return
		}

		g.CurrentInput = ""
		g.SentenceIndex++

		if g.SentenceIndex < len(g.Sentences) {
			g.CurrentTarget = g.Sentences[g.SentenceIndex]
			g.Judge = NewJudge(config.Config.InputMode, g.CurrentTarget.RomajiPatterns)
			g.DisplayManager.Initialize()
			g.DisplayManager.UpdateDisplay(g.CurrentTarget.Text, yet, done, g.Stats)
			g.DisplayManager.ShowProgress(g.SentenceIndex+1, len(g.Sentences))
			g.updateStats()
		}
	}

	g.Stats.StopTimer()
	g.WaitGroup.Wait()

	g.Results.TotalAccuracy = g.Stats.GetAccuracy()
	g.Results.TotalWPM = g.Stats.GetTotalWPM()
	g.Results.TotalTime = fmt.Sprintf("%02d.%03d", int(time.Since(g.Stats.StartTime).Seconds()), int(time.Since(g.Stats.StartTime).Milliseconds()%1000))

	ShowResult(g.Results, onExit)
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

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	g.Sentences = GetSentences(GenerateRomajiPatterns, rng)
	g.CurrentInput = ""

	g.DisplayManager.Initialize()
	g.updateStats()

	g.Results = Result{
		Sentences: []SentenceResult{},
	}
}

func (g *Play) handleUserInput(onExit func()) string {
	g.DisplayManager.ShowProgress(g.SentenceIndex+1, len(g.Sentences))

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
			ShowResult(g.Results, onExit)
			return ""
		}

		g.CurrentInput = g.Judge.ProcessInput(char)
		correct := g.Judge.IsCorrect(g.CurrentInput)
		if correct {
			g.Judge.ShiftPosition()
		} else {
			g.DisplayManager.ShowMissMessage()
		}

		g.Stats.Update(correct)
		g.CurrentTarget.UpdateStats(correct)

		if g.Judge.IsNext() {
			g.Results.Sentences = append(g.Results.Sentences, SentenceResult{
				Text:     g.CurrentTarget.Text,
				Accuracy: g.Stats.GetAccuracy(),
				WPM:      g.Stats.GetCurrentWPM(),
			})

			g.SentenceIndex++
			if g.SentenceIndex < len(g.Sentences) {
				g.CurrentTarget = g.Sentences[g.SentenceIndex]
				g.CurrentInput = ""
				g.Stats.ResetInterval()
				g.Judge = NewJudge(config.Config.InputMode, g.CurrentTarget.RomajiPatterns)
				g.DisplayManager.Initialize()

				done, yet := g.Judge.ToString()
				g.DisplayManager.UpdateDisplay(g.CurrentTarget.Text, done, yet, g.Stats)
				g.DisplayManager.ShowProgress(g.SentenceIndex+1, len(g.Sentences))
				g.updateStats()
				continue
			} else if g.SentenceIndex == len(g.Sentences) {
				return g.CurrentInput
			}
		}

		done, yet := g.Judge.ToString()
		g.DisplayManager.UpdateDisplay(g.CurrentTarget.Text, done, yet, g.Stats)
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
