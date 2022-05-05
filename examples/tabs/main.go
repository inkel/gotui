package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/inkel/gotui/tabs"
)

type model struct {
	tabs     tabs.Model
	selected tabs.Tab
}

func initialModel() tea.Model {
	m := model{
		tabs: tabs.New(stringTab("lorem"), intTab(1234), tab{title: "ipsum", data: tabData{"foo", true}}),
	}
	return m
}

func (m model) Init() tea.Cmd {
	return m.tabs.TabSelected()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tabs.TabSelectedMsg:
		m.selected = msg
	}

	var cmd tea.Cmd
	m.tabs, cmd = m.tabs.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.selected == nil {
		return "" //panic("wahaaaaaa")
	}
	return lipgloss.JoinVertical(
		lipgloss.Left, m.tabs.View(),
		"\n\n",
		fmt.Sprintf("You have selected tab %q with data of type %[2]T and value %#[2]v\n", m.selected.Title(), m.selected.Data()))
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "error running example:", err)
		os.Exit(1)
	}
}

// custom stringTab for a string
type stringTab string

func (t stringTab) Title() string     { return string(t) }
func (t stringTab) Data() interface{} { return string(t) }

// custom intTab for a storing int
type intTab int

func (t intTab) Title() string     { return strconv.Itoa(int(t)) }
func (t intTab) Data() interface{} { return int(t) }

// custom composite tab
type tabData struct {
	foo string
	bar bool
}

type tab struct {
	title string
	data  tabData
}

func (t tab) Title() string     { return t.title }
func (t tab) Data() interface{} { return t.data }
