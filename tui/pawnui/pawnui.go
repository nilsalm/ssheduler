package pawnui

import (
	"ssheduler/tui/manualui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	windowSize tea.WindowSizeMsg
	list       list.Model
	manual     tea.Model
	link       tea.Model
	autorun    tea.Model
	state      State
}

type State int

const (
	pawn State = iota
	manual
	link
	autorun
)

type item struct {
	title, desc, id string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func New(windowSize tea.WindowSizeMsg) model {
	items := []list.Item{
		item{title: "Manual execution", desc: "Choose a shell-file from the queen and execute it", id: "manual"},
		item{title: "Link to queen", desc: "Add a new pawn to your queen", id: "link"},
		item{title: "Autorun", desc: "Add a new pawn to your queen", id: "autorun"},
	}
	m := model{
		windowSize: windowSize,
		list:       list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
	m.list.SetSize(m.windowSize.Width, m.windowSize.Height)
	m.list.Title = "Choose an option"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyEnter:
			switch m.list.SelectedItem().FilterValue() {
			case "manual":
				m.manual = manualui.New(m.windowSize)
				m.state = manual

			case "link":
				m.state = link
			case "autorun":
				m.state = autorun
			}
		}
	}

	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.state {
	case manual:
		return m.manual.View()
	// case link:
	// 	return m.link.View()
	// case autorun:
	// 	return m.autorun.View()
	default:
		return docStyle.Render(m.list.View())
	}
}
