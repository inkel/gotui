package tabs

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Styles Styles
	KeyMap KeyMap
	titles []string
	cur    int
}

func New(titles ...string) Model {
	return Model{
		titles: titles,
		Styles: Styles{
			Normal: DefaultTabStyle,
			Active: DefaultActiveTabStyle,
		},
		KeyMap: DefaultKeyMap(),
	}
}

func (m Model) Init() tea.Cmd { return nil }

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

func (m Model) View() string {
	tabs := make([]string, len(m.titles))

	for i, title := range m.titles {
		style := m.Styles.Normal
		if i == m.cur {
			style = m.Styles.Active
		}
		tabs[i] = style.Render(title)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m Model) Selected() string {
	return m.titles[m.cur]
}

func (m *Model) Next() {
	m.cur = (m.cur + 1) % len(m.titles)
}

func (m *Model) Prev() {
	if m.cur == 0 {
		m.cur = len(m.titles) - 1
	} else {
		m.cur = (m.cur - 1) % len(m.titles)
	}
}

type TabSelectedMsg string

func (m Model) TabSelected() tea.Cmd {
	return func() tea.Msg {
		return TabSelectedMsg(m.Selected())
	}
}
