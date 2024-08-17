package menu

import (
	"fmt"
	"gotty/config"
	"gotty/internal/typing"
	"gotty/pkg/display"
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
		Items:     []string{"Play", "Sentence", "Exit"},
		Templates: templates,
		HideHelp:  true,
		Size:      3,
	}

	_, result, err := menu.Run()

	if err != nil {
		fmt.Printf("Menu selection failed: %v\n", err)
		return
	}

	switch result {
	case "Play":
		g := typing.Play{}
		g.Start(ShowMainMenu)
	case "Sentence":
		ShowSentenceSubMenu()
	case "Exit":
		return
	}
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
