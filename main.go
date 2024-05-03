package main

import (
	"fmt"
	"os"

	main_menu "github.com/PailosNicolas/SaveReaderCLI/mainMenu"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(main_menu.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
