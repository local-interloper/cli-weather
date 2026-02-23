package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/utils"
)

type RootModel struct {
	current tea.Model
	style   lipgloss.Style
}

func MakeRootModel() RootModel {
	width, height := utils.TermSize()
	style := lipgloss.NewStyle().
		Width(width).
		Height(height)

	city, err := utils.GetDefaultCity()
	if err == nil {
		return RootModel{
			current: MakeForecastModel(city),
			style:   style,
		}
	}

	return RootModel{
		current: MakeSetupModel(),
		style:   style,
	}
}

func (m RootModel) Init() tea.Cmd {
	return tea.Batch(m.current.Init())
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(msgs.SwitchModelMsg); ok {
		m.current = msg.Next
		return m, m.current.Init()
	}

	var cmd tea.Cmd
	m.current, cmd = m.current.Update(msg)

	return m, cmd
}

func (m RootModel) View() string {
	return m.style.Render(m.current.View())
}
