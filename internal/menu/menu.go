package menu

import (
	"fmt"
	"gotty/config"
	"gotty/internal/typing"
	"gotty/pkg/display"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func ShowMainMenu() {

	display.ClearTerminal()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   `{{ "❯" | cyan }} {{ . | cyan | underline | bold }}`,
		Inactive: "  {{ . | white }}",
		Selected: `{{ "✔" | green | bold }} {{ . | bold }}`,
	}

	menu := promptui.Select{
		Label:     "Select an option",
		Items:     []string{"Play", "Input Mode", "Sentence", "Exit"},
		Templates: templates,
		HideHelp:  true,
		Size:      4,
	}

	_, result, err := menu.Run()

	if err != nil {
		fmt.Printf("Menu selection failed: %v\n", err)
		return
	}

	switch result {
	case "Play":

		var displayManager typing.DisplayManager

		if config.Config.InputMode == "kana" {
			displayManager = typing.NewKanaDisplayManager()
		} else {
			displayManager = typing.NewRomajiDisplayManager()
		}

		g := typing.Play{
			DisplayManager: displayManager,
			Judge:          typing.NewJudge(config.Config.InputMode),
		}
		g.Start(ShowMainMenu)

	case "Input Mode":
		ShowInputModeSubMenu()

	case "Sentence":
		ShowSentenceSubMenu()

	case "Exit":
		display.ClearTerminal()
		os.Exit(0)
	}
}

func ShowInputModeSubMenu() {
	display.ClearTerminal()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   `{{ "❯" | cyan }} {{ . | cyan | underline | bold }}`,
		Inactive: "  {{ . | white }}",
		Selected: `{{ "✔" | green | bold }} {{ . | bold }}`,
	}

	inputModeOptions := []string{"Romaji", "Kana"}

	menu := promptui.Select{
		Label:     "Select Input Mode",
		Items:     inputModeOptions,
		Templates: templates,
		HideHelp:  true,
		Size:      2,
	}

	_, result, err := menu.Run()

	if err != nil {
		fmt.Printf("Input Mode selection failed: %v\n", err)
		return
	}

	switch result {
	case "Romaji":
		config.Config.InputMode = "romaji"
	case "Kana":
		config.Config.InputMode = "kana"
	}

	if err := config.SaveConfig(); err != nil {
		fmt.Printf("Failed to save config: %v\n", err)
	}

	ShowMainMenu()
}

func ShowSentenceSubMenu() {
	display.ClearTerminal()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   `{{ "❯" | cyan }} {{ . | cyan | underline | bold }}`,
		Inactive: "  {{ . | white }}",
		Selected: `{{ "✔" | green | bold }} {{ . | bold }}`,
	}

	subMenu := promptui.Select{
		Label:     "Configure Sentence Settings",
		Items:     []string{"Set Number of Sentences", "Back to Main Menu"},
		Templates: templates,
		HideHelp:  true,
		Size:      2,
	}

	_, result, err := subMenu.Run()

	if err != nil {
		fmt.Printf("Menu selection failed: %v\n", err)
		return
	}

	switch result {
	case "Set Number of Sentences":
		ShowSentenceConfig()
	case "Back to Main Menu":
		ShowMainMenu()
	}
}

func ShowSentenceConfig() {
	prompt := promptui.Prompt{
		Label:    "Enter the number of sentences",
		Validate: validateNumber,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Failed to get input: %v\n", err)
		return
	}

	number, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Invalid number: %v\n", err)
		return
	}

	config.Config.NumberOfSentences = number
	err = config.SaveConfig()
	if err != nil {
		fmt.Printf("Failed to save config: %v\n", err)
		return
	}

	fmt.Printf("Number of sentences set to: %d\n", number)
	ShowSentenceSubMenu()
}

func validateNumber(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("Invalid number")
	}
	return nil
}
