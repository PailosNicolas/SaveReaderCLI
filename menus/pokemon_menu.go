package menus

import (
	"github.com/PailosNicolas/GoPkmSaveReader/pokemon"
	tea "github.com/charmbracelet/bubbletea"
)

type modelPokemonMenu struct {
	selectedFile    string
	pokemon         pokemon.Pokemon
	selectedCode    string
	choices         []choices
	mainMenuChoices []choices
	cursor          int
	errorStr        string
}

func (m modelPokemonMenu) Init() tea.Cmd {
	return nil
}

func (m *modelPokemonMenu) SetVariables() {
	m.mainMenuChoices = []choices{{name: "General information", code: "general_info"}, {name: "Go back", code: "go_back"}}
	m.choices = m.mainMenuChoices
	m.selectedCode = "main_menu"
	m.cursor = 0
}
