package typing

type Judge struct {
	IsCorrectFunc func(string, string, int) bool
}

func NewJudge() *Judge {
	return &Judge{
		IsCorrectFunc: DefaultIsCorrect,
	}
}

func DefaultIsCorrect(userInput, targetText string, index int) bool {
	return userInput[index] == targetText[index]
}

func (j *Judge) isCorrect(userInput, targetText string, index int) bool {
	return j.IsCorrectFunc(userInput, targetText, index)
}

func (j *Judge) IsExit(char rune) bool {
	// Escape
	return char == 27
}

func (j *Judge) ProcessInput(char rune, currentInput, targetText string) string {
	// Backspace
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
