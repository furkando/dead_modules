package ui

import (
	"dead_modules/delete"
	"dead_modules/search"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	table        *tview.Table
	searchStatus *tview.TextView
)

const (
	title = "[yellow]Dead Modules v1.0.0[white]"
)

func ShowModules() {
	table = tview.NewTable().
		SetFixed(1, 0).
		SetSelectable(true, false)

	headers := []string{"Path", "Last Modified", "Size"}
	for i, header := range headers {
		table.SetCell(0, i, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false).SetExpansion(1))
	}

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == ' ' {
			row, _ := table.GetSelection()
			if row > 0 {
				index := row - 1
				selectedModule := search.Modules[index]
				search.SelectedModules[selectedModule.Path] = !search.SelectedModules[selectedModule.Path]
				if search.SelectedModules[selectedModule.Path] {
					updateRowColor(row, tcell.ColorGreen)
				} else {
					updateRowColor(row, tcell.ColorWhite)
				}
			}
			return nil
		} else if event.Key() == tcell.KeyEnter {
			delete.DeleteSelectedModules(app, table, logDebug)
			return nil
		}
		return event
	})

	// Create title and status text views
	titleTextView := tview.NewTextView().
		SetText(title).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	infoTextView := tview.NewTextView().
		SetText("> Press [green]Space[] to select modules, [green]Enter[] to delete selected modules").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	searchStatus = tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true)

	// Create the main Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleTextView, 0, 1, false).
		AddItem(infoTextView, 0, 1, false).
		AddItem(table, 0, 10, true).
		AddItem(debugTextView, 0, 1, false).
		AddItem(searchStatus, 1, 1, false) // Add the searchStatus directly

	updateSearchStatus("[yellow]Searching...")
	app.SetRoot(flex, true)
}

func updateSearchStatus(status string) {
	searchStatus.SetText(status)
}

func updateRowColor(row int, color tcell.Color) {
	for i := 0; i < 3; i++ {
		table.GetCell(row, i).SetTextColor(color)
	}
}
