package main

import (
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
	C       = MenuBar.Copy().Width(100).Height(2)

	// Width(30).Height(5)
	MenuitemStyle     = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).
				Foreground(lipgloss.Color("170")).
				Bold(true).BorderBottom(true)

	MenuOpt = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
		Width(10).Height(1).Margin(1).Align(lipgloss.Center)
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
