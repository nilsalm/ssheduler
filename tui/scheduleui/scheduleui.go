package scheduleui

import (
	"fmt"
	"ssheduler/queen"

	tea "github.com/charmbracelet/bubbletea"
)

func New() model {
	return model{
		cursor:        0,
		availableCmds: queen.BrowseCommands(),
		qfs:           &queen.FileSystem{Files: queen.GetFS()},
		execOut:       "",
		debug:         "debug",
	}
}

type model struct {
	cursor        int
	availableCmds []string
	qfs           *queen.FileSystem
	execOut       string
	debug         string
}

func (m model) Init() tea.Cmd {
	return nil
}

type SelectMsg struct {
	Choice  int
	CmdPath string
}
type BackMsg bool

func returnBackCmd() tea.Cmd {
	return func() tea.Msg {
		return BackMsg(true)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.availableCmds)-1 {
				m.cursor++
			}
		}

		switch msg.Type {
		case tea.KeyBackspace:
			m.debug = "BACKMSG"
			return m, returnBackCmd()

		case tea.KeyEnter:
			p := m.availableCmds[m.cursor]
			m.execOut = m.qfs.UploadFileToCharm(p, p)

			return m, func() tea.Msg {
				return SelectMsg{Choice: m.cursor, CmdPath: m.availableCmds[m.cursor]}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	var c string

	s += "Please choose a command:\n"
	for i, choice := range m.availableCmds {

		if i == m.cursor {
			c = ">"
		} else {
			c = " "
		}
		s += fmt.Sprintf("%s %s\n", c, choice)
	}
	s += fmt.Sprintf("%s", m.execOut)
	s += fmt.Sprintf("%s", m.debug)
	s += "Press q to exit"
	return s
}
