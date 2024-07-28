package ui

import (
	"dead_modules/search"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const version = "1.0.0"

var app *tview.Application
var debugTextView *tview.TextView
var debugEnabled = false

func StartApp() error {
	app = tview.NewApplication()

	textView := tview.NewTextView().
		SetText("[yellow]Dead Modules v" + version + "[white]\n\nPress [green]Enter[white] to start searching for node_modules...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	debugTextView = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			showLoading()
			go search.SearchOldModules("/Users/furkan/Developer", app, updateTable, logDebug)
		}
	})

	// Center the textView horizontally and vertically using a Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).           // Add a spacer at the top
		AddItem(textView, 0, 1, true).       // Add the textView
		AddItem(debugTextView, 10, 1, false) // Add debug view at the bottom

	return app.SetRoot(flex, true).Run()
}

func showLoading() {
	textView := tview.NewTextView().
		SetText("[yellow]Searching for old node_modules...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Center the textView horizontally and vertically using a Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).     // Add a spacer at the top
		AddItem(textView, 0, 1, true). // Add the textView
		AddItem(debugTextView, 10, 1, false)

	app.SetRoot(flex, true)
}

func updateTable(final bool) {
	if table == nil {
		ShowModules()
	}
	row := len(search.Modules)
	if row == 0 {
		return
	}
	module := search.Modules[row-1]
	table.SetCell(row, 0, tview.NewTableCell(module.Path).
		SetTextColor(tcell.ColorWhite).
		SetAlign(tview.AlignLeft))
	table.SetCell(row, 1, tview.NewTableCell(module.Modified.Format("2006-01-02 15:04:05")).
		SetTextColor(tcell.ColorWhite).
		SetAlign(tview.AlignCenter))
	table.SetCell(row, 2, tview.NewTableCell(fmt.Sprintf("%d bytes", module.Size)).
		SetTextColor(tcell.ColorWhite).
		SetAlign(tview.AlignRight))

	if final {
		updateSearchStatus("[green]Search complete.")
	}
}

func logDebug(format string, args ...interface{}) {
	if !debugEnabled {
		return
	}
	fmt.Fprintf(debugTextView, format+"\n", args...)
}
