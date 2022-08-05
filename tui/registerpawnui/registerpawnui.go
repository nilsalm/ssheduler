package registerpawnui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func New() model {
	return model{}
}

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Press q to exit"
}
