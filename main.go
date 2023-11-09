package main

import (
	"log"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sessionState uint

const (
	defaultTime              = time.Minute
	timerView   sessionState = iota
	spinnerView
)

var (
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}

	modelStyle = lipgloss.NewStyle().
			Width(15).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.DoubleBorder())

	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))

	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state   sessionState
	timer   timer.Model
	spinner spinner.Model
	index   int
}

func newModel(timeout time.Duration) mainModel {
	m := mainModel{state: timerView}
	m.timer = timer.New(timeout)
	m.spinner = spinner.New()
	return m
}

func (m mainModel) Init() tea.Cmd {
	return tea.Batch(m.timer.Init(), m.spinner.Tick)
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "q":
			return m, tea.Quit

		}
	}
	return m, nil

}

func (m mainModel) View() string {
	var s string
	// model := m.currentFocusedModel()
	// if m.state == timerView {
	// s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), modelStyle.Render(m.spinner.View()))
	// } else {
	// 	s += lipgloss.JoinHorizontal(lipgloss.Top, modelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), focusedModelStyle.Render(m.spinner.View()))

	// }
	// s += helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	blockA := "AAA \nAAA \nAAA "
	blockB := "BBB \nBBB \nBBB \nBBB \nBBB "

	// Join 20% from the top
	// s := lipgloss.JoinHorizontal(0.2, blockA, blockB)

	// Join on the top edge
	s = lipgloss.JoinHorizontal(lipgloss.Top, blockA, blockB)
	return s
}

func (m mainModel) currentFocusedModel() string {
	if m.state == timerView {
		return "timer"
	}
	return "spinner"
}

func main() {
	p := tea.NewProgram(newModel(defaultTime))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
