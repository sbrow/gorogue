package gorogue

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Log is the standard Logger for gorogue.
var Log *logger

type logger struct {
	*log.Logger
	debug bool
}

// SetLog sets the output for the standard Logger.
func SetLog(name string, debug bool) (*os.File, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("Something went wrong.")
	}
	f, err := os.Create(filepath.Join(filepath.Dir(filename), name))
	if err != nil {
		return nil, err
	}
	Log = &logger{Logger: log.New(f, "", log.LstdFlags), debug: debug}
	return f, nil
}

func (l *logger) Debug(v ...interface{}) {
	if l.debug {
		l.Print(v...)
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.Printf(format, v...)
	}
}

func (l *logger) Debugln(v ...interface{}) {
	if l.debug {
		l.Println(v...)
	}
}
