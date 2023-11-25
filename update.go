package main

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// m.tWidth, m.tHeight = msg.Width, msg.Height
		// viewportWidth := lipgloss.Width(m.menuBar())
		// viewportHeight := lipgloss.Height(C.String())
		// m.views[m.state] = viewport.New(viewportWidth, viewportHeight)
		// m.views[m.state].MouseWheelEnabled = true
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC, tea.KeyCtrlZ, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyUp, tea.KeyDown, tea.KeyLeft, tea.KeyRight:
			m.list, cmd = m.list.Update(msg)
			cmds = append(cmds, cmd)
			m.viewports.GotoTop()
			m.viewports.SetContent(renderTextArea())
			m.viewports, cmd = m.viewports.Update(msg)
			cmds = append(cmds, cmd)

		case tea.KeyRunes:
			switch {
			case key.Matches(msg, keys.Quit):
				cmd = tea.Quit
				cmds = append(cmds, cmd)
			}
		}

	default:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)

		m.viewports, cmd = m.viewports.Update(msg)
		cmds = append(cmds, cmd)

	}

	return m, tea.Batch(cmds...)
}
