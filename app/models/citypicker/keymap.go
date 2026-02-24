package citypicker

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymapType struct {
	Accept key.Binding
	Up     key.Binding
	Down   key.Binding
	Filter key.Binding
	Back   key.Binding
}

func (k keymapType) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Accept, k.Up, k.Down, k.Filter, k.Back,
	}
}

func (s keymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{s.Accept, s.Back},
	}
}

var keymap keymapType = keymapType{
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
