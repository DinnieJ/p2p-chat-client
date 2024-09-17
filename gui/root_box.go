package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const APP_TITLE = "P2P Chat Client"

type Feature struct {
	Name        string
	Description string
	Layout      tview.Primitive
}

type View struct {
	layout    tview.Primitive
	sublayout *tview.Flex
}

func (w *View) initBox() {
	sidebar := tview.NewList()
	sidebar.SetBackgroundColor(tcell.Color16)
	sidebar.AddItem("Fucked", "online", '\u070D', func() {})
	sidebar.AddItem("Fucked2", "online", tview.BlockDarkShade, func() {})
	sidebar.AddItem("Fucked3", "online", tview.BlockLeftHalfBlock, func() {})
	sidebar.SetSelectedBackgroundColor(tcell.ColorGreen)

	userinfo := tview.NewTextView()
	userinfo.SetBackgroundColor(tcell.Color16)
	userinfo.SetText("Hello world !")

	// Create a Grid layout
	grid := tview.NewGrid()
	grid.SetBorders(true)
	grid.SetColumns(30, 0)
	grid.SetRows(30, 0)
	grid.SetBordersColor(tcell.Color102)
	grid.SetBackgroundColor(tcell.Color16)

	sublayout := tview.NewFlex().SetDirection(tview.FlexRow)

	// grid.AddItem(header, 0, 0, 1, 2, 0, 0, false)  // Header spans 2 columns
	grid.AddItem(sidebar, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(userinfo, 1, 0, 1, 1, 0, 0, true)
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
	w.initBox()
	newApp := tview.NewApplication()
	newApp = newApp.SetRoot(w.layout, true)
	newApp = newApp.EnableMouse(true)
	// w.SetLayout()
	return newApp.Run()
}
