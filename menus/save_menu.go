package menus

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/savereader"
	tea "github.com/charmbracelet/bubbletea"
)

type modelSaveMenu struct {
	selectedFile           string
	save                   savereader.Save
	selectedCode           string
	choices                []choice
	mainMenuChoices        []choice
	generalInfoChoices     []choice
	exportMenuChoices      []choice
	teamDetailsMenuChoices []choice
	cursor                 int
	errorStr               string
}

func (m modelSaveMenu) Init() tea.Cmd {
	return nil
}

func (m *modelSaveMenu) SetVariables() {
	m.mainMenuChoices = []choice{{name: "General information", code: "general_info"}, {name: "Export pokemon", code: "export_pokemon"}, {name: "Go back", code: "go_back"}}
	m.generalInfoChoices = []choice{{name: "Team details", code: "team_details"}, {name: "Go back", code: "main_menu"}}
	m.choices = m.mainMenuChoices
	m.selectedCode = "main_menu"
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
				m.teamDetailsMenuChoices = append(m.teamDetailsMenuChoices, choice{pkmn.Nickname(), "team_details"})
			} else {
				m.exportMenuChoices = append(m.exportMenuChoices, choice{pkmn.Species(), string(rune(id))})
				m.teamDetailsMenuChoices = append(m.teamDetailsMenuChoices, choice{pkmn.Species(), "team_details"})
			}
		}
	}
	m.teamDetailsMenuChoices = append(m.teamDetailsMenuChoices, choice{name: "Go back", code: "general_info"})
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
	case "team_details":
		if m.cursor < len(m.teamDetailsMenuChoices)-1 {
			s.WriteString(m.pkmnDetail(m.cursor))
		} else {
			s.WriteString(m.pkmnDetail(m.cursor - 1))
		}
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
			s.WriteString(" Lvl: ")
			s.WriteString(strconv.Itoa(pkmn.Level()))
		}
	}

	s.WriteString("\n")

	return s.String()
}

func (m modelSaveMenu) generalInfoMenu() string {
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

func (m *modelSaveMenu) changeChoices() {
	switch m.selectedCode {
	case "export_pokemon":
		m.choices = m.exportMenuChoices
	case "general_info":
		m.choices = m.generalInfoChoices
	case "team_details":
		m.choices = m.teamDetailsMenuChoices
	default:
		m.choices = m.mainMenuChoices
	}
}

func (m modelSaveMenu) pkmnDetail(id int) string {
	var s strings.Builder
	team := m.save.Trainer.Team()
	if team[id].Nickname() != "" {
		s.WriteString(team[id].Nickname())
	} else {
		s.WriteString(team[id].Species())
	}
	s.WriteString(" Lvl: ")
	s.WriteString(strconv.Itoa(team[id].Level()))
	s.WriteString("\n Item held: ")
	s.WriteString(team[id].ItemHeld().Name)
	s.WriteString("\n Moves:\n\t")
	for index, move := range team[id].Moves() {
		budget := 25
		if move.Id != 0 {
			s.WriteString(move.Name)
		} else {
			s.WriteString("Empty")
		}
		if mod := index % 2; mod != 0 {
			s.WriteString("\n\t")
		} else {
			space, _ := spaceCalculator(budget, move.Name)
			s.WriteString(space)
		}
	}

	budget := 25
	budgetHeader := 10

	s.WriteString("\n Stats:\n\t")
	stats := team[id].Stats()
	ivs := team[id].IVs()
	space, _ := spaceCalculator(budget, " Stats:")
	s.WriteString(space)
	s.WriteString("\t Stat")
	space, _ = spaceCalculator(budgetHeader, "Stat")
	s.WriteString(space)
	s.WriteString("IV")
	space, _ = spaceCalculator(budgetHeader, "IV")
	s.WriteString(space)
	s.WriteString("EVs")

	s.WriteString("\n\tHp:")
	space, _ = spaceCalculator(budget, "Hp:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.TotalHP))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Hp))

	s.WriteString("\n\tAttack:")
	space, _ = spaceCalculator(budget, "Attack:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Attack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Attack))

	s.WriteString("\n\tDefense:")
	space, _ = spaceCalculator(budget, "Defense:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Defense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Defense))

	s.WriteString("\n\tSpecial Defense:")
	space, _ = spaceCalculator(budget, "Special Defense:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.SpecialDefense))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.SpecialDefense))

	s.WriteString("\n\tSpecial Attack:")
	space, _ = spaceCalculator(budget, "Special Attack:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.SpecialAttack))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.SpecialAttack))

	s.WriteString("\n\tSpeed:")
	space, _ = spaceCalculator(budget, "Speed:")
	s.WriteString(space)
	s.WriteString(strconv.Itoa(stats.Speed))
	space, _ = spaceCalculator(budgetHeader, strconv.Itoa(stats.TotalHP))
	s.WriteString(space)
	s.WriteString(strconv.Itoa(ivs.Speed))

	s.WriteString("\n")

	return s.String()
}

func spaceCalculator(budget int, previosWrods ...string) (string, error) {
	s := ""

	for _, word := range previosWrods {
		if len(word) > budget {
			return "", errors.New("previosWrods longer than budget")
		}

		budget -= len(word)
	}

	s += strings.Repeat(" ", budget)
	return s, nil

}
