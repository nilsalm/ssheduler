package queenui

import (
	"ssheduler/tui/registerpawnui"
	"ssheduler/tui/scheduleui"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	schedule   tea.Model
	register   tea.Model
	state      State
	list       list.Model
	windowSize tea.WindowSizeMsg
}

type State int

const (
	choosing State = iota
	schedule
	register
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func New(windowSize tea.WindowSizeMsg) model {
	items := []list.Item{
		item{title: "Schedule Command", desc: "Choose a shell-file and submit it to the pawns"},
		item{title: "Register new Pawn", desc: "Add a new pawn to your queen"},
	}

	m := model{
		list:       list.New(items, list.NewDefaultDelegate(), 0, 0),
		windowSize: windowSize,
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
	switch msg.(type) {
	case scheduleui.BackMsg:
		m.state = choosing
	case scheduleui.SelectMsg:
		m.state = choosing
	}

	switch m.state {
	case schedule:
		newModel, newCmd := m.schedule.Update(msg)
		m.schedule = newModel
		cmd = newCmd
	case register:
		newModel, newCmd := m.register.Update(msg)
		m.register = newModel
		cmd = newCmd
	case choosing:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.windowSize = msg
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			}
			switch msg.Type {
			case tea.KeyEnter:
				if m.list.Index() == 0 {
					m.schedule = scheduleui.New(m.windowSize)
					m.state = schedule
				} else if m.list.Index() == 1 {
					m.register = registerpawnui.New()
					m.state = register
				}
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
	case register:
		return m.register.View()
	case schedule:
		return m.schedule.View()
	case choosing:
		return docStyle.Render(m.list.View())
	default: // choosing
		return docStyle.Render(m.list.View())
	}
}
