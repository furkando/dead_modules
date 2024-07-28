package delete

import (
	"dead_modules/search"
	"fmt"
	"os"
	"sync"

	"github.com/rivo/tview"
)

func DeleteSelectedModules(app *tview.Application, table *tview.Table, logDebug func(format string, args ...interface{})) {
	var wg sync.WaitGroup

	for i, selected := range search.SelectedModules {
		if selected {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				module := search.Modules[i]
				app.QueueUpdateDraw(func() {
					table.GetCell(i+1, 0).SetText(fmt.Sprintf("[yellow][Deleting[] %s", module.Path))
				})
				err := os.RemoveAll(module.Path)
				if err != nil {
					app.QueueUpdateDraw(func() {
						table.GetCell(i+1, 0).SetText(fmt.Sprintf("[red]Error deleting %s: %v", module.Path, err))
					})
				} else {
					app.QueueUpdateDraw(func() {
						table.GetCell(i+1, 0).SetText(fmt.Sprintf("[green][DELETED[] %s", module.Path))
					})
				}
			}(i)
		}
	}

	// Wait for all deletions to complete
	go func() {
		wg.Wait()
		app.QueueUpdateDraw(func() {
			// Update the status or UI after all deletions are complete
			logDebug("All selected modules have been deleted")
		})
	}()
}
