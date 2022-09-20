package log

import (
	"github.com/owner/repository/core/domain/store"
	"github.com/vite-cloud/go-zoup"
	"os"
)

const (
	// Store is the unique name of the logger store.
	Store = store.Store("logs")
	// DefaultLogFile is the name of the log file.
	DefaultLogFile = ":bin.log"

	// DefaultFileMode is the file mode used for the log file.
	DefaultFileMode = os.FileMode(0600)
)

// Log logs an internal event to the global logger.
func Log(level zoup.Level, message string, fields zoup.Fields) {
	dir, err := Store.Dir()
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(dir+"/"+DefaultLogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultFileMode)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}(file)

	logger := &zoup.FileWriter{File: file}

	err = logger.Write(level, message, fields)
	if err != nil {
		panic(err)
	}
}
