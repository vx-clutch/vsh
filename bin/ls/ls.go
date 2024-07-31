package ls

import (
	"io/fs"
	"os"
)

func LS(path string) ([]fs.DirEntry, error) {
	if path == "" {
		path = "./"
	}
	return os.ReadDir(path)
}
