package typing_test

import (
	"gotty/internal/typing"
	"gotty/pkg/display"
	"gotty/test/utils"
	"reflect"
	"testing"
)

func TestKana_Initialize(t *testing.T) {
	t.Parallel()

	display.ClearTerminal = func() {}
	sut := typing.NewKanaDisplayManager()
	sut.Initialize()

	want := typing.NewKanaDisplayManager()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected display manager to be %+v, but got %+v", want, sut)
	}
}

func TestKana_Initialize_NoConstructor(t *testing.T) {
	t.Parallel()

	display.ClearTerminal = func() {}
	sut := typing.KanaDisplayManager{}
	sut.Initialize()

	want := typing.KanaDisplayManager{
		TextLine:     &display.TerminalLine{LineNumber: 1},
		RomajiLine:   &display.TerminalLine{LineNumber: 2},
		MissLine:     &display.TerminalLine{LineNumber: 3},
		TimerLine:    &display.TerminalLine{LineNumber: 4},
		StatsLine:    &display.TerminalLine{LineNumber: 5},
		ProgressLine: &display.TerminalLine{LineNumber: 6},
	}

	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected display manager to be %+v, but got %+v", want, sut)
	}
}

func TestKana_UpdateDisplay(t *testing.T) {
	rec := utils.NewLogRecorder()

	stats := typing.NewStats()

	sut := typing.NewKanaDisplayManager()
	sut.UpdateDisplay("こんにちは", "こん", "にちは", stats)

	got := rec.ToArray()
	want := []string{
		"こんにちは",
		"こんにちは",
		"Accuracy: --- WPM: 0.00 (Current: 0.00)",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected output to be %+v, but got %+v", want, got)
	}
}

func TestKana_ShowMissMessage(t *testing.T) {
	rec := utils.NewLogRecorder()

	sut := typing.NewKanaDisplayManager()
	sut.ShowMissMessage()

	got := rec.ToArray()
	want := []string{
		"MISS!",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected output to be %+v, but got %+v", want, got)
	}
}

func TestKana_ShowProgress(t *testing.T) {
	rec := utils.NewLogRecorder()

	sut := typing.NewKanaDisplayManager()
	sut.ShowProgress(1, 10)

	got := rec.ToArray()
	want := []string{
		"1 / 10 [==>------------------]",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected output to be %+v, but got %+v", want, got)
	}
}
