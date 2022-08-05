package queenui

import (
	"fmt"
	"ssheduler/tui/registerpawnui"
	"ssheduler/tui/scheduleui"

	tea "github.com/charmbracelet/bubbletea"
)

func New() model {
	return model{
		cursor:  0,
		choices: []string{"Schedule Command", "Register Pawn"},
	}
}

type State int

const (
	choosing State = iota
	schedule
	register
)

type model struct {
	cursor   int
	choices  []string
	schedule tea.Model
	register tea.Model
	state    State
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
		case tea.KeyMsg:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			}
			switch msg.Type {
			case tea.KeyEnter:
				if m.cursor == 0 {
					m.schedule = scheduleui.New()
					m.state = schedule
				} else if m.cursor == 1 {
					m.register = registerpawnui.New()
					m.state = register
				}
			}
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)

}

func (m model) View() string {
	var s string
	var c string

	switch m.state {
	case choosing:
		s += "Please choose an option:\n"
		for i, choice := range m.choices {
			if i == m.cursor {
				c = ">"
			} else {
				c = " "
			}
			s += fmt.Sprintf("%s %s\n", c, choice)
		}
	case register:
		return m.register.View()
	case schedule:
		return m.schedule.View()
	}
	s += "Press q to exit"
	return s
}
