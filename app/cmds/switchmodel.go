package cmds

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/msgs"
)

func SwitchModel(model tea.Model) tea.Cmd {
	return func() tea.Msg {
		return msgs.SwitchModelMsg{
			Next: model,
		}
	}
}
