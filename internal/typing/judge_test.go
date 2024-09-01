package typing_test

import (
	"gotty/internal/typing"
	"reflect"
	"slices"
	"testing"
)

func TestNewJudge_Kana(t *testing.T) {
	t.Parallel()

	sut := typing.NewJudge("kana", [][]string{{"a"}})
	want := typing.NewKanaJudge([][]string{{"a"}})

	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_IsCorrect_SingleCharacter_CorrectInput(t *testing.T) {
	t.Parallel()

	patterns := [][]string{
		{"a"},
	}
	judge := typing.NewKanaJudge(patterns)

	if !judge.IsCorrect("a") {
		t.Errorf("Expected IsCorrect('a') to be true, but got false")
	}
}

func TestKanaJudge_IsCorrect_SingleCharacter_CorrectInput_NoRomaIndexChange(t *testing.T) {
	t.Parallel()

	patterns := [][]string{
		{"a"},
	}
	judge := typing.NewKanaJudge(patterns)

	if judge.RomaIndex != 0 {
		t.Errorf("Expected RomaIndex to be 0, but got %d", judge.RomaIndex)
	}
}

func TestKanaJudge_IsCorrect_SingleCharacter_CorrectInput_NoKanaIndexChange(t *testing.T) {
	t.Parallel()

	patterns := [][]string{
		{"a"},
	}
	judge := typing.NewKanaJudge(patterns)

	if judge.KanaIndex != 0 {
		t.Errorf("Expected KanaIndex to be 0, but got %d", judge.KanaIndex)
	}
}

func TestKanaJudge_IsCorrect_SingleCharacter_IncorrectInput(t *testing.T) {
	t.Parallel()

	patterns := [][]string{
		{"a"},
	}
	judge := typing.NewKanaJudge(patterns)

	if judge.IsCorrect("b") {
		t.Errorf("Expected IsCorrect('b') to be false, but got true")
	}
}

func TestKanaJudge_IsCorrect_MultiCharacter_Indexed(t *testing.T) {
	t.Parallel()
	patterns := [][]string{
		{"o"},
		{"ha"},
		{"yo"},
	}
	judge := typing.NewKanaJudge(patterns)

	tests := []struct {
		input     string
		expected  bool
		kanaIndex int
		romaIndex int
	}{
		{"o", true, 0, 0},
		{"h", true, 1, 0},
		{"a", true, 1, 1},
		{"y", true, 2, 0},
		{"o", true, 2, 1},
		{"x", false, 2, 0},
		{"z", false, 2, 0},
	}

	for _, tt := range tests {
		judge.KanaIndex = tt.kanaIndex
		judge.RomaIndex = tt.romaIndex
		if judge.IsCorrect(tt.input) != tt.expected {
			t.Errorf("IsCorrect('%s') at KanaIndex: %d, RomaIndex: %d = %v, expected %v",
				tt.input, tt.kanaIndex, tt.romaIndex, !tt.expected, tt.expected)
		}
	}
}

func TestKanaJudge_IsCorrect_MultiplePatterns_SingleCharacter(t *testing.T) {
	t.Parallel()
	originalPatterns := [][]string{
		{"o", "wo"},
	}

	tests := []struct {
		input    string
		expected bool
	}{
		{"o", true},
		{"w", true},
		{"x", false},
	}

	for _, tt := range tests {
		patterns := slices.Clone(originalPatterns)
		judge := typing.NewKanaJudge(patterns)
		if judge.IsCorrect(tt.input) != tt.expected {
			t.Errorf("IsCorrect('%s') = %v, expected %v", tt.input, !tt.expected, tt.expected)
		}
	}
}

func TestKanaJudge_IsCorrect_Patterns(t *testing.T) {
	t.Parallel()
	patterns := [][]string{
		{"a", "ka", "kya", "xyz"},
	}

	tests := []struct {
		input            string
		kanaIndex        int
		romaIndex        int
		expectedPatterns []string
	}{
		{"k", 0, 0, []string{"ka", "kya"}},
		{"x", 0, 0, []string{"xyz"}},
		{"a", 0, 0, []string{"a"}},
		{"y", 0, 1, []string{"kya"}},
		{"z", 0, 0, []string{"a", "ka", "kya", "xyz"}},
	}

	for _, tt := range tests {
		judge := typing.KanaJudge{
			KanaIndex: tt.kanaIndex,
			RomaIndex: tt.romaIndex,
			Patterns:  slices.Clone(patterns),
		}

		judge.IsCorrect(tt.input)
		for i, pattern := range tt.expectedPatterns {
			if judge.Patterns[tt.kanaIndex][i] != pattern {
				t.Errorf("Expected pattern '%s', but got '%s'", pattern, judge.Patterns[tt.kanaIndex][i])
			}
		}
	}
}

func TestKanaJudge_ShiftPosition_IndexUpdated(t *testing.T) {
	t.Parallel()

	sut := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 0,
		Patterns:  [][]string{{"ab"}},
	}

	want := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 1,
		Patterns:  [][]string{{"ab"}},
	}

	sut.ShiftPosition()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_ShiftPosition_KanaIndexNotUpdated(t *testing.T) {
	t.Parallel()

	sut := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 1,
		Patterns:  [][]string{{"abc"}},
	}

	want := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 2,
		Patterns:  [][]string{{"abc"}},
	}

	sut.ShiftPosition()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_ShiftPosition_KanaIndexUpdated(t *testing.T) {
	t.Parallel()

	sut := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 2,
		Patterns:  [][]string{{"abc"}},
	}

	want := typing.KanaJudge{
		KanaIndex: 1,
		RomaIndex: 0,
		Patterns:  [][]string{{"abc"}},
	}

	sut.ShiftPosition()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_IsNext_IndexUpdated(t *testing.T) {
	t.Parallel()

	sut := typing.KanaJudge{
		KanaIndex: 1,
		RomaIndex: 0,
		Patterns:  [][]string{{"ab"}},
	}

	want := typing.KanaJudge{
		KanaIndex: 2,
		RomaIndex: 0,
		Patterns:  [][]string{{"ab"}},
	}

	sut.IsNext()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_IsNext_IndexNotUpdated(t *testing.T) {
	t.Parallel()

	sut := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 1,
		Patterns:  [][]string{{"ab"}},
	}

	want := typing.KanaJudge{
		KanaIndex: 0,
		RomaIndex: 1,
		Patterns:  [][]string{{"ab"}},
	}

	sut.IsNext()
	if !reflect.DeepEqual(sut, want) {
		t.Errorf("Expected judge to be %+v, but got %+v", want, sut)
	}
}

