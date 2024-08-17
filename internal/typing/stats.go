package typing

import (
	"fmt"
	"gotty/pkg/display"
	"strings"
	"time"
)

const progressBase = 20

type Stats struct {
	CorrectCount int
	TotalCount   int
	Timer        *Timer
}

func NewStats() *Stats {
	return &Stats{
		CorrectCount: 0,
		TotalCount:   0,
		Timer:        NewTimer(),
	}
}

func (s *Stats) Update(correct bool) {
	s.TotalCount++
	if correct {
		s.CorrectCount++
	}
}

func (s *Stats) GetAccuracy() string {
	return CalculateAccuracy(s.CorrectCount, s.TotalCount)
}

func (s *Stats) StartTimer(timerLine *display.TerminalLine) {
	go s.Timer.RunTimer(timerLine)
}

func (s *Stats) StopTimer() {
	s.Timer.Stop()
}

func CalculateAccuracy(correctCount, totalCount int) string {
	if totalCount == 0 {
		return "---"
	}
	return fmt.Sprintf("%.2f%%", float64(correctCount)/float64(totalCount)*100)
}

type Timer struct {
	StartTime  time.Time
	StopTimer  chan struct{}
	PauseTimer chan bool
}

func NewTimer() *Timer {
	return &Timer{
		StartTime:  time.Now(),
		StopTimer:  make(chan struct{}),
		PauseTimer: make(chan bool),
	}
}

func (t *Timer) RunTimer(timerLine *display.TerminalLine) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	running := true

	for {
		select {
		case <-t.StopTimer:
			return
		case pause := <-t.PauseTimer:
			running = !pause
		case <-ticker.C:
			if running {
				elapsed := time.Since(t.StartTime)
				timerLine.SetText(fmt.Sprintf("%02d.%03d", int(elapsed.Seconds()), int(elapsed.Milliseconds()%1000)))
			}
		}
	}
}

func (t *Timer) Stop() {
	close(t.StopTimer)
}

func ShowProgressBar(current, total int, progressLine *display.TerminalLine) {
	progress := int(float64(current) / float64(total) * progressBase)
	bar := strings.Repeat("=", progress-1) + ">" + strings.Repeat("-", progressBase-progress)
	progressLine.SetText(fmt.Sprintf("%d / %d [%s]", current, total, bar))
}
