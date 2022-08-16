package pawnui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func New(windowSize tea.WindowSizeMsg) model {
	return model{}
}

type model struct {
	windowSize tea.WindowSizeMsg
}

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
	var s string
	s += "Chose PAWN \n"
	s += "Press q to exit"
	return s
}
