package typing

import (
	"fmt"
	"gotty/pkg/display"

	"github.com/fatih/color"
)

type DisplayManager interface {
	Initialize()
	UpdateDisplay(sentence Sentence, patternIndex int, charIndex int, stats *Stats)
	UpdateStats(stats *Stats)
	ShowMissMessage()
	ShowProgress(current, total int)
}

type RomajiDisplayManager struct {
	TextLine     *display.TerminalLine
	RomajiLine   *display.TerminalLine
	MissLine     *display.TerminalLine
	TimerLine    *display.TerminalLine
	StatsLine    *display.TerminalLine
	ProgressLine *display.TerminalLine
}

func NewRomajiDisplayManager() *RomajiDisplayManager {
	return &RomajiDisplayManager{
		TextLine:     display.NewTerminalLine(1),
		RomajiLine:   display.NewTerminalLine(2),
		MissLine:     display.NewTerminalLine(3),
		TimerLine:    display.NewTerminalLine(4),
		StatsLine:    display.NewTerminalLine(5),
		ProgressLine: display.NewTerminalLine(6),
	}
}

func (d *RomajiDisplayManager) Initialize() {
	display.ClearTerminal()
	if d.TextLine == nil {
		d.TextLine = display.NewTerminalLine(1)
	}
	if d.RomajiLine == nil {
		d.RomajiLine = display.NewTerminalLine(2)
	}
	if d.MissLine == nil {
		d.MissLine = display.NewTerminalLine(3)
	}
	if d.TimerLine == nil {
		d.TimerLine = display.NewTerminalLine(4)
	}
	if d.StatsLine == nil {
		d.StatsLine = display.NewTerminalLine(5)
	}
	if d.ProgressLine == nil {
		d.ProgressLine = display.NewTerminalLine(6)
	}
}

func (d *RomajiDisplayManager) UpdateDisplay(sentence Sentence, currentIndex int, charIndex int, stats *Stats) {
	romajiDisplay := ""
	for i, patterns := range sentence.RomajiPatterns {
		if len(patterns) > 0 {
			if i < currentIndex {
				romajiDisplay += color.New(color.FgCyan).Sprint(patterns[0])
			} else {
				romajiDisplay += patterns[0]
			}
		}
	}
	d.TextLine.SetText(sentence.Text)
	d.RomajiLine.SetText(romajiDisplay)
	d.UpdateStats(stats)
}

func (d *RomajiDisplayManager) UpdateStats(stats *Stats) {
	d.StatsLine.SetText(fmt.Sprintf("Accuracy: %s WPM: %s (Current: %s)",
		stats.GetAccuracy(), stats.GetTotalWPM(), stats.GetCurrentWPM()))
}

func (d *RomajiDisplayManager) ShowMissMessage() {
	d.MissLine.ShowMissMessage()
}

func (d *RomajiDisplayManager) ShowProgress(current, total int) {
	display.ShowProgressBar(current, total, d.ProgressLine)
}

type KanaDisplayManager struct {
	TextLine     *display.TerminalLine
	RomajiLine   *display.TerminalLine
	MissLine     *display.TerminalLine
	TimerLine    *display.TerminalLine
	StatsLine    *display.TerminalLine
	ProgressLine *display.TerminalLine
}

func NewKanaDisplayManager() *KanaDisplayManager {
	return &KanaDisplayManager{
		TextLine:     display.NewTerminalLine(1),
		RomajiLine:   display.NewTerminalLine(2),
		MissLine:     display.NewTerminalLine(3),
		TimerLine:    display.NewTerminalLine(4),
		StatsLine:    display.NewTerminalLine(5),
		ProgressLine: display.NewTerminalLine(6),
	}
}

func (d *KanaDisplayManager) Initialize() {
	display.ClearTerminal()
	if d.TextLine == nil {
		d.TextLine = display.NewTerminalLine(1)
	}
	if d.RomajiLine == nil {
		d.RomajiLine = display.NewTerminalLine(2)
	}
	if d.MissLine == nil {
		d.MissLine = display.NewTerminalLine(3)
	}
	if d.TimerLine == nil {
		d.TimerLine = display.NewTerminalLine(4)
	}
	if d.StatsLine == nil {
		d.StatsLine = display.NewTerminalLine(5)
	}
	if d.ProgressLine == nil {
		d.ProgressLine = display.NewTerminalLine(6)
	}
}

func (d *KanaDisplayManager) UpdateDisplay(sentence Sentence, patternIndex int, charIndex int, stats *Stats) {
	romajiDisplay := ""
	correctColored := ""

	for i, patterns := range sentence.RomajiPatterns {
		if len(patterns) > 0 {
			if i < patternIndex {
				correctColored += color.New(color.FgCyan).Sprint(patterns[0])
			} else if i == patternIndex {
				correctColored += color.New(color.FgCyan).Sprint(patterns[0][:charIndex])
				correctColored += color.New(color.FgWhite).Sprint(patterns[0][charIndex:])
			} else {
				correctColored += color.New(color.FgWhite).Sprint(patterns[0])
			}
			romajiDisplay += patterns[0]
		}
	}

	d.TextLine.SetText(sentence.Text)
	d.RomajiLine.SetText(correctColored)
	d.UpdateStats(stats)
}

func (d *KanaDisplayManager) UpdateStats(stats *Stats) {
	d.StatsLine.SetText(fmt.Sprintf("Accuracy: %s WPM: %s (Current: %s)",
		stats.GetAccuracy(), stats.GetTotalWPM(), stats.GetCurrentWPM()))
}

func (d *KanaDisplayManager) ShowMissMessage() {
	d.MissLine.ShowMissMessage()
}

func (d *KanaDisplayManager) ShowProgress(current, total int) {
	display.ShowProgressBar(current, total, d.ProgressLine)
}
