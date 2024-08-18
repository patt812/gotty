package typing

type Judge interface {
	IsCorrect(userInput, targetText string, index int) bool
	IsExit(char rune) bool
	ProcessInput(char rune, currentInput, targetText string) string
}

type RomajiJudge struct{}

func NewRomajiJudge() *RomajiJudge {
	return &RomajiJudge{}
}

func (j *RomajiJudge) IsCorrect(userInput, targetText string, index int) bool {
	return userInput[index] == targetText[index]
}

func (j *RomajiJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *RomajiJudge) ProcessInput(char rune, currentInput, targetText string) string {
	if char == 127 {
		if len(currentInput) > 0 {
			return currentInput[:len(currentInput)-1]
		}
		return currentInput
	} else if len(currentInput) < len(targetText) {
		return currentInput + string(char)
	}
	return currentInput
}

type KanaJudge struct{}

func NewKanaJudge() *KanaJudge {
	return &KanaJudge{}
}

func (j *KanaJudge) IsCorrect(userInput, targetText string, index int) bool {

	return userInput[index] == targetText[index]
}

func (j *KanaJudge) IsExit(char rune) bool {
	return char == 27
}

func (j *KanaJudge) ProcessInput(char rune, currentInput, targetText string) string {
	if char == 127 {
		if len(currentInput) > 0 {
			return currentInput[:len(currentInput)-1]
		}
		return currentInput
	} else if len(currentInput) < len(targetText) {
		return currentInput + string(char)
	}
	return currentInput
}

func NewJudge(inputMode string) Judge {
	if inputMode == "kana" {
		return NewKanaJudge()
	}
	return NewRomajiJudge()
}
