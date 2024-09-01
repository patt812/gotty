package typing

type Judge interface {
	IsCorrect(userInput string) bool
	IsNext() bool
	IsExit(char rune) bool
	ShiftPosition()
	ProcessInput(char rune) string
	ToString() (string, string)
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

func (j *RomajiJudge) IsCorrect(userInput string) bool {
	if len(j.CurrentInput) > len(j.ExpectedInput) || j.ExpectedInput[:len(j.CurrentInput)] != j.CurrentInput {
		return false
	}

	if j.CurrentInput == j.ExpectedInput {
		return true
	}

	return false
}

func (j *RomajiJudge) IsNext() bool {
	return j.CurrentInput == ""
}

func (j *RomajiJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *RomajiJudge) ProcessInput(char rune) string {
	return string(char)
}

func (j *RomajiJudge) ShiftPosition() {
	j.CurrentInput = j.CurrentInput[1:]
}

func (j *RomajiJudge) ToString() (string, string) {
	return j.CurrentInput, j.ExpectedInput
}

type KanaJudge struct {
	KanaIndex int
	RomaIndex int
	Patterns  [][]string
}

func NewKanaJudge(patterns [][]string) *KanaJudge {
	return &KanaJudge{
		KanaIndex: 0,
		RomaIndex: 0,
		Patterns:  patterns,
	}
}

func (j *KanaJudge) IsCorrect(userInput string) bool {
	candidates := j.Patterns[j.KanaIndex]
	correctPatterns := []string{}

	for _, pattern := range candidates {
		if len(pattern) > j.RomaIndex && string(pattern[j.RomaIndex]) == userInput {
			correctPatterns = append(correctPatterns, pattern)
		}
	}

	if len(correctPatterns) > 0 {
		j.Patterns[j.KanaIndex] = correctPatterns
		return true
	}
	return false
}

func (j *KanaJudge) ShiftPosition() {
	for _, pattern := range j.Patterns[j.KanaIndex] {
		if len(pattern) == j.RomaIndex+1 {
			j.RomaIndex = 0
			j.KanaIndex++
			return
		}
	}
	j.RomaIndex++
}

func (j *KanaJudge) IsNext() bool {
	if j.KanaIndex >= len(j.Patterns) {
		j.RomaIndex = 0
		j.KanaIndex++
		return true
	}
	return false
}

func (j *KanaJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *KanaJudge) ProcessInput(char rune) string {
	return string(char)
}

func (j *KanaJudge) ToString() (string, string) {
	var correctColored, remainingColored string

	for i := 0; i < j.KanaIndex; i++ {
		correctColored += j.Patterns[i][0]
	}

	if j.KanaIndex < len(j.Patterns) {
		currentPattern := j.Patterns[j.KanaIndex][0]
		correctColored += currentPattern[:j.RomaIndex]
		remainingColored += currentPattern[j.RomaIndex:]
	}

	for i := j.KanaIndex + 1; i < len(j.Patterns); i++ {
		remainingColored += j.Patterns[i][0]
	}

	return correctColored, remainingColored
}

func NewJudge(inputMode string, patterns [][]string) Judge {
	if inputMode == "kana" {
		return NewKanaJudge(patterns)
	}
	return NewRomajiJudge("")
}
