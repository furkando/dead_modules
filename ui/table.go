package ui

import (
	"dead_modules/delete"
	"dead_modules/search"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var table *tview.Table

func ShowModules() {
	table = tview.NewTable().
		SetFixed(1, 0).
		SetSelectable(true, false)

	headers := []string{"Folder Path", "Last Modified", "Size"}
	for i, header := range headers {
		table.SetCell(0, i, tview.NewTableCell(header).
			SetTextColor(tcell.ColorYellow).
			SetAlign(tview.AlignCenter).
			SetSelectable(false))
	}

	for i, module := range search.Modules {
		table.SetCell(i+1, 0, tview.NewTableCell(module.Path).
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
				table.GetCell(row, 0).SetTextColor(tcell.ColorGreen)
			} else {
				table.GetCell(row, 0).SetTextColor(tcell.ColorWhite)
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
					table.GetCell(row, 0).SetTextColor(tcell.ColorGreen)
				} else {
					table.GetCell(row, 0).SetTextColor(tcell.ColorWhite)
				}
			}
			return nil
		} else if event.Key() == tcell.KeyEnter {
			delete.DeleteSelectedModules(app, table)
			return nil
		}
		return event
	})

	// Center the table horizontally and vertically using a Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(table, 0, 1, true). // Add the table
		AddItem(debugTextView, 10, 1, false)

	app.SetRoot(flex, true)
}
