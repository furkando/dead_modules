package ui

import (
	"dead_modules/search"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const version = "1.0.0"

var app *tview.Application

func StartApp() error {
	app = tview.NewApplication()

	textView := tview.NewTextView().
		SetText("[yellow]Dead Modules v" + version + "[white]\n\nPress [green]Enter[white] to start searching for node_modules...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			showLoading()
			go search.SearchOldModules("/Users/furkan", app, updateTable)
		}
	})

	// Center the textView horizontally and vertically using a Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).     // Add a spacer at the top
		AddItem(textView, 0, 1, true). // Add the textView
		AddItem(nil, 0, 1, false)      // Add a spacer at the bottom

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
		AddItem(nil, 0, 1, false)      // Add a spacer at the bottom

	app.SetRoot(flex, true)
}

func updateTable() {
	if table == nil {
		ShowModules()
	}
	row := len(search.Modules)
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
}
