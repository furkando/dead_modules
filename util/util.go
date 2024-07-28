package util

const (
	MaxPathLength = 75
)

func TruncatePath(path string, max int) string {
	if len(path) > max {
		return "..." + path[len(path)-max+3:]
	}
	return path
}
