package typing

type Judge interface {
	IsCorrect(userInput string, romajiPatterns [][]string, index int) bool
	IsExit(char rune) bool
	ProcessInput(char rune, currentInput string, targetText string) string
	GetPatternIndex() int
	GetCharIndex() int
	ShouldGoNext() bool
}

type RomajiJudge struct {
	ExpectedInput string
	CurrentInput  string
}

func NewRomajiJudge(expectedInput string) *RomajiJudge {
	return &RomajiJudge{
		ExpectedInput: expectedInput,
		CurrentInput:  "",
	}
}

func (j *RomajiJudge) IsCorrect(userInput string, _ [][]string, _ int) bool {

	if len(j.CurrentInput) > len(j.ExpectedInput) || j.ExpectedInput[:len(j.CurrentInput)] != j.CurrentInput {
		return false
	}

	if j.CurrentInput == j.ExpectedInput {
		return true
	}

	return false
}

func (j *RomajiJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *RomajiJudge) ProcessInput(char rune, currentInput string, targetText string) string {
	if char == 127 {
		if len(j.CurrentInput) > 0 {
			j.CurrentInput = j.CurrentInput[:len(j.CurrentInput)-1]
		}
		return j.CurrentInput
	} else if len(j.CurrentInput) < len(targetText) {
		j.CurrentInput += string(char)
	}
	return j.CurrentInput
}

func (j *RomajiJudge) GetPatternIndex() int {
	return len(j.CurrentInput)
}

func (j *RomajiJudge) GetCharIndex() int {
	return len(j.CurrentInput)
}

func (j *RomajiJudge) ShouldGoNext() bool {
	return j.CurrentInput == j.ExpectedInput
}

type KanaJudge struct {
	PatternIndex    int
	CharIndex       int
	CorrectPatterns []string
	Patterns        [][]string
	GoNext          bool
}

func NewKanaJudge(patterns [][]string) *KanaJudge {
	return &KanaJudge{
		PatternIndex:    0,
		CharIndex:       0,
		CorrectPatterns: nil,
		Patterns:        patterns,
		GoNext:          false,
	}
}

func (j *KanaJudge) IsCorrect(userInput string, romajiPatterns [][]string, index int) bool {

	if index >= len(userInput) || j.PatternIndex >= len(romajiPatterns) {
		return false
	}

	candidates := romajiPatterns[j.PatternIndex]
	newCorrectPatterns := []string{}

	for _, pattern := range candidates {
		if len(pattern) > j.CharIndex && pattern[j.CharIndex] == userInput[index] {
			newCorrectPatterns = append(newCorrectPatterns, pattern)
		}
	}

	if len(newCorrectPatterns) == 0 {
		return false
	}

	j.CorrectPatterns = newCorrectPatterns

	if len(j.CorrectPatterns[0]) == j.CharIndex+1 {
		j.PatternIndex++
		j.CharIndex = 0
		j.CorrectPatterns = nil
		if j.PatternIndex >= len(romajiPatterns) {
			j.GoNext = true
		}
	} else {
		j.CharIndex++
		j.GoNext = false
	}

	return true
}

func (j *KanaJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *KanaJudge) ProcessInput(char rune, currentInput string, targetText string) string {
	if char == 127 {
		if len(currentInput) > 0 {
			currentInput = currentInput[:len(currentInput)-1]
			j.CharIndex = max(0, j.CharIndex-1)
		}
		return currentInput
	}

	if len(currentInput) < len(targetText) {
		currentInput += string(char)
	}

	return currentInput
}

func (j *KanaJudge) GetPatternIndex() int {
	return j.PatternIndex
}

func (j *KanaJudge) GetCharIndex() int {
	return j.CharIndex
}

func (j *KanaJudge) ShouldGoNext() bool {
	return j.GoNext
}

func NewJudge(inputMode string, patterns [][]string) Judge {
	if inputMode == "kana" {
		return NewKanaJudge(patterns)
	}
	return NewRomajiJudge("")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
