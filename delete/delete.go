package delete

import (
	"dead_modules/search"
	"fmt"
	"os"

	"github.com/rivo/tview"
)

func DeleteSelectedModules(app *tview.Application, table *tview.Table) {
	for i, selected := range search.SelectedModules {
		if selected {
			module := search.Modules[i]
			app.QueueUpdateDraw(func() {
				table.GetCell(i+1, 0).SetText(fmt.Sprintf("[yellow][Deleting] %s", module.Path))
			})
			err := os.RemoveAll(module.Path)
			if err != nil {
				app.QueueUpdateDraw(func() {
					table.GetCell(i+1, 0).SetText(fmt.Sprintf("[red]Error deleting %s: %v", module.Path, err))
				})
			} else {
				app.QueueUpdateDraw(func() {
					table.GetCell(i+1, 0).SetText(fmt.Sprintf("[green][DELETED] %s", module.Path))
				})
			}
		}
	}
}
