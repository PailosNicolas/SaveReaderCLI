package menus

import (
	"strings"

	"github.com/PailosNicolas/GoPkmSaveReader/savereader"
	tea "github.com/charmbracelet/bubbletea"
)

type modelSaveMenu struct {
	selectedFile string
	save         savereader.Save
}

func (m modelSaveMenu) Init() tea.Cmd {
	return nil
}

func (m modelSaveMenu) Update(msg tea.Msg) (modelSaveMenu, tea.Cmd) {
	return m, nil
}

func (m *modelSaveMenu) readSave() {
	m.save, _ = savereader.ReadDataFromSave(m.selectedFile)
}

func (m modelSaveMenu) View() string {
	var s strings.Builder
	s.WriteString("Game: ")
	s.WriteString(m.save.Game())
	s.WriteString("\nTrainer info:")
	s.WriteString("\n\tName: ")
	s.WriteString(m.save.Trainer.Name())
	s.WriteString("\n\tGender: ")
	s.WriteString(m.save.Trainer.Gender())

	return s.String()
}
