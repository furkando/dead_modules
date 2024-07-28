package search

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rivo/tview"
)

type ModuleInfo struct {
	Path     string
	Modified time.Time
	Size     int64
}

var Modules []ModuleInfo
var SelectedModules = make(map[int]bool)

func SearchOldModules(rootDir string, app *tview.Application, updateTable func()) {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == "node_modules" {
			Modules = append(Modules, ModuleInfo{Path: path, Modified: info.ModTime(), Size: dirSize(path)})
			app.QueueUpdateDraw(func() {
				updateTable()
			})
			return filepath.SkipDir
		}
		return nil
	})

	app.QueueUpdateDraw(func() {
		updateTable()
	})
}

func dirSize(path string) int64 {
	var size int64
	filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size
}
