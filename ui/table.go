package ui

import (
	"dead_modules/delete"
	"dead_modules/search"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	table        *tview.Table
	searchStatus *tview.TextView
)

const (
	maxPathLength = 50
	title         = "[yellow]Dead Modules v1.0.0[white]"
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

	for i, module := range search.Modules {
		path := truncatePath(module.Path, maxPathLength)
		table.SetCell(i+1, 0, tview.NewTableCell(path).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft))

		table.SetCell(i+1, 1, tview.NewTableCell(module.Modified.Format("2006-01-02 15:04:05")).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignCenter))
		table.SetCell(i+1, 2, tview.NewTableCell(fmt.Sprintf("%d bytes", module.Size)).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignRight))
	}

	table.SetSelectedFunc(func(row, column int) {
		if row > 0 {
			index := row - 1
			search.SelectedModules[index] = !search.SelectedModules[index]
			if search.SelectedModules[index] {
				updateRowColor(row, tcell.ColorGreen)
			} else {
				updateRowColor(row, tcell.ColorWhite)
			}
		}
	})

	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == ' ' {
			row, _ := table.GetSelection()
			if row > 0 {
				index := row - 1
				search.SelectedModules[index] = !search.SelectedModules[index]
				if search.SelectedModules[index] {
					updateRowColor(row, tcell.ColorGreen)
				} else {
					updateRowColor(row, tcell.ColorWhite)
				}
			}
			return nil
		} else if event.Key() == tcell.KeyEnter {
			delete.DeleteSelectedModules(app, table)
			return nil
		}
		return event
	})

	// Create title and status text views
	titleTextView := tview.NewTextView().
		SetText(title).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	searchStatus = tview.NewTextView().
		SetText("").
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true)

	// Create the main Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(titleTextView, 1, 1, false).
		AddItem(nil, 0, 1, false).
		AddItem(table, 0, 10, true).
		AddItem(debugTextView, 0, 1, false).
		AddItem(searchStatus, 1, 1, false) // Add the searchStatus directly

	updateSearchStatus("[yellow]Searching...")
	app.SetRoot(flex, true)
}

func updateSearchStatus(status string) {
	searchStatus.SetText(status)
}

func truncatePath(path string, maxLength int) string {
	if len(path) > maxLength {
		return "..." + path[len(path)-maxLength+3:]
	}
	return path
}

func updateRowColor(row int, color tcell.Color) {
	for i := 0; i < 3; i++ {
		table.GetCell(row, i).SetTextColor(color)
	}
}
