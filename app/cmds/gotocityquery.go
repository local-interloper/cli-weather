package cmds

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/msgs"
)

func GoToCityQuery() tea.Cmd {
	return func() tea.Msg {
		return msgs.GoToCityQueryMsg{}
	}
}
