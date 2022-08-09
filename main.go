package main

import (
	"fmt"

	"os"

	common "ssheduler/common"
	queen "ssheduler/queen"
	"ssheduler/tui/commonui"
	"ssheduler/tui/pawnui"
	"ssheduler/tui/queenui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	fs     *queen.FileSystem
	mode   common.Mode
	queen  tea.Model
	pawn   tea.Model
	common tea.Model
}

// func main() {
// 	cmd_filepath_synced := "/cmd"
// 	cmd_filepath := "ssheduler_cmd.sh"

// 	initialModel := Model{

// 		fs:   nil,
// 		mode: Init,
// 	}

// 	initialModel.mode = Queen
// 	initialModel.fs = &queen.FileSystem{Files: queen.GetFS()}

// 	initialModel.fs.UploadFileToCharm(cmd_filepath, cmd_filepath_synced)
// }

func main() {
	common.ReadConfig()
	err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initialModel() model {
	mode := common.GetMode()
	return model{
		mode:   mode,
		queen:  queenui.New(),
		pawn:   pawnui.New(),
		common: commonui.New(),
	}
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
			common.SaveConfig()
			return m, tea.Quit
		}
	case commonui.SelectMsg:
		switch msg.Choice {
		case 0:
			m.mode = common.Queen
			m.queen = queenui.New()
		case 1:
			m.mode = common.Pawn
			m.pawn = pawnui.New()
		}
		common.SetMode(m.mode)
	}

	switch m.mode {
	case common.Init:
		newModel, newCmd := m.common.Update(msg)
		m.common = newModel
		cmd = newCmd
	case common.Queen:
		newModel, newCmd := m.queen.Update(msg)
		m.queen = newModel
		cmd = newCmd
	case common.Pawn:
		newModel, newCmd := m.pawn.Update(msg)
		m.pawn = newModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.mode {
	case common.Init:
		return m.common.View()
	case common.Queen:
		if m.queen != nil {
			return m.queen.View()
		}
	case common.Pawn:
		if m.pawn != nil {
			return m.pawn.View()
		}
	default:
		return "Press q to exit."
	}
	return "Press q to exit."
}
