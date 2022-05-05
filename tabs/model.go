// Package tabs provides a component that draws selectable tabs.
package tabs

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model contains the state for the tabs. Use New to create a new
// models rather than using Model as a struct literal.
type Model struct {
	Styles Styles
	KeyMap KeyMap
	tabs   []Tab
	cur    int
}

// New returns a Model with the given titles as tabs, using default
// styles and keymaps.
func New(tabs ...Tab) Model {
	return Model{
		tabs: tabs,
		Styles: Styles{
			Normal: DefaultTabStyle,
			Active: DefaultActiveTabStyle,
		},
		KeyMap: DefaultKeyMap(),
	}
}

// Update is the Tea update function. It will move the tab selection
// to the next or previous one based on the key passed; if that
// happens a TabSelectedMsg will be broadcasted with the return value
// of Selected.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch {
	case key.Matches(keyMsg, m.KeyMap.Next):
		m.Next()

	case key.Matches(keyMsg, m.KeyMap.Prev):
		m.Prev()
	}

	return m, m.TabSelected()
}

// View renders the models' view.
func (m Model) View() string {
	tabs := make([]string, len(m.tabs))

	for i, t := range m.tabs {
		style := m.Styles.Normal
		if i == m.cur {
			style = m.Styles.Active
		}
		tabs[i] = style.Render(t.Title())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

// Selected returns the currently selected tab.
func (m Model) Selected() Tab {
	return m.tabs[m.cur]
}

// Next selects the next tab, wrapping back to the first one if at the
// last one.
func (m *Model) Next() {
	m.cur = (m.cur + 1) % len(m.tabs)
}

// Prev selectes the previous tab, wrapping forward to the last one if
// at the first one.
func (m *Model) Prev() {
	if m.cur == 0 {
		m.cur = len(m.tabs) - 1
	} else {
		m.cur = (m.cur - 1) % len(m.tabs)
	}
}

// TabSelectedMsg indicates a new tab was selected.
type TabSelectedMsg Tab

// TabSelected is the command used to broadcast the selected tab.
func (m Model) TabSelected() tea.Cmd {
	return func() tea.Msg {
		return TabSelectedMsg(m.Selected())
	}
}

// Tab is the interface used to create tabs.
type Tab interface {
	// Title is the text used to render the tab.
	Title() string

	// Data contains data associated with the tab.
	Data() interface{}
}
