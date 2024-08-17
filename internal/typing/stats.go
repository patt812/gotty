package typing

import (
	"fmt"
	"gotty/pkg/display"
	"strings"
	"time"
)

const progressBase = 20

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

type Stats struct {
	CorrectCount int
	TotalCount   int
}

func NewStats() *Stats {
	return &Stats{
		CorrectCount: 0,
		TotalCount:   0,
	}
}

func (s *Stats) Update(correct bool) {
	s.TotalCount++
	if correct {
		s.CorrectCount++
	}
}

func (s *Stats) Accuracy() string {
	if s.TotalCount == 0 {
		return "---"
	}
	return fmt.Sprintf("%.2f%%", float64(s.CorrectCount)/float64(s.TotalCount)*100)
}
