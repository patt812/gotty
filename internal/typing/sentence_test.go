package typing_test

import (
	"gotty/config"
	"gotty/internal/typing"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func setupConfigMock(numberOfSentences int, sentences []string) {
	config.Config.NumberOfSentences = numberOfSentences
	config.Sentences = sentences
	config.Patterns = config.PatternConfig{
		"こんにちは": {"konnichiwa"},
		"さようなら": {"sayounara"},
	}
}

func TestGetSentences(t *testing.T) {
	t.Parallel()

	rng := rand.New(rand.NewSource(1))
	setupConfigMock(2, []string{"こんにちは", "さようなら"})
	mockGeneratePatterns := func(text string) [][]string {
		return [][]string{{"mocked_pattern"}}
	}

	got := typing.GetSentences(mockGeneratePatterns, rng)
	want := []typing.Sentence{
		{Text: "こんにちは", RomajiPatterns: [][]string{{"mocked_pattern"}}},
		{Text: "さようなら", RomajiPatterns: [][]string{{"mocked_pattern"}}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected sentences to be %+v, but got %+v", want, got)
	}
}

func TestUpdateStats_Correct(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	sut.UpdateStats(true)
	want := typing.Sentence{
		CorrectCount: 1,
		TotalCount:   1,
		StartTime:    sut.StartTime,
	}

	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected sentence to be %+v, but got %+v", want, sut)
	}
}

func TestUpdateStats_FirstUpdate(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	sut.UpdateStats(true)

	if sut.StartTime.IsZero() {
		t.Error("Expected start time to be set, but it was not")
	}
}

func TestUpdateStats_Incorrect(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	sut.UpdateStats(false)
	want := typing.Sentence{
		CorrectCount: 0,
		TotalCount:   1,
		StartTime:    sut.StartTime,
	}

	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected sentence to be %+v, but got %+v", want, sut)
	}
}

func TestAccuracy_ZeroTotalCount(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	got := sut.Accuracy()
	want := "---"

	if got != want {
		t.Errorf("Expected accuracy to be %s, but got %s", want, got)
	}
}

func TestWPM_ZeroTotalCount(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	got := sut.WPM()
	want := "0.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}

func TestWPM_StartTimeNotSet(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{}
	got := sut.WPM()
	want := "0.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}

func TestWPM_ZeroElapsedTime(t *testing.T) {
	t.Parallel()

	sut := typing.Sentence{
		CorrectCount: 10,
		StartTime:    time.Now(),
	}
	got := sut.WPM()
	want := "0.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}

func TestWPM_NonZeroElapsedTime(t *testing.T) {
	t.Parallel()

	startTime := time.Now().Add(-time.Minute)
	sut := typing.Sentence{
		TotalCount:   100,
		CorrectCount: 100,
		StartTime:    startTime,
	}
	got := sut.WPM()
	want := "100.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}

func TestWPM_ElapsedTimeGreaterThanZero(t *testing.T) {
	t.Parallel()

	startTime := time.Now().Add(-30 * time.Second)
	sut := typing.Sentence{
		TotalCount:   100,
		CorrectCount: 50,
		StartTime:    startTime,
	}
	got := sut.WPM()
	want := "100.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}

func TestWPM_ElapsedTimeLessThanZero(t *testing.T) {
	t.Parallel()

	startTime := time.Now().Add(30 * time.Second)
	sut := typing.Sentence{
		TotalCount:   100,
		CorrectCount: 50,
		StartTime:    startTime,
	}
	got := sut.WPM()
	want := "0.00"

	if got != want {
		t.Errorf("Expected WPM to be %s, but got %s", want, got)
	}
}
