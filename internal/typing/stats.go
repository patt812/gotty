package typing

import (
	"fmt"
	"gotty/pkg/display"
	"strings"
	"time"
)

const progressBase = 20

type Stats struct {
	CorrectCount     int
	TotalCount       int
	TotalWPM         float64
	CurrentWPM       float64
	CurrentCorrect   int
	CurrentTotal     int
	Timer            *Timer
	StartTime        time.Time
	IntervalStart    time.Time
	CurrentStartTime time.Time
}

func NewStats() *Stats {
	return &Stats{
		CorrectCount:     0,
		TotalCount:       0,
		TotalWPM:         0,
		CurrentWPM:       0,
		CurrentCorrect:   0,
		CurrentTotal:     0,
		Timer:            NewTimer(),
		StartTime:        time.Now(),
		IntervalStart:    time.Now(),
		CurrentStartTime: time.Now(),
	}
}

func (s *Stats) Update(correct bool) {
	s.TotalCount++
	s.CurrentTotal++
	if correct {
		s.CorrectCount++
		s.CurrentCorrect++
	}
	s.calculateWPM()
}

func (s *Stats) calculateWPM() {
	elapsedTime := time.Since(s.StartTime).Minutes()
	if elapsedTime > 0 {
		s.TotalWPM = float64(s.CorrectCount) / elapsedTime
	}

	currentElapsedTime := time.Since(s.CurrentStartTime).Minutes()
	if currentElapsedTime > 0 {
		s.CurrentWPM = float64(s.CurrentCorrect) / currentElapsedTime
	}
}

func (s *Stats) ResetInterval() {
	s.IntervalStart = time.Now()
	s.CurrentStartTime = time.Now()
	s.CurrentCorrect = 0
	s.CurrentTotal = 0
}

func (s *Stats) GetAccuracy() string {
	return CalculateAccuracy(s.CorrectCount, s.TotalCount)
}

func (s *Stats) GetCurrentWPM() string {
	return fmt.Sprintf("%.2f", s.CurrentWPM)
}

func (s *Stats) GetTotalWPM() string {
	return fmt.Sprintf("%.2f", s.TotalWPM)
}

func (s *Stats) StartTimer(timerLine *display.TerminalLine) {
	s.StartTime = time.Now()
	s.IntervalStart = s.StartTime
	s.CurrentStartTime = s.StartTime
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
	if timerLine == nil {
		fmt.Println("Error: timerLine is nil")
		return
	}

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
