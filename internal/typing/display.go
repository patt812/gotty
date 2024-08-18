package typing

import (
	"fmt"
	"gotty/pkg/display"
)

type DisplayManager interface {
	Initialize()
	UpdateDisplay(sentence Sentence, currentInput string)
	UpdateStats(stats *Stats)
	ShowMissMessage()
	ShowProgress(current, total int)
}

type RomajiDisplayManager struct {
	TextLine     *display.TerminalLine
	MissLine     *display.TerminalLine
	StatsLine    *display.TerminalLine
	ProgressLine *display.TerminalLine
}

func NewRomajiDisplayManager() *RomajiDisplayManager {
	return &RomajiDisplayManager{
		TextLine:     display.NewTerminalLine(1),
		MissLine:     display.NewTerminalLine(2),
		StatsLine:    display.NewTerminalLine(3),
		ProgressLine: display.NewTerminalLine(4),
	}
}

func (d *RomajiDisplayManager) Initialize() {
	display.ClearTerminal()
	if d.TextLine == nil {
		d.TextLine = display.NewTerminalLine(1)
	}
	if d.MissLine == nil {
		d.MissLine = display.NewTerminalLine(2)
	}
	if d.StatsLine == nil {
		d.StatsLine = display.NewTerminalLine(3)
	}
	if d.ProgressLine == nil {
		d.ProgressLine = display.NewTerminalLine(4)
	}
}

func (d *RomajiDisplayManager) UpdateDisplay(sentence Sentence, currentInput string) {
	d.TextLine.UpdateDisplay(sentence.Text, currentInput)
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
	KanaLine     *display.TerminalLine
	MissLine     *display.TerminalLine
	StatsLine    *display.TerminalLine
	ProgressLine *display.TerminalLine
}

func NewKanaDisplayManager() *KanaDisplayManager {
	return &KanaDisplayManager{
		TextLine:     display.NewTerminalLine(1),
		KanaLine:     display.NewTerminalLine(2),
		MissLine:     display.NewTerminalLine(3),
		StatsLine:    display.NewTerminalLine(4),
		ProgressLine: display.NewTerminalLine(5),
	}
}

func (d *KanaDisplayManager) Initialize() {
	display.ClearTerminal()
	if d.TextLine == nil {
		d.TextLine = display.NewTerminalLine(1)
	}
	if d.KanaLine == nil {
		d.KanaLine = display.NewTerminalLine(2)
	}
	if d.MissLine == nil {
		d.MissLine = display.NewTerminalLine(3)
	}
	if d.StatsLine == nil {
		d.StatsLine = display.NewTerminalLine(4)
	}
	if d.ProgressLine == nil {
		d.ProgressLine = display.NewTerminalLine(5)
	}
}

func (d *KanaDisplayManager) UpdateDisplay(sentence Sentence, currentInput string) {
	d.TextLine.SetText(sentence.Text)
	d.KanaLine.SetText(sentence.Kana)
	d.TextLine.UpdateDisplay(sentence.Text, currentInput)
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
