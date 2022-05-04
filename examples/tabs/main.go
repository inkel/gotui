package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/inkel/gotui/tabs"
)

type model struct {
	tabs     tabs.Model
	selected string
}

func initialModel() tea.Model {
	m := model{
		tabs: tabs.New("lorem", "ipsum", "dolor sit", "foo\nbar"),
	}
	return m
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tabs.TabSelectedMsg:
		m.selected = string(msg)
	}

	var cmd tea.Cmd
	m.tabs, cmd = m.tabs.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.tabs.View(), "\n\n", fmt.Sprintf("You have selected %q\n", m.selected))
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "error running example:", err)
		os.Exit(1)
	}
}
