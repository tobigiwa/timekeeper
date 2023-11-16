package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/ts"
)

func main() {

	settingsAndUpdate := lipgloss.JoinHorizontal(lipgloss.Bottom,
		settingsBox.Render("Settings"), updateBox.Render("Updates"))

	// combinedBox :=

	sideBarOptions := lipgloss.JoinVertical(
		lipgloss.Top,
		activeTab.Render("Remainder System"),
		tab.Render("Timer/Alert ⏳"),
		tab.Render("Monitoring"),
		tab.Render("Notes & Secrets"),
		settingsAndUpdate,
	)

	row2 := lipgloss.JoinHorizontal(lipgloss.Top, sidebar.Render(sideBarOptions),
		viewspace.Render())

	row := lipgloss.JoinHorizontal(lipgloss.Top, firsMenubar.Render(),
		secondMenurbar.Render("2"), thirdMenubar.Render("3"),
	)

	allView := lipgloss.JoinVertical(lipgloss.Top, row, row2)

	dialog := lipgloss.Place(
		terminalDim.Col(), terminalDim.Row(), lipgloss.Center,
		lipgloss.Top, allView,
		lipgloss.WithWhitespaceChars("猫咪 猫咪"),
		lipgloss.WithWhitespaceForeground(subtle))

	fmt.Println(dialog)
	fmt.Println()
}

var (
	h           = lipgloss.NewStyle()
	terminalDim = getwindowSize()
	subtle      = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

	firsMenubar = lipgloss.NewStyle().
			Width(int(float32(terminalDim.Col())*0.29)).
			Height(3).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.DoubleBorder()).MarginLeft(1).MarginRight(1)

	secondMenurbar = firsMenubar
	thirdMenubar   = firsMenubar
	sidebar        = lipgloss.NewStyle().
			Width(int(float32(terminalDim.Col()) * 0.29)).
			Height(int(float32(terminalDim.Row()) * 0.75)).
			Border(lipgloss.DoubleBorder()).MarginLeft(1).MarginRight(1)
	viewspace = lipgloss.NewStyle().
			Width(int(float32(terminalDim.Col()) * 0.61)).
			Height(int(float32(terminalDim.Row()) * 0.75)).
			Border(lipgloss.DoubleBorder()).MarginLeft(1).MarginRight(1)

	activeTabBorder = lipgloss.Border{
		Top:         "-",
		Bottom:      "-",
		Left:        "►",
		Right:       "│",
		TopRight:    "◝",
		BottomRight: "◞",
	}

	tabBorder = lipgloss.Border{
		Top:    "─",
		Bottom: "─",
		Left:   " ",
		Right:  "│",
		// TopLeft:     "╭",
		TopRight: "╮",
		// BottomLeft:  "╰",
		BottomRight: "╯",
	}
	settingsAndUpdateBorder = lipgloss.Border{
		Top:    "─",
		Bottom: "─",
		// Left:        " ",
		TopLeft:     "◜",
		MiddleLeft:  "╻",
		BottomLeft:  "◟",
		Right:       "│",
		TopRight:    "◝",
		BottomRight: "◞",
	}

	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		MarginTop(1).
		MarginBottom(1).
		BorderForeground(highlight).
		Padding(0, 1).Faint(true)

	activeTab = tab.Copy().Border(activeTabBorder, true).Blink(true).Faint(false).Underline(true)

	settingsAndUpdateTab = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(highlight).Margin(1)
	settingsBox = settingsAndUpdateTab.Copy()
	updateBox   = settingsAndUpdateTab.Copy()
)

func GetAsciiArt(n int) string {

	if n > 1 {
		os.Exit(1)
	}

	var (
		colorArr  = [7]string{"\u001b[33m", "\u001b[33m", "\u001b[33m", "\u001b[31m", "\u001b[35m", "\u001b[35m", "\u001b[33m"}
		turnoff   = "\u001b[0m\n"
		ch        = make(chan string)
		fileNames = [2]string{"asciiArtLogo.txt", "asciiArt.txt"}
		container = make([]string, 0, 2)
		count     = 0
		size      = len(colorArr)
	)

	for _, v := range fileNames {
		fileName := v

		go func(fileName string, ch chan<- string) {
			var (
				buf   strings.Builder
				bytes []byte
				err   error
			)
			source := rand.NewSource(time.Now().Unix())
			r := rand.New(source)

			if bytes, err = os.ReadFile(fileName); err != nil {
				ch <- err.Error()
				return
			}

			strSlice := strings.Split(strings.TrimSpace(string(bytes)), "\n")
			for _, art := range strSlice {
				buf.WriteString(fmt.Sprintf("%s %s %s", colorArr[r.Intn(size)], art, turnoff))
			}
			ch <- buf.String()

		}(fileName, ch)
	}

	for v := range ch {
		count++
		container = append(container /*strings.TrimSpace(v)*/, v)
		if count >= 2 {
			close(ch)
			fmt.Println("channel closed successfully")
			break
		}
	}

	// fmt.Print(container[0])
	// fmt.Print(container[1])

	// return fmt.Sprintf("%s%s", container[0], container[1])
	return container[n]

}

func getwindowSize() ts.Size {
	var (
		size ts.Size
		err  error
	)

	if size, err = ts.GetSize(); err != nil {
		log.Fatal(err)
	}

	return size

}
