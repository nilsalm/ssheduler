package scheduleui

import (
	"ssheduler/queen"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	qfs        *queen.FileSystem
	execOut    string
	list       list.Model
	windowSize tea.WindowSizeMsg
	spinner    spinner.Model
	uploading  bool
}
type SelectMsg struct {
	Err     error
	CmdPath string
}
type BackMsg bool

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func returnBackCmd() tea.Cmd {
	return func() tea.Msg {
		return BackMsg(true)
	}
}
func New(windowSize tea.WindowSizeMsg) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	// Build up a list of the available commands
	availableCmds := queen.BrowseCommands()
	items := make([]list.Item, len(availableCmds))
	for i, path := range availableCmds {
		items[i] = list.Item(item{title: path, desc: "more ta"})
	}

	m := model{
		qfs:        &queen.FileSystem{Files: queen.GetFS()},
		execOut:    "",
		list:       list.New(items, list.NewDefaultDelegate(), 0, 0),
		windowSize: windowSize,
		spinner:    s,
		uploading:  false,
	}
	m.list.SetSize(m.windowSize.Width, m.windowSize.Height)
	m.list.Title = "Choose a command to schedule"
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowSize = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

		switch msg.Type {
		case tea.KeyBackspace:
			return m, returnBackCmd()

		case tea.KeyEnter:
			m.uploading = true
			items := m.list.Items()
			activeItem := items[m.list.Index()]
			path := activeItem.FilterValue()

			return m, tea.Batch(spinner.Tick, m.sendFile(path, path)) // batch tick and upload together
		}
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}
	if m.uploading == true {
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m model) sendFile(from string, to string) tea.Cmd {
	return func() tea.Msg {
		loc, err := m.qfs.UploadFileToCharm(from, to)
		if err != nil {
			return SelectMsg{Err: err}
		}
		return SelectMsg{CmdPath: loc}
	}
}

func (m model) View() string {
	var s string
	s = docStyle.Render(m.list.View())
	if m.uploading == true {
		s += m.spinner.View() + " uploading ... please wait."
	}
	return s
}