func TestKanaJudge_ProcessInput(t *testing.T) {
	t.Parallel()

	sut := typing.NewKanaJudge([][]string{{"a"}})
	got := sut.ProcessInput(97)

	if got != "a" {
		t.Errorf("Expected ProcessInput('a') to return 'a', but got %s", got)
	}
}

func TestKanaJudge_IsExit(t *testing.T) {
	t.Parallel()

	sut := typing.NewKanaJudge([][]string{{"a"}})
	got := sut.IsExit(27)

	if !got {
		t.Errorf("Expected IsExit(27) to return true, but got false")
	}
}

func TestKanaJudge_IsExit_NotEscapeKey(t *testing.T) {
	t.Parallel()

	sut := typing.NewKanaJudge([][]string{{"a"}})
	got := sut.IsExit(28)

	if got {
		t.Errorf("Expected IsExit(28) to return false, but got true")
	}
}

func TestKanaJudge_ToString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		judge             typing.KanaJudge
		expectedCorrect   string
		expectedRemaining string
	}{
		{
			name: "joined correct input",
			judge: typing.KanaJudge{
				KanaIndex: 3,
				RomaIndex: 1,
				Patterns: [][]string{
					{"o"},
					{"ha"},
					{"yo"},
					{"wu"},
				},
			},
			expectedCorrect:   "ohayow",
			expectedRemaining: "u",
		},
		{
			name: "joined multi correct input",
			judge: typing.KanaJudge{
				KanaIndex: 3,
				RomaIndex: 1,
				Patterns: [][]string{
					{"o"},
					{"ha"},
					{"yo"},
					{"wu", "ab"},
				},
			},
			expectedCorrect:   "ohayow",
			expectedRemaining: "u",
		},
		{
			name: "multi patterns input",
			judge: typing.KanaJudge{
				KanaIndex: 3,
				RomaIndex: 2,
				Patterns: [][]string{
					{"o"},
					{"ha"},
					{"yo"},
					{"wu"},
				},
			},
			expectedCorrect:   "ohayowu",
			expectedRemaining: "",
		},
		{
			name: "long strings input",
			judge: typing.KanaJudge{
				KanaIndex: 0,
				RomaIndex: 3,
				Patterns: [][]string{
					{"hello"},
				},
			},
			expectedCorrect:   "hel",
			expectedRemaining: "lo",
		},
		{
			name: "remaining kana patterns test",
			judge: typing.KanaJudge{
				KanaIndex: 1,
				RomaIndex: 1,
				Patterns: [][]string{
					{"ab"},
					{"cd"},
					{"ef"},
				},
			},
			expectedCorrect:   "abc",
			expectedRemaining: "def",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			correctColored, remainingColored := tt.judge.ToString()

			if correctColored != tt.expectedCorrect {
				t.Errorf("Expected correctColored to be '%s', but got '%s'", tt.expectedCorrect, correctColored)
			}
			if remainingColored != tt.expectedRemaining {
				t.Errorf("Expected remainingColored to be '%s', but got '%s'", tt.expectedRemaining, remainingColored)
			}
		})
	}
}
