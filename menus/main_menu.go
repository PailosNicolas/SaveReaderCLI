package menus

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices         []choices
	mainMenuChoices []choices // items on the to-do list
	readSaveChoices []choices // items on the to-do list
	cursor          int       // which to-do list item our cursor is pointing at
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
		// Our to-do list is a grocery list
		mainMenuChoices: []choices{{name: "Read save", code: "read_save"}, {name: "Load pokemon", code: "load_pokemon"}},
		readSaveChoices: []choices{{name: "Read file", code: "read_file"}, {name: "Go to main menu", code: "main_menu"}},
		choices:         []choices{{name: "Read save", code: "read_save"}, {name: "Load pokemon", code: "load_pokemon"}},
		filePicker:      fp,
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selectedCode: "main_menu",
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
		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice.name)
		}

		// The footer
		s += "\nPress q to quit.\n"

	case "read_save":
		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice.name)
		}

		// The footer
		s += "\nPress q to quit.\n"

	case "read_file":
		s += m.filePicker.View()

	case "save_menu":
		s += m.saveMenu.View()

	}

	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	// Is it a key press?
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
			// Cool, what was the actual key pressed?
			switch msg.String() {

			// These keys should exit the program.
			case "ctrl+c", "q":
				return m, tea.Quit

			// The "up" and "k" keys move the cursor up
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			// The "down" and "j" keys move the cursor down
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
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

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}
