package utils

import (
	"log"
	"os"
)

var (
	logFile *os.File
	loggers = make(map[string]*log.Logger)
)

func Logger(prefix string) *log.Logger {
	if logger, ok := loggers[prefix]; ok {
		return logger
	}

	label := "[" + prefix + "] "
	flag := log.Ldate | log.Ltime | log.Lshortfile

	if os.Getenv("WAILS_DEV") != "1" && logFile == nil {
		// if in production mode and log file isn't opened yet
		var err error
		if logFile, err = os.OpenFile(
			GetExecutablePath("my-playground.log"),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		); err == nil {
			// open log file and cache logger
			loggers[prefix] = log.New(logFile, label, flag)
			return loggers[prefix]
		}
	}

	// in development mode, cache and return standard logger
	loggers[prefix] = log.New(log.Writer(), label, flag)
	return loggers[prefix]
}
