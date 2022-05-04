package tabs

import "github.com/charmbracelet/lipgloss"

var (
	DefaultTabStyle = lipgloss.NewStyle().
			Faint(true).
			Padding(0, 2).
			BorderStyle(lipgloss.NormalBorder())

	DefaultActiveTabStyle = DefaultTabStyle.Copy().
				Faint(false).
				BorderStyle(lipgloss.RoundedBorder())
)

type Styles struct {
	Normal, Active lipgloss.Style
}
