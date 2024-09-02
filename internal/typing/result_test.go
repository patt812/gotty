package typing_test

import (
	"gotty/internal/typing"
	"gotty/pkg/display"
	"os"
	"testing"
)

func SetupTestInput(t *testing.T, onExitCalled *bool) func() {
	oldStdin := os.Stdin
	t.Cleanup(func() { os.Stdin = oldStdin })

	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		defer w.Close()
		w.Write([]byte{27})
	}()

	return func() {
		*onExitCalled = true
	}
}

func TestShowResult_MultipleSentences(t *testing.T) {
	display.ClearTerminal = func() {}

	sut := typing.Result{
		Sentences: []typing.SentenceResult{
			{Text: "こんにちは", Accuracy: "100%", WPM: "10"},
			{Text: "さようなら", Accuracy: "50%", WPM: "5"},
		},
		TotalWPM:      "7.5",
		TotalAccuracy: "75%",
		TotalTime:     "10",
	}

	called := false
	onExit := SetupTestInput(t, &called)

	typing.ShowResult(sut, onExit)

	if !called {
		t.Errorf("Expected onExit to be called, but it wasn't")
	}
}
