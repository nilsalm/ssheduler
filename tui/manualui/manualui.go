package main

import (
	pawn "ssheduler/pawn"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	list       list.Model
	pfs        *pawn.FileSystem
	windowSize tea.WindowSizeMsg
}
type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func New(windowSize tea.WindowSizeMsg) model {
	m := model{
		windowSize: windowSize,
		pfs:        &pawn.FileSystem{Files: pawn.GetFS()},
	}

	// Build up a list of the available commands
	availableCmds := m.pfs.BrowseCommands()
	items := make([]list.Item, len(availableCmds))
	for i, path := range availableCmds {
		items[i] = list.Item(item{title: path, desc: "more ta"})
	}

	m.list = list.New(items, list.NewDefaultDelegate(), 0, 0)

	m.list.SetSize(m.windowSize.Width, m.windowSize.Height)

	return m
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
	return "Press q to exit"
}
