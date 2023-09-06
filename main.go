package main

import (
	"fmt"
	"os"
	"timekeeper/tab"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	items := []list.Item{
		MenuItem{"ToDo", "Tickets to be completed"},
		MenuItem{"Alert", "Remainder system"},
		MenuItem{"Monitoring", "Monitor screentime"},
		MenuItem{"Time log", "Notes & Secrets"},
	}

	tabs := []string{"Lip Gloss", "Blush", "Eye Shadow", "Mascara", "Foundation"}
	tabContent := []string{"Lip Gloss Tab", "Blush Tab", "Eye Shadow Tab", "Mascara Tab", "Foundation Tab"}
	tabBar := tab.Tabmodel{
		Tabs: tabs, TabContent: tabContent,
	}

	m := MainModel{List: list.New(items, itemDelegateMenuOpt{}, 15, 12), Tabs: tabBar}
	m.List.SetShowStatusBar(false)
	m.List.SetFilteringEnabled(false)
	m.List.SetShowHelp(false)
	m.List.SetShowTitle(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
