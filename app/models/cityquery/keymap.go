package cityquery

import "github.com/charmbracelet/bubbles/key"

type keymapType struct {
	Accept key.Binding
	Quit   key.Binding
}

func (k keymapType) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Accept, k.Quit,
	}
}

func (k keymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Accept, k.Quit},
	}
}

var keymap keymapType = keymapType{
	Accept: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("ENTER", "Accept"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("Ctrl+C", "Quit"),
	),
}
