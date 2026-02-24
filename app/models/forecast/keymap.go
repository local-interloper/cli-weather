package forecast

import "github.com/charmbracelet/bubbles/key"

type keymapType struct {
	Search key.Binding
	Quit   key.Binding
}

func (k keymapType) ShortHelp() []key.Binding {
	return []key.Binding{k.Search, k.Quit}
}

func (k keymapType) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Search, k.Quit}}
}

var keymap keymapType = keymapType{
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "Search"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("Ctrl+C", "Quit"),
	),
}
