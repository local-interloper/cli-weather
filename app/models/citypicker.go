package models

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/local-interloper/cli-weather/app/cmds"
	"github.com/local-interloper/cli-weather/app/msgs"
	"github.com/local-interloper/cli-weather/app/types"
	"github.com/local-interloper/cli-weather/app/utils"
)

type CityPickerModel struct {
	query      string
	citiesList list.Model
	cities     []types.City
	help       help.Model
}

type cityPickerModelKeymapType struct {
	Accept key.Binding
	Up     key.Binding
	Down   key.Binding
	Filter key.Binding
	Back   key.Binding
}

var cityPickerModelKeymap cityPickerModelKeymapType = cityPickerModelKeymapType{
	Accept: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("ENTER", "Accept"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("Up", "Move selection down"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("Down", "Move selection up"),
	),
	Filter: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "Filter"),
	),
	Back: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("Ctrl+C", "Back"),
	),
}

func (k cityPickerModelKeymapType) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Accept, k.Up, k.Down, k.Filter, k.Back,
	}
}

func (s cityPickerModelKeymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{s.Accept, s.Back},
	}
}

func MakeCityPickerModel(query string) CityPickerModel {
	width, height := utils.TermSize()

	ls := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height-1)
	ls.Title = "Select a city"
	ls.InfiniteScrolling = true
	ls.SetShowHelp(false)

	return CityPickerModel{
		query:      query,
		citiesList: ls,
		help:       help.New(),
	}
}

func (m CityPickerModel) Init() tea.Cmd {
	return tea.Batch(cmds.QueryCities(m.query))
}

func (m CityPickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgs.CitiesQueryResultMsg:
		m.cities = msg.Cities
		items := utils.CitiesToItems(m.cities)
		m.citiesList.SetItems(items)
		m.citiesList.FilterInput.Focus()
		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, setupModelKeymap.Accept) {
			index := m.citiesList.Index()
			city := m.cities[index]
			utils.SetDefaultCity(city)
			return m, cmds.SwitchModel(MakeForecastModel(city))
		}

		if key.Matches(msg, forecastModelKeymap.Quit) {
			return m, cmds.SwitchModel(MakeSetupModel())
		}
	}

	{
		var cmd tea.Cmd
		if m.citiesList, cmd = m.citiesList.Update(msg); cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m CityPickerModel) View() string {
	out := ""
	out += m.citiesList.View()
	out += "\n"
	out += m.help.View(cityPickerModelKeymap)
	return out
}
