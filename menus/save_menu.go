package menus

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/savereader"
	tea "github.com/charmbracelet/bubbletea"
)

type modelSaveMenu struct {
	selectedFile string
	save         savereader.Save
	selectedCode string
	choices      []choice
	cursor       int
}

func (m modelSaveMenu) Init() tea.Cmd {
	return nil
}

func (m *modelSaveMenu) SetVariables() {
	m.choices = []choice{{name: "Export pokemon", code: "export_pokemon"}, {name: "Go back", code: "go_back"}}
	m.selectedCode = "general_info"
	m.cursor = 0
}

type choice struct {
	name string
	code string
}

func (m modelSaveMenu) Update(msg tea.Msg) (modelSaveMenu, tea.Cmd) {
	return m, nil
}

func (m *modelSaveMenu) readSave() {
	m.save, _ = savereader.ReadDataFromSave(m.selectedFile)
}

func (m modelSaveMenu) View() string {
	var s strings.Builder
	switch m.selectedCode {
	case "general_info":
		s.WriteString(m.generalInfo())
		s.WriteString(m.generalInfoMenu())
	}

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
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice.name)
	}

	s += "\n"

	return s
}
