package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	// activeTabBorder = lipgloss.Border{
	// 	Top:         "─",
	// 	Bottom:      " ",
	// 	Left:        "│",
	// 	Right:       "│",
	// 	TopLeft:     "╭",
	// 	TopRight:    "╮",
	// 	BottomLeft:  "┘",
	// 	BottomRight: "└",
	// }

	// tabBorder = lipgloss.Border{
	// 	Top:         "─",
	// 	Bottom:      "─",
	// 	Left:        "│",
	// 	Right:       "│",
	// 	TopLeft:     "╭",
	// 	TopRight:    "╮",
	// 	BottomLeft:  "┴",
	// 	BottomRight: "┴",
	// }

	width = 96

	terminalDim = getwindowSize()
	logoBanner  = lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder()).
			Width(int(float32(terminalDim.Col()) * 1 / 4)).Height(int(float32(terminalDim.Row()) * 1 / 4))
	MenuBar = lipgloss.NewStyle().Border(lipgloss.DoubleBorder())

	// Width(30).Height(5)
	MenuitemStyle     = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).
				Foreground(lipgloss.Color("170")).
				Bold(true).BorderBottom(true)

	MenuOpt = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
		Width(10).Height(1).Margin(1).Align(lipgloss.Center)

	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")

	inactiveTabStyle = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)

	activeTabStyle = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	windowStyle = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
