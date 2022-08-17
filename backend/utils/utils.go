package utils

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

var (
	isDev          bool
	executablePath string
)

func init() {
	var err error

	isDev = os.Getenv("WAILS_DEV") == "1"

	executablePath, err = os.Executable()
	if err != nil {
		log.Fatalf("fail to get executable path: %+v\n", err)
	}
	executablePath = filepath.Dir(executablePath)
}

func IsDev() bool {
	return isDev
}

func GetExecutablePath(elem ...string) string {
	return filepath.Join(append([]string{executablePath}, elem...)...)
}

func GetLocaleFromJSON(rawJson []byte) map[string]string {
	locale := make(map[string]string)
	json.Unmarshal(rawJson, &locale)
	return locale
}

func IsDirectoryExist(dir string) bool {
	d, err := os.Stat(dir)
	if os.IsExist(err) && d.IsDir() {
		return true
	}
	return false
}
