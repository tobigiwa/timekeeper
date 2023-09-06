package main

import tea "github.com/charmbracelet/bubbletea"

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// var menubarFocus string

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width + 20)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC, tea.KeyCtrlZ, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyRight:
			if m.focus {
				// menuOpt, _ := m.List.SelectedItem().(MenuItem)
				
			}
			return m, nil
		}

	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}
