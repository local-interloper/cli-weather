package root

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/models/citypicker"
	"github.com/local-interloper/cli-weather/app/models/cityquery"
	"github.com/local-interloper/cli-weather/app/models/forecast"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/utils"
)

type Model struct {
	current tea.Model
	style   lipgloss.Style
}

func New() Model {
	width, height := utils.TermSize()
	style := lipgloss.NewStyle().
		Width(width).
		Height(height)

	city, err := utils.GetDefaultCity()
	if err == nil {
		return Model{
			current: forecast.New(city),
			style:   style,
		}
	}

	return Model{
		current: cityquery.New(),
		style:   style,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.current.Init())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.GoToCityQueryMsg:
		m.current = cityquery.New()
		return m, m.current.Init()
	case msgs.GoToCityPickerMsg:
		m.current = citypicker.New(msg.Query)
		return m, m.current.Init()
	case msgs.GoToForecastMsg:
		m.current = forecast.New(msg.City)
		return m, m.current.Init()
	}

	var cmd tea.Cmd
	m.current, cmd = m.current.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	return m.style.Render(m.current.View())
}
