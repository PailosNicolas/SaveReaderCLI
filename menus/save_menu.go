package menus

import (
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

func (m modelSaveMenu) View() string {
	return ""
}
