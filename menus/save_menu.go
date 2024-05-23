package menus

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/savereader"
	tea "github.com/charmbracelet/bubbletea"
)

type modelSaveMenu struct {
	selectedFile      string
	save              savereader.Save
	selectedCode      string
	choices           []choice
	mainMenuChoices   []choice
	exportMenuChoices []choice
	cursor            int
	errorStr          string
}

func (m modelSaveMenu) Init() tea.Cmd {
	return nil
}

func (m *modelSaveMenu) SetVariables() {
	m.mainMenuChoices = []choice{{name: "Export pokemon", code: "export_pokemon"}, {name: "Go back", code: "go_back"}}
	m.choices = m.mainMenuChoices
	m.selectedCode = "general_info"
	m.cursor = 0
}

type choice struct {
	name string
	code string
}

func (m modelSaveMenu) Update(msg tea.Msg) (modelSaveMenu, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if m.selectedCode == "error" {
			m.selectedCode = "general_info"
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
			if m.selectedCode == "export_pokemon" {
				dir, err := os.Getwd()
				if err != nil {
					m.selectedCode = "error"
					m.errorStr = err.Error()
					return m, nil
				}
				dir += "/"
				team := m.save.Trainer.Team()
				err = team[m.cursor].ExportPokemonToFile(dir)
				if err != nil {
					m.selectedCode = "error"
					m.errorStr = err.Error()
					return m, nil
				}
			}

			m.selectedCode = m.choices[m.cursor].code
			m.changeChoices()
		}
	}
	return m, nil
}

func (m *modelSaveMenu) readSave() {
	m.save, _ = savereader.ReadDataFromSave(m.selectedFile)
	m.exportMenuChoices = []choice{}
	for id, pkmn := range m.save.Trainer.Team() {
		if pkmn.OTName() != "" { // improve empy validation
			if pkmn.Nickname() != "" {
				m.exportMenuChoices = append(m.exportMenuChoices, choice{pkmn.Nickname(), string(rune(id))})
			} else {
				m.exportMenuChoices = append(m.exportMenuChoices, choice{pkmn.Species(), string(rune(id))})
			}
		}
	}
}

func (m modelSaveMenu) View() string {
	var s strings.Builder
	switch m.selectedCode {
	case "error":
		s.WriteString("An error has occurred:\n")
		s.WriteString(m.errorStr + "\n")
		s.WriteString("Press any key to continue.\n")
		return s.String()
	case "general_info":
		s.WriteString(m.generalInfo())
	}
	s.WriteString(m.generalInfoMenu())

	return s.String()
}

func (m modelSaveMenu) generalInfo() string {
	var s strings.Builder
	s.WriteString("Game: ")
	s.WriteString(m.save.Game())
	s.WriteString("\n\nTrainer info:")
	s.WriteString("\n\tName: ")
	s.WriteString(m.save.Trainer.Name())
	s.WriteString("\n\tGender: ")
	s.WriteString(m.save.Trainer.Gender())
	s.WriteString("\n\nTeam info:")
	for _, pkmn := range m.save.Trainer.Team() {
		if pkmn.SpeciesIndex() != 0 {
			s.WriteString("\n\t")
			if pkmn.Nickname() != "" {
				s.WriteString(pkmn.Nickname())
			} else {
				s.WriteString(pkmn.Species())
			}
			s.WriteString("\n\tLvl: ")
			s.WriteString(strconv.Itoa(pkmn.Level()))
		}
	}

	s.WriteString("\n")

	return s.String()
}

func (m modelSaveMenu) generalInfoMenu() string {
	s := "\n"
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

func (m *modelSaveMenu) changeChoices() {
	switch m.selectedCode {
	case "export_pokemon":
		m.choices = m.exportMenuChoices
	default:
		m.choices = m.mainMenuChoices
	}
}
