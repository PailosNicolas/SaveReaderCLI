package menus

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices         []choices
	mainMenuChoices []choices
	readSaveChoices []choices
	cursor          int // which choise the cursor is pointing at
	selectedCode    string
	filePicker      filepicker.Model
	saveMenu        modelSaveMenu
}

type choices struct {
	name string
	code string
}

func InitialModel() model {
	fp := filepicker.New()
	fp.AllowedTypes = []string{".sav", ".pkmn", ".SAV"}
	fp.ShowPermissions = false
	fp.CurrentDirectory, _ = os.Getwd()
	return model{
		mainMenuChoices: []choices{{name: "Read save", code: "read_save"}, {name: "Load pokemon", code: "load_pokemon"}},
		readSaveChoices: []choices{{name: "Read file", code: "read_file"}, {name: "Go to main menu", code: "main_menu"}},
		choices:         []choices{{name: "Read save", code: "read_save"}, {name: "Load pokemon", code: "load_pokemon"}},
		filePicker:      fp,
		selectedCode:    "main_menu",
	}
}

func (m model) Init() tea.Cmd {
	return m.filePicker.Init()
}

func (m model) View() string {
	// The header
	s := "Pokemon save reader CLI:\n\n"

	switch m.selectedCode {
	case "main_menu":
		for i, choice := range m.choices {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, choice.name)
		}

		s += "\nPress q to quit.\n"

	case "read_save":
		for i, choice := range m.choices {

			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, choice.name)
		}

		s += "\nPress q to quit.\n"

	case "read_file":
		s += m.filePicker.View()
		s += "\nPress q to quit.\n"

	case "save_menu":
		s += m.saveMenu.View()
		s += "\nPress q to quit.\n"

	}

	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch m.selectedCode {
		case "save_menu":
			switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

			case tea.KeyEsc.String():
				m.selectedCode = "read_save"
				return m, nil

			}

			m.saveMenu, cmd = m.saveMenu.Update(msg)
			return m, cmd
		case "read_file":
			switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

			case "enter", " ":
				m.filePicker, cmd = m.filePicker.Update(msg)
				if ok, _ := m.filePicker.DidSelectDisabledFile(msg); !ok {
					m.saveMenu.selectedFile = m.filePicker.Path
					m.saveMenu.readSave()
					m.saveMenu.SetVariables()
					m.selectedCode = "save_menu"
				}
				return m, cmd

			case tea.KeyEsc.String():
				m.selectedCode = "read_save"
				return m, nil
			}
			m.filePicker, cmd = m.filePicker.Update(msg)
			return m, cmd
		default:
			switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

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
				switch m.selectedCode {
				case "main_menu":
					m.choices = m.mainMenuChoices
				case "read_save":
					m.choices = m.readSaveChoices
				}
			}
			return m, nil
		}
	}

	m.filePicker, cmd = m.filePicker.Update(msg)

	return m, cmd
}
