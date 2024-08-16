package menu

import (
	"fmt"
	"gotty/internal/game"
	"gotty/pkg/display"

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
		Items:     []string{"Play", "Exit"},
		Templates: templates,
		HideHelp:  true,
		Size:      2,
	}

	_, result, err := menu.Run()

	if err != nil {
		fmt.Printf("Menu selection failed: %v\n", err)
		return
	}

	switch result {
	case "Play":
		game.Start(ShowMainMenu)
	case "Exit":
		return
	}
}
