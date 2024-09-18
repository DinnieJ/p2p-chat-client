package gui

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	// "github.com/nathan-fiscaletti/consolesize-go"
)

const APP_TITLE = "P2P Chat Client"
const APP_MIN_WIDTH = 80
const APP_MIN_HEIGHT = 24

type Feature struct {
	Name        string
	Description string
	Layout      tview.Primitive
}

type View struct {
	layout          tview.Primitive
	sublayout       *tview.Grid
	peerListSidebar *tview.List
	currentHostInfo *tview.TextView
	rootClient      interface{} // root client library, will be implemented later
}

func (w *View) InitLayout() {
	sidebar := tview.NewList()
	sidebar.SetBackgroundColor(tcell.Color16)
	sidebar.SetSelectedBackgroundColor(tcell.ColorGreen)

	userinfo := tview.NewTextView()
	userinfo.SetBackgroundColor(tcell.Color16)
	userinfo.SetText("Hello world !")

	// Create a Grid layout
	grid := tview.NewGrid()
	grid.SetMinSize(8, 30)
	grid.SetBorders(true)
	grid.SetColumns(25, -1)
	grid.SetRows(24, 0)
	grid.SetBordersColor(tcell.Color102)
	grid.SetBackgroundColor(tcell.Color16)

	sublayout := tview.NewGrid()
	chatLayout := tview.NewGrid()
	chatLayout.SetRows(20, 0)
	chatTextView := tview.NewTextView()
	chatTextView.SetBackgroundColor(tcell.Color16)
	chatTextView.SetBorder(true)
	// chatTextView.SetText("Fuckdnaskdn")
	chatInput := tview.NewTextArea()
	chatInput.SetBorder(true)
	chatInput.SetBackgroundColor(tcell.Color16)
	chatInput.SetBackgroundColor(tcell.ColorBlack)
	chatInput.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		text := chatInput.GetText()
		if event.Modifiers() == tcell.ModShift && event.Key() == tcell.KeyEnter {
			fmt.Println("Helloworld")
			chatInput.SetText(text, true)
			return nil
		} else if event.Key() == tcell.KeyEnter {
			if strings.TrimSpace(text) == "" {
				return event
			}
			chatLayoutText := chatTextView.GetText(true)
			chatTextView.SetText(chatLayoutText + text + "\n")
			chatInput.SetText("", true)
		}
		return event
	})
	// chatInput.SetFieldWidth(30)

	chatLayout.AddItem(chatTextView, 0, 0, 1, 1, 0, 0, false)
	chatLayout.AddItem(chatInput, 1, 0, 1, 1, 8, 0, true)

	sublayout.AddItem(chatLayout, 0, 0, 1, 1, 0, 0, false)
	// grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)  // Header spans 2 columns
	grid.AddItem(sidebar, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(userinfo, 1, 0, 1, 1, 8, 0, false)
	grid.AddItem(sublayout, 0, 1, 2, 1, 0, 0, false)
	// grid.AddItem(content, 1, 1, 1, 1, 0, 0, false) // Content on the right
	// grid.AddItem(footer, 2, 0, 1, 2, 0, 0, false)  // Footer spans 2 columns
	w.layout = grid
	w.sublayout = sublayout
}

func (w *View) SetLayout() {
	w.sublayout.Draw(tcell.NewSimulationScreen(tview.NewLine))
}

func (w *View) StartApp() error {
	w.InitLayout()
	newApp := tview.NewApplication()
	newApp = newApp.SetRoot(w.layout, true)
	newApp = newApp.EnableMouse(true)
	newApp.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		width, height := screen.Size()
		if width < APP_MIN_WIDTH || height < APP_MIN_HEIGHT {
			// fmt.Printf("please set your terminal to larger size. Minimum (%d * %d)\n", width, height)
			message := fmt.Sprintf("Please set your terminal to minimum size of (%d, %d)\nCurrent size: %d, %d", APP_MIN_WIDTH, APP_MIN_HEIGHT, width, height)
			screen.SetContent(0, 0, '\u0700', []rune(message), tcell.StyleDefault)
			return true
		}
		return false
	})
	// w.SetLayout()
	return newApp.Run()
}
