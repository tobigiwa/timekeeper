package main

import (
	"github.com/charmbracelet/lipgloss"
)

func (m MainModel) View() string {

	MenuBar := lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).
		Width(20).Height(3)
	c := MenuBar.Copy().Width(100).Height(2)
	d := MenuBar.Copy().Height(13).Width(70)

	// doc := strings.Builder{}
	// {
	// 	var renderedTabs []string

	// 	for i, t := range m.Tabs {
	// 		var style lipgloss.Style
	// 		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
	// 		if isActive {
	// 			style = activeTabStyle.Copy()
	// 		} else {
	// 			style = inactiveTabStyle.Copy()
	// 		}
	// 		border, _, _, _, _ := style.GetBorder()
	// 		if isFirst && isActive {
	// 			border.BottomLeft = "│"
	// 		} else if isFirst && !isActive {
	// 			border.BottomLeft = "├"
	// 		} else if isLast && isActive {
	// 			border.BottomRight = "│"
	// 		} else if isLast && !isActive {
	// 			border.BottomRight = "┤"
	// 		}
	// 		style = style.Border(border)
	// 		renderedTabs = append(renderedTabs, style.Render(t))
	// 	}

	// 	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	// 	doc.WriteString(row)
	// 	doc.WriteString("\n")
	// 	doc.WriteString(windowStyle.Width(lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize()).Render(m.TabContent[m.activeTab]))
	// }

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			MenuBar.Render(m.List.View()),
			d.Render(m.Tabs.View()),
		),

		c.String(),
	)
}

type status int

const (
	todo status = iota
	alert
	monitor
	timeLog

)

