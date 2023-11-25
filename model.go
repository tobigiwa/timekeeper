package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuItem struct {
	option string
	desc   string
}

func (m MenuItem) Title() string       { return m.option }
func (m MenuItem) Description() string { return m.desc }

func (m MenuItem) FilterValue() string { return m.option }

type itemDelegateMenuOpt struct{}

func (i itemDelegateMenuOpt) Height() int  { return 0 }
func (i itemDelegateMenuOpt) Spacing() int { return 2 }

func (i itemDelegateMenuOpt) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (i itemDelegateMenuOpt) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	var (
		item MenuItem
		ok   bool
	)

	if item, ok = listItem.(MenuItem); !ok {
		return
	}

	fn := MenuitemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(lipgloss.JoinHorizontal(
				lipgloss.Top,
				"|☛ "+strings.Join(s, " "),
			))
		}
	}

	fmt.Fprint(w, fn(item.option))
}

type MainModel struct {
	list        list.Model
	viewports   viewport.Model
	focusOnList bool
	state       status
	focus       bool
	tHeight     int
	tWidth      int
}
