package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/utils"
)

type setupModelKeymapType struct {
	Accept key.Binding
	Quit   key.Binding
}

var setupModelKeymap setupModelKeymapType = setupModelKeymapType{
	Accept: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("ENTER", "Accept"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("Ctrl+C", "Quit"),
	),
}

func (k setupModelKeymapType) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Accept, k.Quit,
	}
}

func (k setupModelKeymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Accept, k.Quit},
	}
}

type SetupModel struct {
	queryInput textinput.Model
	help       help.Model
	style      lipgloss.Style
}

func MakeSetupModel() SetupModel {
	width, height := utils.TermSize()

	ti := textinput.New()
	ti.Placeholder = "City name"
	ti.Focus()
	ti.CharLimit = 32
	ti.Width = 20

	style := lipgloss.NewStyle().Width(width).Height(height - 1)

	return SetupModel{
		queryInput: ti,
		help:       help.New(),
		style:      style,
	}
}

func (m SetupModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m SetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if key.Matches(msg, setupModelKeymap.Accept) {
			return m, cmds.SwitchModel(MakeCityPickerModel(m.queryInput.Value()))

		}

		if key.Matches(msg, forecastModelKeymap.Quit) {
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

func (m SetupModel) View() string {
	out := ""

	out += m.style.Render(
		"Enter target city name:\n",
		m.queryInput.View(),
	)

	out += m.help.View(setupModelKeymap)

	return out
}
