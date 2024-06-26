package menus

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/pokemon"
	tea "github.com/charmbracelet/bubbletea"
)

type modelPokemonMenu struct {
	selectedFile     string
	pokemon          pokemon.Pokemon
	selectedCode     string
	choices          []choices
	mainMenuChoices  []choices
	statMenuChoices  []choices
	movesMenuChoices []choices
	cursor           int
	errorStr         string
}

func (m modelPokemonMenu) Init() tea.Cmd {
	return nil
}

func (m *modelPokemonMenu) SetVariables() {
	m.mainMenuChoices = []choices{{name: "Stats", code: "stats_info"}, {name: "Moves", code: "moves_info"}, {name: "Go back", code: "go_back"}}
	m.statMenuChoices = []choices{{name: "Go back", code: "main_menu"}}
	m.choices = m.mainMenuChoices
	m.selectedCode = "main_menu"
	m.cursor = 0
}

func (m *modelPokemonMenu) readPokemon() {
	m.pokemon, _ = pokemon.ReadPokemonFromFile(m.selectedFile)
	for _, move := range m.pokemon.Moves() {
		m.movesMenuChoices = append(m.movesMenuChoices, choices{name: move.Name, code: "moves_info"})
	}
	m.movesMenuChoices = append(m.movesMenuChoices, choices{name: "Go back", code: "main_menu"})
}

func (m modelPokemonMenu) View() string {
	var s strings.Builder

	switch m.selectedCode {
	case "main_menu":
		s.WriteString(m.mainMenuView())
	case "stats_info":
		s.WriteString(pokemonStatView(m.pokemon))
	case "moves_info":
		moves := m.pokemon.Moves()
		if m.cursor < len(m.movesMenuChoices)-1 {
			s.WriteString(moveView(moves[m.cursor]))
		} else {
			s.WriteString(moveView(moves[m.cursor-1]))
		}
		s.WriteString("\n")
	case "error":
		s.WriteString("An error has occurred:\n")
		s.WriteString(m.errorStr + "\n")
		s.WriteString("Press any key to continue.\n")
		return s.String()
		//case "general_info":
		//	s.WriteString(m.generalInfo())
	}

	s.WriteString(m.generalInfoMenu())

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
		case "enter", " ":
			m.selectedCode = m.choices[m.cursor].code
			m.changeChoices()
			m.cursor = 0
		}

	}
	return m, nil
}

func (m modelPokemonMenu) generalInfoMenu() string {
	s := ""
	for i, choice := range m.choices {
		if m.cursor == i {
			s += fmt.Sprintf("\033[31m%s \033[0m\n", choice.name)
		} else {
			s += fmt.Sprintf("%s\n", choice.name)
		}
	}

	s += "\n"

	return s
}

func (m *modelPokemonMenu) changeChoices() {
	switch m.selectedCode {
	case "general_info":
		m.choices = m.mainMenuChoices //TODO: it should be change to a generalInfoChoices when it is aviable
	case "stats_info":
		m.choices = m.statMenuChoices
	case "moves_info":
		m.choices = m.movesMenuChoices
	default:
		m.choices = m.mainMenuChoices
	}
}

func (m *modelPokemonMenu) mainMenuView() string {
	var s strings.Builder
	budget := 15
	if m.pokemon.Nickname() != "" {
		s.WriteString(m.pokemon.Nickname())
	} else {
		s.WriteString(m.pokemon.Species())
	}
	s.WriteString("\n")
	s.WriteString("Lvl: ")
	space, _ := spaceCalculator(budget, "Lvl: ")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(m.pokemon.Level()))
	s.WriteString("\n")
	s.WriteString("Item held: ")
	space, _ = spaceCalculator(budget, "Item held: ")
	s.WriteString(space)
	s.WriteString(m.pokemon.ItemHeld().Name)
	s.WriteString("\n")
	s.WriteString("Pokeball: ")
	space, _ = spaceCalculator(budget, "Pokeball: ")
	s.WriteString(space)
	s.WriteString(m.pokemon.PokeBall())
	s.WriteString("\n")
	s.WriteString("OT Name: ")
	space, _ = spaceCalculator(budget, "OT Name: ")
	s.WriteString(space)
	s.WriteString(m.pokemon.OTName())
	s.WriteString("\n\n")

	return s.String()
}
