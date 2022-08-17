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

	if !isDev {
		// if in production mode
		if logFile == nil {
			// open log file if it isn't opened yet
			logFile, _ = os.OpenFile(
				GetExecutablePath("my-playground.log"),
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0666,
			)
		}

		if logFile != nil {
			// log file has been opened
			loggers[prefix] = log.New(logFile, label, flag)
			return loggers[prefix]
		}

		// ...pass to standard logger if cannot open log file
	}

	// in development mode, cache and return standard logger
	loggers[prefix] = log.New(log.Writer(), label, flag)
	return loggers[prefix]
}
