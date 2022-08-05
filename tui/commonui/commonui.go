package commonui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// func main() {
// 	err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start()
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

func New() model {
	return model{
		choices: []string{"Queen", "Pawn"},
		cursor:  0,
	}
}

type model struct {
	choices []string
	cursor  int
}

func (m model) Init() tea.Cmd {
	return nil
}

type SelectMsg struct {
	Choice int
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			return m, func() tea.Msg {
				return SelectMsg{Choice: m.cursor}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	var c string

	s += "Please choose a mode:\n"
	for i, choice := range m.choices {

		if i == m.cursor {
			c = ">"
		} else {
			c = " "
		}
		s += fmt.Sprintf("%s %s\n", c, choice)
	}
	s += "Press q to exit"
	return s
}
