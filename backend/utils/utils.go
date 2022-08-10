package utils

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() (path string, err error) {
	path, err = os.Executable()
	if err != nil {
		return
	}

	path = filepath.Dir(path)
	return
}
