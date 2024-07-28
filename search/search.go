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

func SearchOldModules(rootDir string, app *tview.Application, updateTable func(final bool), logDebug func(format string, args ...interface{})) {
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logDebug("Error walking path %s: %v", path, err)
			return nil
		}
		logDebug("Scanning: %s, IsDir: %v, Base: %s", path, info.IsDir(), filepath.Base(path))

		if filepath.Base(path)[0] == '.' {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() && filepath.Base(path) == "node_modules" {
			Modules = append(Modules, ModuleInfo{Path: path, Modified: info.ModTime(), Size: dirSize(path)})
			app.QueueUpdateDraw(func() {
				updateTable(false)
			})
			return filepath.SkipDir
		}
		return nil
	})

	app.QueueUpdateDraw(func() {
		updateTable(true)
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
