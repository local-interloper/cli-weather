package cityquery

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/utils"
)

type Model struct {
	queryInput textinput.Model
	help       help.Model
	style      lipgloss.Style
}

func New() Model {
	width, height := utils.TermSize()

	ti := textinput.New()
	ti.Placeholder = "City name"
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 20

	style := lipgloss.NewStyle().Width(width).Height(height - 1)

	return Model{
		queryInput: ti,
		help:       help.New(),
		style:      style,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if key.Matches(msg, keymap.Accept) {
			return m, cmds.GoToCityPicker(m.queryInput.Value())
		}

		if key.Matches(msg, keymap.Quit) {
			return m, tea.Quit
		}
	}

	{
		var cmd tea.Cmd
		if m.queryInput, cmd = m.queryInput.Update(msg); cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	out := ""

	out += m.style.Render(
		"Enter target city name:\n",
		m.queryInput.View(),
	)

	out += m.help.View(keymap)

	return out
}
