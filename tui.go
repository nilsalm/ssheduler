package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	p *tea.Program
)

type applicationState int

const (
	startState applicationState = iota
	adminState
	remoteState
)

type viewState int

const (
	listView viewState = iota
	startView
	scheduleCommandView
	linkNewRemoteView
	registerToAdminView
	executeCommandView
	changeModeView
)

// func main() {
// 	err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

// func initialModel() model {
// 	return model{}
// }

type model struct {
	state      applicationState
	viewState  viewState
	list       tea.Model // ?
	action     tea.Model // ?
	pages      []string
	activePage string
	cursor     int
	windowSize tea.WindowSizeMsg
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Press Ctrl+C to exit"
}
