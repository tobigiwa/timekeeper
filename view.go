package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func (m MainModel) View() string {

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			m.menuBar(),
			m.viewBar(),
		),

		C.String(),
	)

}

func (m MainModel) menuBar() string {
	menuBar := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Width(19)

	return menuBar.Render(m.list.View())
}

func (m MainModel) viewBar() string {

	// currectView := m.views[m.state]
	// currectView.Style = lipgloss.NewStyle().
	// 	Width(80).
	// 	Height(30).Border(lipgloss.DoubleBorder())

	m.viewports = viewport.New(80, 25)
	m.viewports.Style = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Foreground(lipgloss.Color("170"))

	return m.viewports.View()

}

func (m MainModel) viewportContent(width int) string {
	var builder strings.Builder
	for i := 0; i <= width; i++ {
		builder.WriteString(strings.Repeat("---\n", width))
	}
	return wordwrap.String(builder.String(), width)
}

type status int

const (
	todo status = iota
	alert
	monitor
	timeLog
)

func renderTextArea() string {
	x := textarea.New()
	x.SetWidth(40)
	x.Placeholder = "jing yang"
	x.Cursor.Blink = true
	x.Focus()
	return x.View()
}
