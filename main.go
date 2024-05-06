package main

import (
	"fmt"
	"os"

	menus "github.com/PailosNicolas/SaveReaderCLI/menus"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(menus.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
