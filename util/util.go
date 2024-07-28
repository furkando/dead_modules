package util

import "dead_modules/search"

const (
	MaxPathLength = 100
)

func TruncatePath(path string, max int) string {
	if len(path) > max {
		return "..." + path[len(path)-max+3:]
	}
	return path
}

func FindModuleByPath(path string, modules []search.ModuleInfo) (int, *search.ModuleInfo) {
	for i, module := range modules {
		if module.Path == path {
			return i, &module
		}
	}
	return -1, nil
}
