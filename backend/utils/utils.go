package utils

import (
	"encoding/json"
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

func GetLocaleFromJSON(rawJson []byte) map[string]string {
	locale := make(map[string]string)
	json.Unmarshal(rawJson, &locale)
	return locale
}
