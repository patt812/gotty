package main

import (
	"gotty/config"
	"gotty/internal/menu"
	"log"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	menu.ShowMainMenu()
}
