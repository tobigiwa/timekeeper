package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (m MainModel) Init() tea.Cmd {
	return nil
}

func newModel() MainModel {

	var (
		m MainModel
	)

	items := []list.Item{
		MenuItem{"ToDo", "Tickets to be completed"},
		MenuItem{"Alert", "Remainder system"},
		MenuItem{"Monitoring", "Monitor screentime"},
		MenuItem{"Time log", "Notes & Secrets"},
	}

	m.list = list.New(items, itemDelegateMenuOpt{}, 15, 10)
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)

	m.state = todo
	m.focus = true

	return m
}
