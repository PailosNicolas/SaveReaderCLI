package menus

import (
	"strings"

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

func (m *modelPokemonMenu) readPokemon() {
	m.pokemon, _ = pokemon.ReadPokemonFromFile(m.selectedFile)
}

func (m modelPokemonMenu) View() string {
	var s strings.Builder

	s.WriteString("Pokemon menu")

	return s.String()
}

func (m modelPokemonMenu) Update(msg tea.Msg) (modelPokemonMenu, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if m.selectedCode == "error" {
			m.selectedCode = "main_menu"
			m.choices = m.mainMenuChoices
			return m, nil
		}
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}
